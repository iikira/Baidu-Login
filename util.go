package main

import (
	"fmt"
	"github.com/astaxie/beego/session"
	"github.com/dop251/goja"
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
	"log"
	"net"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

// listAddresses 列出本地可用的 IP 地址
func listAddresses() (addresses []string) {
	ifaces, _ := net.Interfaces()
	addresses = make([]string, 0, len(ifaces))
	for _, iface := range ifaces {
		ifAddrs, _ := iface.Addrs()
		for _, ifAddr := range ifAddrs {
			switch v := ifAddr.(type) {
			case *net.IPNet:
				addresses = append(addresses, v.IP.String())
			case *net.IPAddr:
				addresses = append(addresses, v.IP.String())
			}
		}
	}
	return
}

// registerCookiejar 为 sess 如果没有 cookiejar , 就添加
func registerCookiejar(sess *session.Store) {
	if (*sess).Get("cookiejar") == nil { // 找不到 cookie 储存器
		initJar, _ := cookiejar.New(nil) // 初始化cookie储存器
		(*sess).Set("cookiejar", initJar)
	}
}

// getCookiejar 查找该 sessionID 下是否存在 cookiejar
func getCookiejar(sessionID string) (*cookiejar.Jar, error) {
	sessionStore, err := globalSessions.GetSessionStore(sessionID)
	if err != nil {
		return nil, err
	}
	jarInterface := sessionStore.Get("cookiejar")
	switch value := jarInterface.(type) {
	case *cookiejar.Jar:
		return jarInterface.(*cookiejar.Jar), nil
	default:
		return nil, fmt.Errorf("Unknown session cookiejar type: %s", value)
	}
}

// encryptePassword 通过调用 goja javascript虚拟机 来运行用来加密百度密码的javascript代码, 达到加密 password 的目的, 返回值为加密后的密码 encryptedPassword.
func encryptePassword(password string) (encryptedPassword string) {
	content, err := httpFilesBox.String("js/encrypt_password.tmpl.js")
	if err != nil {
		log.Println(err)
		return
	}
	vm := goja.New() // 初始化 javascript虚拟机
	rep := map[string]string{
		"RSAString":  getRSAString(),
		"ServerTime": serverTime,
		"Password":   password,
	}
	content = parseTemplate(content, rep)
	value, err := vm.RunString(content)
	if err != nil {
		log.Println(err)
		return
	}
	encryptedPassword = value.String()
	return
}

// parseTemplate 自己写的简易 template 解析器
func parseTemplate(content string, rep map[string]string) string {
	for k, v := range rep {
		content = strings.Replace(content, "{{."+k+"}}", v, 1)
	}
	return content
}

// parsePhoneAndEmail 抓取绑定百度账号的邮箱和手机号并插入至 json 结构
func (lj *loginJSON) parsePhoneAndEmail(sessionID string) {
	gotoURL, ok := (*lj)["data"]["gotoUrl"]
	if !ok {
		return
	}

	jar, err := getCookiejar(sessionID)
	if err != nil {
		log.Println(err)
		return
	}

	body, err := baiduUtil.Fetch("GET", gotoURL, jar, nil, nil)
	baiduUtil.PrintErrIfExist(err)

	// 使用正则表达式匹配
	rawPhone := regexp.MustCompile("您的手机号(.*?)是否能接收短信？").FindSubmatch(body)
	rawTokenAndU := regexp.MustCompile("token=(.*?)&u=(.*?)&secstate=").FindStringSubmatch(gotoURL)
	if len(rawPhone) >= 1 {
		(*lj)["data"]["phone"] = string(rawPhone[1])
	} else {
		(*lj)["data"]["phone"] = "未找到你的手机号"
	}
	if len(rawTokenAndU) >= 2 {
		(*lj)["data"]["token"] = rawTokenAndU[1]
		if u, err := url.Parse(rawTokenAndU[2]); err == nil {
			(*lj)["data"]["u"] = u.Path
		}
	}
	body, err = baiduUtil.Fetch("GET", gotoURL+"&finance=&clientfrom=&client=&adapter=2&enabledPage=email", jar, nil, nil)
	baiduUtil.PrintErrIfExist(err)
	rawEmail := regexp.MustCompile("您帐号绑定的邮箱(.*?)，能否接收邮件").FindSubmatch(body)
	if len(rawEmail) >= 1 {
		(*lj)["data"]["email"] = string(rawEmail[1])
	} else {
		(*lj)["data"]["email"] = "未找到你的邮箱地址"
	}
}

// parseCookies 解析 STOKEN, PTOKEN, BDUSS 并插入至 json 结构.
func (lj *loginJSON) parseCookies(targetURL string, jar *cookiejar.Jar) {
	url, _ := url.Parse(targetURL)
	targetList := []string{"STOKEN", "PTOKEN", "BDUSS"}
	cookies := jar.Cookies(url)
	for _, cookie := range cookies {
		for _, name := range targetList {
			if cookie.Name == name {
				(*lj)["data"][strings.ToLower(name)] = cookie.Value
			}
		}
	}
	(*lj)["data"]["cookieString"] = baiduUtil.GetURLCookieString(targetURL, jar) // 插入 cookie 字串
}

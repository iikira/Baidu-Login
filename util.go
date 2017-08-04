package main

import (
	"github.com/dop251/goja"
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
	"log"
	"net"
	"net/url"
	"regexp"
	"strings"
)

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

func encryptePassword(password string) (encryptedPassword string) {
	content, err := httpFilesBox.String("js/encrypt-password-tmpl.js")
	if err != nil {
		log.Println(err)
		return
	}
	vm := goja.New()
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

func parseTemplate(content string, rep map[string]string) string {
	for k, v := range rep {
		content = strings.Replace(content, "{{."+k+"}}", v, 1)
	}
	return content
}

// parsePhoneAndEmail 抓取绑定百度账号的邮箱和手机号并插入至 json 结构
func (lj *loginJSON) parsePhoneAndEmail() {
	gotoURL, ok := (*lj)["data"]["gotoUrl"]
	if !ok {
		return
	}
	body, err := baiduUtil.Fetch(gotoURL, jar, nil, nil)
	baiduUtil.PrintErrIfExist(err)

	// 使用正则表达式匹配
	rawPhone := regexp.MustCompile("您的手机号(.*?)是否能接收短信？").FindStringSubmatch(body)
	rawTokenAndU := regexp.MustCompile("token=(.*?)&u=(.*?)&secstate=").FindStringSubmatch(gotoURL)
	if len(rawPhone) >= 1 {
		(*lj)["data"]["phone"] = rawPhone[1]
	} else {
		(*lj)["data"]["phone"] = "未找到你的手机号"
	}
	if len(rawTokenAndU) >= 2 {
		(*lj)["data"]["token"] = rawTokenAndU[1]
		if u, err := url.Parse(rawTokenAndU[2]); err == nil {
			(*lj)["data"]["u"] = u.Path
		}
	}
	body, err = baiduUtil.Fetch(gotoURL+"&finance=&clientfrom=&client=&adapter=2&enabledPage=email", jar, nil, nil)
	baiduUtil.PrintErrIfExist(err)
	rawEmail := regexp.MustCompile("您帐号绑定的邮箱(.*?)，能否接收邮件").FindStringSubmatch(body)
	if len(rawEmail) >= 1 {
		(*lj)["data"]["email"] = rawEmail[1]
	} else {
		(*lj)["data"]["email"] = "未找到你的邮箱地址"
	}
}

func (lj *loginJSON) parseCookies(urlString string) {
	url, _ := url.Parse(urlString)
	targetList := []string{"STOKEN", "PTOKEN", "BDUSS"}
	cookies := jar.Cookies(url)
	for _, cookie := range cookies {
		for _, name := range targetList {
			if cookie.Name == name {
				(*lj)["data"][strings.ToLower(name)] = cookie.Value
			}
		}
	}
	(*lj)["data"]["cookieString"] = baiduUtil.GetURLCookieString(urlString, jar) // 插入 cookie 字串
}

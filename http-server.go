package main

import (
	"encoding/json"
	"fmt"
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

type loginJSON map[string]map[string]string

// startServer 启动服务
func startServer() {
	jar, _ = cookiejar.New(nil) // 初始化cookie储存器
	serverTime = getServerTime()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index.html", indexPage)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/template/js/login.js", loginJs)
	http.HandleFunc("/template/js/jquery.tiny.js", jquery)
	// http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(httpFilesBox.HTTPBox())))
	http.HandleFunc("/execBaiduLogin", execBaiduLogin)
	http.HandleFunc("/execVerifiedLogin", execVerify)
	http.HandleFunc("/sendCode", sendCode)

	fmt.Println("Server is starting...")

	// Print available URLs.
	for _, address := range listAddresses() {
		fmt.Printf(
			"URL: %s\n",
			(&url.URL{
				Scheme: "http",
				Host:   net.JoinHostPort(address, *port),
				Path:   "/",
			}).String(),
		)
	}
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":"+*port, nil))
}

// 根目录处理
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// 跳转到 /index.html
		w.Header().Set("Location", "/index.html")
		http.Error(w, "", 302)
	} else {
		http.Error(w, "404 Not Found", 404)
	}
}

// execBaiduLogin 发送 百度登录 请求
func execBaiduLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                                        // 解析 url 传递的参数
	username := strings.Join(r.Form["username"], "")     // 百度 用户名
	password := strings.Join(r.Form["password"], "")     // 密码
	verifycode := strings.Join(r.Form["verifycode"], "") // 图片验证码
	vcodestr := strings.Join(r.Form["vcodestr"], "")     // 与 图片验证码 相对应

	body, _ := baiduLogin(username, password, verifycode, vcodestr) //发送登录请求
	var lj loginJSON

	// 如果 json 解析出错, 直接输出{"error","错误信息"}
	if err := json.Unmarshal([]byte(body), &lj); err != nil {
		w.Write([]byte(`{"error","` + err.Error() + `"}`))
		return
	}

	lj.parseCookies("https://wappass.baidu.com")
	if lj != nil {
		switch lj["errInfo"]["no"] {
		case "400023", "400101": // 需要验证手机或邮箱
			lj.parsePhoneAndEmail()
		}
	}
	// 输出 json 编码
	byteBody, _ := json.MarshalIndent(&lj, "", "\t")
	w.Write(byteBody)
}

// execVerifiy 发送 提交验证码 请求
func execVerify(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                                  // 解析 url 传递的参数
	verifyType := strings.Join(r.Form["type"], "") // email/mobile
	token := strings.Join(r.Form["token"], "")     // token 不可或缺
	vcode := strings.Join(r.Form["vcode"], "")     // email/mobile 收到的验证码
	u := strings.Join(r.Form["u"], "")

	h := map[string]string{
		"Connection":                "keep-alive",
		"Host":                      "wappass.baidu.com",
		"Pragma":                    "no-cache",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	}

	url := fmt.Sprintf("https://wappass.baidu.com/passport/authwidget?v=1501743656994&vcode=%s&token=%s&u=%s&action=check&type=%s&tpl=&skin=&clientfrom=&adapter=2&updatessn=&bindToSmsLogin=&isnew=&card_no=&finance=&callback=%s", vcode, token, u, verifyType, "jsonp1")
	body, err := baiduUtil.Fetch(url, jar, nil, &h)
	if err != nil {
		log.Println(err)
		return
	}

	// 去除 body 的 callback 嵌套 "jsonp1(...)"
	body = strings.TrimLeft(body, "jsonp1(")
	body = strings.TrimRight(body, ")")

	var lj loginJSON

	// 如果 json 解析出错, 直接输出{"error","错误信息"}
	if err := json.Unmarshal([]byte(body), &lj); err != nil {
		fmt.Fprintf(w, `{"error","%s"}`, err.Error())
		return
	}

	u = fmt.Sprintf("%s&authsid=%s&fromtype=%s&bindToSmsLogin=", u, lj["data"]["authsid"], verifyType)

	_, err = baiduUtil.Fetch(u, jar, nil, nil)
	if err != nil {
		log.Println(err)
		return
	}

	lj.parseCookies(u)

	// 输出 json 编码
	byteBody, _ := json.MarshalIndent(&lj, "", "\t")
	w.Write(byteBody)
}

// sendCode 发送 获取验证码 请求
func sendCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	verifyType := strings.Join(r.Form["type"], "")
	token := strings.Join(r.Form["token"], "")
	if token == "" {
		w.Write([]byte(`{"error":"Token is null."}`))
		return
	}

	url := fmt.Sprintf("https://wappass.baidu.com/passport/authwidget?action=send&tpl=&type=%s&token=%s&from=&skin=&clientfrom=&adapter=2&updatessn=&bindToSmsLogin=&upsms=&finance=", verifyType, token)
	body, _ := baiduUtil.Fetch(url, jar, nil, nil)

	v := map[string]string{}
	rawMsg := regexp.MustCompile(`<p class="mod-tipinfo-subtitle">\s+(.*?)\s+</p>`).FindStringSubmatch(body)
	if len(rawMsg) >= 1 {
		v["msg"] = rawMsg[1]
	}

	byteBody, _ := json.MarshalIndent(&v, "", "\t")
	w.Write(byteBody)
}

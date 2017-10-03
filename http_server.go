package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iikira/baidu-tools/util"
	"log"
	"net"
	"net/http"
	"net/url"
	"regexp"
)

type loginJSON map[string]map[string]string

// startServer 启动服务
func startServer() {
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

// rootHandler 根目录处理
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		// 跳转到 /index.html
		w.Header().Set("Location", "/index.html")
		http.Error(w, "", 301)
	} else {
		http.Error(w, "404 Not Found", 404)
	}
}

// execBaiduLogin 发送 百度登录 请求
func execBaiduLogin(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	registerCookiejar(&sess)     // 如果没有 cookiejar , 就添加
	defer sess.SessionRelease(w) // 更新 session 储存

	r.ParseForm()                          // 解析 url 传递的参数
	username := r.Form.Get("username")     // 百度 用户名
	password := r.Form.Get("password")     // 密码
	verifycode := r.Form.Get("verifycode") // 图片验证码
	vcodestr := r.Form.Get("vcodestr")     // 与 图片验证码 相对应的字串

	jar, err := getCookiejar(sess.SessionID()) // 查找该 sessionID 下是否存在 cookiejar
	if err != nil {
		log.Println(err)
		return
	}

	serverTime = getServerTime(jar) // 趁此机会，访问一次百度页面，以初始化百度的 Cookie

	body, _ := baiduLogin(username, password, verifycode, vcodestr, jar) //发送登录请求
	var lj loginJSON

	// 如果 json 解析出错, 直接输出{"error","错误信息"}
	if err := json.Unmarshal([]byte(body), &lj); err != nil {
		w.Write([]byte(`{"error","` + err.Error() + `"}`))
		return
	}

	lj.parseCookies("https://wappass.baidu.com", jar)
	if lj != nil {
		switch lj["errInfo"]["no"] {
		case "400023", "400101": // 需要验证手机或邮箱
			lj.parsePhoneAndEmail(sess.SessionID())
		}
	}
	// 输出 json 编码
	byteBody, _ := json.MarshalIndent(&lj, "", "\t")
	w.Write(byteBody)
}

// execVerifiy 发送 提交验证码 请求
func execVerify(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	jar, err := getCookiejar(sess.SessionID()) // 查找该 sessionID 下是否存在 cookiejar
	if err != nil {
		log.Println(err)
		return
	}
	r.ParseForm()                    // 解析 url 传递的参数
	verifyType := r.Form.Get("type") // email/mobile
	token := r.Form.Get("token")     // token 不可或缺
	vcode := r.Form.Get("vcode")     // email/mobile 收到的验证码
	u := r.Form.Get("u")

	h := map[string]string{
		"Connection":                "keep-alive",
		"Host":                      "wappass.baidu.com",
		"Pragma":                    "no-cache",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	}

	url := fmt.Sprintf("https://wappass.baidu.com/passport/authwidget?v=1501743656994&vcode=%s&token=%s&u=%s&action=check&type=%s&tpl=&skin=&clientfrom=&adapter=2&updatessn=&bindToSmsLogin=&isnew=&card_no=&finance=&callback=%s", vcode, token, u, verifyType, "jsonp1")
	body, err := baiduUtil.Fetch("GET", url, jar, nil, h)
	if err != nil {
		log.Println(err)
		return
	}

	// 去除 body 的 callback 嵌套 "jsonp1(...)"
	body = bytes.TrimLeft(body, "jsonp1(")
	body = bytes.TrimRight(body, ")")

	var lj loginJSON

	// 如果 json 解析出错, 直接输出{"error","错误信息"}
	if err := json.Unmarshal([]byte(body), &lj); err != nil {
		fmt.Fprintf(w, `{"error","%s"}`, err.Error())
		return
	}

	// 最后一步要访问的 URL
	u = fmt.Sprintf("%s&authsid=%s&fromtype=%s&bindToSmsLogin=", u, lj["data"]["authsid"], verifyType) // url

	_, err = baiduUtil.Fetch("GET", u, jar, nil, nil)
	if err != nil {
		log.Println(err)
		return
	}

	lj.parseCookies(u, jar)

	// 输出 json 编码
	byteBody, _ := json.MarshalIndent(&lj, "", "\t")
	w.Write(byteBody)
}

// sendCode 发送 获取验证码 请求
func sendCode(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	jar, err := getCookiejar(sess.SessionID())
	if err != nil {
		log.Println(err)
		return
	}
	r.ParseForm()
	verifyType := r.Form.Get("type")
	token := r.Form.Get("token")
	if token == "" {
		w.Write([]byte(`{"error":"Token is null."}`))
		return
	}

	url := fmt.Sprintf("https://wappass.baidu.com/passport/authwidget?action=send&tpl=&type=%s&token=%s&from=&skin=&clientfrom=&adapter=2&updatessn=&bindToSmsLogin=&upsms=&finance=", verifyType, token)
	body, _ := baiduUtil.Fetch("GET", url, jar, nil, nil)

	v := map[string]string{}
	rawMsg := regexp.MustCompile(`<p class="mod-tipinfo-subtitle">\s+(.*?)\s+</p>`).FindSubmatch(body)
	if len(rawMsg) >= 1 {
		v["msg"] = string(rawMsg[1])
	}

	byteBody, _ := json.MarshalIndent(&v, "", "\t")
	w.Write(byteBody)
}

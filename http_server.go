package baidulogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iikira/BaiduPCS-Go/util"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
)

type loginJSON struct {
	ErrInfo struct {
		No  string `json:"no"`
		Msg string `json:"msg"`
	} `json:"errInfo"`
	Data struct {
		CodeString   string `json:"codeString"`
		GotoURL      string `json:"gotoUrl"`
		Token        string `json:"token"`
		U            string `json:"u"`
		AuthSID      string `json:"authsid"`
		Phone        string `json:"phone"`
		Email        string `json:"email"`
		BDUSS        string `json:"bduss"`
		PToken       string `json:"ptoken"`
		SToken       string `json:"stoken"`
		CookieString string `json:"cookieString"`
	} `json:"data"`
}

// StartServer 启动服务
func StartServer(port string) {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index.html", indexPage)
	http.HandleFunc("/favicon.ico", favicon)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(libFilesBox.HTTPBox())))
	http.HandleFunc("/cgi-bin/baidu/login", execBaiduLogin)
	http.HandleFunc("/cgi-bin/baidu/verifylogin", execVerify)
	http.HandleFunc("/cgi-bin/baidu/sendcode", sendCode)

	fmt.Println("Server is starting...")

	// Print available URLs.
	for _, address := range pcsutil.ListAddresses() {
		fmt.Printf(
			"URL: %s\n",
			(&url.URL{
				Scheme: "http",
				Host:   net.JoinHostPort(address, port),
				Path:   "/",
			}).String(),
		)
	}
	log.Fatal("ListenAndServe: ", http.ListenAndServe(":"+port, nil))
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
	registerBaiduClient(&sess)   // 如果没有 baiduClinet , 就添加
	defer sess.SessionRelease(w) // 更新 session 储存

	r.ParseForm()                          // 解析 url 传递的参数
	username := r.Form.Get("username")     // 百度 用户名
	password := r.Form.Get("password")     // 密码
	verifycode := r.Form.Get("verifycode") // 图片验证码
	vcodestr := r.Form.Get("vcodestr")     // 与 图片验证码 相对应的字串

	bc, err := getBaiduClient(sess.SessionID()) // 查找该 sessionID 下是否存在 cookiejar
	if err != nil {
		log.Println(err)
		return
	}

	body, err := bc.baiduLogin(username, password, verifycode, vcodestr) //发送登录请求
	if err != nil {
		fmt.Println(err)
	}

	var lj loginJSON

	// 如果 json 解析出错, 直接输出{"error","错误信息"}
	if err := json.Unmarshal(body, &lj); err != nil {
		lj.ErrInfo.No = "-1"
		lj.ErrInfo.Msg = "发送登录请求错误: " + err.Error()
		ljContent, _ := json.MarshalIndent(lj, "", "\t")
		w.Write(ljContent)
		return
	}

	lj.parseCookies("https://wappass.baidu.com", bc.Jar.(*cookiejar.Jar))

	switch lj.ErrInfo.No {
	case "400023", "400101": // 需要验证手机或邮箱
		lj.parsePhoneAndEmail(sess.SessionID())
	}

	// 输出 json 编码
	byteBody, _ := json.MarshalIndent(&lj, "", "\t")
	w.Write(byteBody)
}

// execVerifiy 发送 提交验证码 请求
func execVerify(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	bc, err := getBaiduClient(sess.SessionID())

	if err != nil {
		log.Println(err)
		return
	}
	r.ParseForm()                    // 解析 url 传递的参数
	verifyType := r.Form.Get("type") // email/mobile
	token := r.Form.Get("token")     // token 不可或缺
	vcode := r.Form.Get("vcode")     // email/mobile 收到的验证码
	u := r.Form.Get("u")

	header := map[string]string{
		"Connection":                "keep-alive",
		"Host":                      "wappass.baidu.com",
		"Pragma":                    "no-cache",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	}

	url := fmt.Sprintf("https://wappass.baidu.com/passport/authwidget?v=1501743656994&vcode=%s&token=%s&u=%s&action=check&type=%s&tpl=&skin=&clientfrom=&adapter=2&updatessn=&bindToSmsLogin=&isnew=&card_no=&finance=&callback=%s", vcode, token, u, verifyType, "jsonp1")
	body, err := bc.Fetch("GET", url, nil, header)
	if err != nil {
		log.Println(err)
		return
	}

	// 去除 body 的 callback 嵌套 "jsonp1(...)"
	body = bytes.TrimLeft(body, "jsonp1(")
	body = bytes.TrimRight(body, ")")

	var lj loginJSON

	// 如果 json 解析出错, 直接输出错误信息
	if err := json.Unmarshal(body, &lj); err != nil {
		lj.ErrInfo.No = "-2"
		lj.ErrInfo.Msg = "提交手机/邮箱验证码错误: " + err.Error()
		ljContent, _ := json.MarshalIndent(lj, "", "\t")
		w.Write(ljContent)
		return
	}

	// 最后一步要访问的 URL
	u = fmt.Sprintf("%s&authsid=%s&fromtype=%s&bindToSmsLogin=", u, lj.Data.AuthSID, verifyType) // url

	_, err = bc.Fetch("GET", u, nil, nil)
	if err != nil {
		log.Println(err)
		return
	}

	lj.parseCookies(u, bc.Jar.(*cookiejar.Jar))

	// 输出 json 编码
	byteBody, _ := json.MarshalIndent(&lj, "", "\t")
	w.Write(byteBody)
}

// sendCode 发送 获取验证码 请求
func sendCode(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r)
	bc, err := getBaiduClient(sess.SessionID())
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
	body, _ := bc.Fetch("GET", url, nil, nil)

	v := map[string]string{}
	rawMsg := regexp.MustCompile(`<p class="mod-tipinfo-subtitle">\s+(.*?)\s+</p>`).FindSubmatch(body)
	if len(rawMsg) >= 1 {
		v["msg"] = string(rawMsg[1])
	}

	byteBody, _ := json.MarshalIndent(&v, "", "\t")
	w.Write(byteBody)
}

package main

import (
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
	"net/http/cookiejar"
	"regexp"
)

// baiduLogin 发送 百度登录请求
func baiduLogin(username, password, verifycode, vcodestr string, jar *cookiejar.Jar) (body string, err error) {
	isPhone := "0"
	if regexp.MustCompile("^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$").MatchString(username) {
		isPhone = "1"
	}
	post := map[string]string{
		"username":     username,
		"password":     encryptePassword(password),
		"verifycode":   verifycode,
		"vcodestr":     vcodestr,
		"isphone":      isPhone,
		"action":       "login",
		"tpl":          "wimn",
		"uid":          "1501513545973_702",
		"subpro":       "wimn",
		"skin":         "default_v2",
		"connect":      "0",
		"dv":           "MDExAAoAZQALAL4ACgAAAF00ABMCAB6Ri4uL45fjk-Da9dqtzLzMrd6tg-GA6Y341rXat5gXAgARkZGOjoLZtOmy3a_yqci6588WAgAisMSvn7GHtYKzgLKKvY-4gLaCsYi5iL2NvIq6i7KFsYSyhgQCAAaTk5GSp54VAgAIkZGQzVMg6cEBAgAGkZmZkBG_BQIABJGRkZoQAgABkQYCACiRkZGGhoaGurq6vj09PTzx8fH3t7e3tDAwMDZ2dnZ1Q0NDR9XV1dbgBwIABJGRkZE",
		"getpassUrl":   "/passport/getpass?clientfrom=&adapter=0&ssid=&from=&authsite=&bd_page_type=&uid=1501513545973_702&pu=&tpl=wimn&u=https://m.baidu.com/usrprofile%3Fuid%3D1501513545973_702%23logined&type=&bdcm=060d5ffd462309f7e5529822720e0cf3d7cad665&tn=&regist_mode=&login_share_strategy=&subpro=wimn&skin=default_v2&client=&connect=0&smsLoginLink=1&loginLink=&bindToSmsLogin=&overseas=&is_voice_sms=&subpro=wimn&hideSLogin=&forcesetpwd=&regdomestic=",
		"mobilenum":    "undefined",
		"servertime":   serverTime,
		"gid":          "95C82F7-3D89-4094-8E13-6B9B756EDA4A",
		"logLoginType": "wap_loginTouch",
		"FP_UID":       "0b58c206c9faa8349576163341ef1321",
	}
	header := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Accept":           "application/json",
		"Referer":          "https://wappass.baidu.com/passport/?login&tpl=wimn&subpro=wimn&regtype=1&u=https://m.baidu.com/usrprofile#logined",
		"X-Requested-With": "XMLHttpRequest",
		"Connection":       "keep-alive",
	}

	body, err = baiduUtil.Fetch("https://wappass.baidu.com/wp/api/login", jar, &post, &header)
	return
}

// 获取百度服务器时间, 形如 "e362bacbae"
func getServerTime(jar *cookiejar.Jar) (serverTime string) {
	body, _ := baiduUtil.Fetch("https://wappass.baidu.com/wp/api/security/antireplaytoken", jar, nil, nil)
	rawServerTime := regexp.MustCompile(`,"time":"(.*?)"`).FindStringSubmatch(body)
	if len(rawServerTime) >= 1 {
		return rawServerTime[1]
	}
	return "e362bacbae"
}

// 获取百度 RSA 字串
func getRSAString() (RSAString string) {
	body, _ := baiduUtil.Fetch("https://wappass.baidu.com/static/touch/js/login_d9bffc9.js", nil, nil, nil)
	rawRSA := regexp.MustCompile(`,rsa:"(.*?)",error:`).FindStringSubmatch(body)
	if len(rawRSA) >= 1 {
		return rawRSA[1]
	}
	return "B3C61EBBA4659C4CE3639287EE871F1F48F7930EA977991C7AFE3CC442FEA49643212E7D570C853F368065CC57A2014666DA8AE7D493FD47D171C0D894EEE3ED7F99F6798B7FFD7B5873227038AD23E3197631A8CB642213B9F27D4901AB0D92BFA27542AE890855396ED92775255C977F5C302F1E7ED4B1E369C12CB6B1822F"
}

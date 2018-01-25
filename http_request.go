package baidulogin

import (
	"fmt"
	"github.com/iikira/Baidu-Login/bdcrypto"
	"github.com/iikira/BaiduPCS-Go/requester"
	"github.com/iikira/baidu-tools/util"
	"regexp"
)

type baiduClient struct {
	*requester.HTTPClient

	serverTime string
	traceid    string
}

func newBaiduClinet() *baiduClient {
	return &baiduClient{
		HTTPClient: requester.NewHTTPClient(),
	}
}

// baiduLogin 发送 百度登录请求
func (bc *baiduClient) baiduLogin(username, password, verifycode, vcodestr string) (body []byte, err error) {
	bc.getServerTime() // 趁此机会，访问一次百度页面，以初始化百度的 Cookie
	bc.getTraceID()

	isPhone := "0"
	if baiduUtil.ChinaPhoneRE.MatchString(username) {
		isPhone = "1"
	}

	enpass, err := bdcrypto.RsaEncrypt(bc.getRSAString(), []byte(password+bc.serverTime))
	if err != nil {
		fmt.Println(err)
	}

	post := map[string]string{
		"username":   username,
		"password":   fmt.Sprintf("%x", enpass),
		"verifycode": verifycode,
		"vcodestr":   vcodestr,
		"isphone":    isPhone,
		"action":     "login",
		"uid":        "1516806244773_357",
		"skin":       "default_v2",
		"connect":    "0",
		"dv":         "tk0.408376350146535171516806245342@oov0QqrkqfOuwaCIxUELn3oYlSOI8f51tbnGy-nk3crkqfOuwaCIxUou2iobENoYBf51tb4Gy-nk3cuv0ounk5vrkBynGyvn1QzruvN6z3drLJi6LsdFIe3rkt~4Lyz5ktfn1Qlrk5v5D5fOuwaCIxUobJWOI3~rkt~4Lyi5kBfni0vrk8~n15fOuwaCIxUobJWOI3~rkt~4Lyz5DQfn1oxrk0v5k5eruvN6z3drLneFYeVEmy-nk3c-qq6Cqw3h7CChwvi5-y-rkFizvmEufyr1By4k5bn15e5k0~n18inD0b5D8vn1Tyn1t~nD5~5T__ivmCpA~op5gr-wbFLhyFLnirYsSCIAerYnNOGcfEIlQ6I6VOYJQIvh515f51tf5DBv5-yln15f5DFy5myl5kqf5DFy5myvnktxrkT-5T__Hv0nq5myv5myv4my-nWy-4my~n-yz5myz4Gyx4myv5k0f5Dqirk0ynWyv5iTf5DB~rk0z5Gyv4kTf5DQxrkty5Gy-5iQf51B-rkt~4B__",
		"getpassUrl": "/passport/getpass?clientfrom=&adapter=0&ssid=&from=&authsite=&bd_page_type=&uid=1501513545973_702&pu=&tpl=wimn&u=https://m.baidu.com/usrprofile%3Fuid%3D1501513545973_702%23logined&type=&bdcm=060d5ffd462309f7e5529822720e0cf3d7cad665&tn=&regist_mode=&login_share_strategy=&subpro=wimn&skin=default_v2&client=&connect=0&smsLoginLink=1&loginLink=&bindToSmsLogin=&overseas=&is_voice_sms=&subpro=wimn&hideSLogin=&forcesetpwd=&regdomestic=",
		"mobilenum":  "undefined",
		"servertime": bc.serverTime,
		// "gid":          "7B3E207-25FD-4DA7-B482-A4039C935C86",
		"logLoginType": "wap_loginTouch",
		"FP_UID":       "0b58c206c9faa8349576163341ef1321",
		"traceid":      bc.traceid,
	}

	header := map[string]string{
		"Content-Type":     "application/x-www-form-urlencoded",
		"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Accept":           "application/json",
		"Referer":          "https://wappass.baidu.com/",
		"X-Requested-With": "XMLHttpRequest",
		"Connection":       "keep-alive",
	}

	body, err = bc.Fetch("POST", "https://wappass.baidu.com/wp/api/login", post, header)

	return
}

func (bc *baiduClient) getTraceID() {
	resp, err := bc.Get("http://wappass.baidu.com")
	if err != nil {
		fmt.Println(err)
		bc.traceid = err.Error()
	}

	defer resp.Body.Close()

	bc.traceid = resp.Header.Get("Trace-Id")
}

// 获取百度服务器时间, 形如 "e362bacbae"
func (bc *baiduClient) getServerTime() {
	body, _ := bc.Fetch("GET", "https://wappass.baidu.com/wp/api/security/antireplaytoken", nil, nil)
	rawServerTime := regexp.MustCompile(`,"time":"(.*?)"`).FindSubmatch(body)
	if len(rawServerTime) >= 1 {
		bc.serverTime = string(rawServerTime[1])
	}
	bc.serverTime = "e362bacbae"
}

// 获取百度 RSA 字串
func (bc *baiduClient) getRSAString() (RSAString string) {
	body, _ := bc.Fetch("GET", "https://wappass.baidu.com/static/touch/js/login_d9bffc9.js", nil, nil)
	rawRSA := regexp.MustCompile(`,rsa:"(.*?)",error:`).FindSubmatch(body)
	if len(rawRSA) >= 1 {
		return string(rawRSA[1])
	}
	return "B3C61EBBA4659C4CE3639287EE871F1F48F7930EA977991C7AFE3CC442FEA49643212E7D570C853F368065CC57A2014666DA8AE7D493FD47D171C0D894EEE3ED7F99F6798B7FFD7B5873227038AD23E3197631A8CB642213B9F27D4901AB0D92BFA27542AE890855396ED92775255C977F5C302F1E7ED4B1E369C12CB6B1822F"
}

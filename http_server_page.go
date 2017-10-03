package main

import (
	"github.com/iikira/baidu-tools/util"
	"log"
	"net/http"
	"text/template"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	sess, _ := globalSessions.SessionStart(w, r) // session start
	registerCookiejar(&sess)                     // 如果没有 cookiejar , 就添加

	// get file contents as string
	contents, err := httpFilesBox.String("index.tmpl.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl, err := template.New("index.html").Parse(contents)
	baiduUtil.PrintErrIfExist(err)
	baiduUtil.PrintErrIfExist(tmpl.Execute(w, version))
}

func loginJs(w http.ResponseWriter, r *http.Request) {
	byteCntents, err := httpFilesBox.Bytes("js/login.js")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(byteCntents)
}

func jquery(w http.ResponseWriter, r *http.Request) {
	byteCntents, err := httpFilesBox.Bytes("js/jquery.tiny.js")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(byteCntents)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "//www.baidu.com/favicon.ico")
	http.Error(w, "", 302)
}

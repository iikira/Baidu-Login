package main

import (
	"github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil"
	"log"
	"net/http"
	"text/template"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	indexContent, err := httpFilesBox.String("index-tmpl.html")
	if err != nil {
		log.Println(err)
		return
	}
	tmpl, err := template.New("index.html").Parse(indexContent)
	baiduUtil.PrintErrIfExist(err)
	baiduUtil.PrintErrIfExist(tmpl.Execute(w, version))
}

func loginJs(w http.ResponseWriter, r *http.Request) {
	// get file contents as string
	content, err := httpFilesBox.Bytes("js/login.js")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(content)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "https://www.baidu.com/favicon.ico")
	http.Error(w, "", 302)
}

func jquery(w http.ResponseWriter, r *http.Request) {
	content, err := httpFilesBox.Bytes("js/jquery.tiny.js")
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(content)
}

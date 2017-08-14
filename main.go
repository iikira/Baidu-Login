package main

import (
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/astaxie/beego/session"
	"os"
)

const (
	version = "beta v1.1"
)

var (
	port = flag.String("p", "9090", "HTTP server port.")
	v    = flag.Bool("v", false, "Displays version information.")

	serverTime     string
	httpFilesBox   *rice.Box
	err            error
	globalSessions *session.Manager
)

//go:generate $GOPATH/bin/rice embed-go
func init() {
	flag.Parse()
	if *v {
		fmt.Printf("百度BDUSS获取工具 %s Go语言.\n", version)
		fmt.Println("Copyright (c) 2017, iikira.")
		os.Exit(0)
	}
	fmt.Printf("百度BDUSS获取工具 %s Go语言.\n", version)
	fmt.Printf("运行 %s -h 以获取更多帮助.\n\n", os.Args[0])
}

func init() {
	globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
	})
	go globalSessions.GC()
}

func main() {
	httpFilesBox, err = rice.FindBox("http-files")
	if err != nil {
		fmt.Println(err)
		return
	}
	startServer()
}

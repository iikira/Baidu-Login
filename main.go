package main

import (
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"net/http/cookiejar"
	"os"
	// "time"
)

const (
	version = "beta v1.0"
)

var (
	httpFilesBox = rice.MustFindBox("http-files")
	port         = flag.String("p", "9090", "HTTP server port.")
	v            = flag.Bool("v", false, "Displays version information.")
	serverTime   string
	jar          *cookiejar.Jar
)

//go:generate $GOPATH/bin/rice embed-go
func init() {
	flag.Parse()
	if *v {
		fmt.Printf("百度BDUSS获取工具 %s (Go runtime go1.8.3).\n", version)
		fmt.Println("Copyright © 2017, iikira.")
		os.Exit(0)
	}
	fmt.Printf("百度BDUSS获取工具 %s (Go runtime go1.8.3).\n", version)
	fmt.Printf("运行 %s -h 以获取更多帮助.\n\n", os.Args[0])
	jar, _ = cookiejar.New(nil)
	serverTime = getServerTime()
	// go func() {
	// 	for {
	// 		fmt.Println(jar)
	// 		time.Sleep(3e9)
	// 	}
	// }()
}

func main() {
	startServer()
}

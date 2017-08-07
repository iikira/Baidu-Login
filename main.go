package main

import (
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"net/http/cookiejar"
	"os"
)

const (
	version = "beta v1.0"
)

var (
	port = flag.String("p", "9090", "HTTP server port.")
	v    = flag.Bool("v", false, "Displays version information.")

	serverTime   string
	jar          *cookiejar.Jar
	httpFilesBox *rice.Box
	err          error
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
}

func main() {
	httpFilesBox, err = rice.FindBox("http-files")
	if err != nil {
		fmt.Println(err)
		return
	}
	startServer()
}

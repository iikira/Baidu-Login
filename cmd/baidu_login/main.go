package main

import (
	"flag"
	"fmt"
	"github.com/iikira/Baidu-Login"
	"os"
)

const (
	version = "beta v1.1"
)

var (
	port = flag.String("p", "9090", "HTTP server port.")
	v    = flag.Bool("version", false, "Displays version information.")
)

func init() {
	flag.Parse()
	baidulogin.Version = version
	if *v {
		fmt.Printf("Baidu BDUSS Tool %s Golang.\n", version)
		fmt.Println("Copyright (c) 2017, iikira.")
		os.Exit(0)
	}
	fmt.Printf("Baidu BDUSS Tool %s Golang.\n", version)
	fmt.Printf("Try '%s -h' for more infomation.\n\n", os.Args[0])
}

func main() {
	baidulogin.StartServer(*port) // 启动服务
}

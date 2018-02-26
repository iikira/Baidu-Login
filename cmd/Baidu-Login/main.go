package main

import (
	"flag"
	"fmt"
	"github.com/iikira/Baidu-Login"
	"os"
)

var (
	port = flag.String("p", "9090", "HTTP server port.")
	v    = flag.Bool("version", false, "Displays version information.")
)

func init() {
	flag.Parse()
	if *v {
		fmt.Printf("Baidu BDUSS Tool %s Golang.\n", baidulogin.Version)
		fmt.Println("Copyright (c) 2017, iikira.")
		os.Exit(0)
	}
	fmt.Printf("Baidu BDUSS Tool %s Golang.\n", baidulogin.Version)
	fmt.Printf("Try '%s -h' for more infomation.\n\n", os.Args[0])
}

func main() {
	baidulogin.StartServer(*port) // 启动服务
}

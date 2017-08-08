# 百度BDUSS获取工具 beta v1.0 (Go runtime go1.8.3)

## 功能

百度: 获取百度账户 BDUSS, PTOKEN, STOKEN值

百度: 支持在线 手机／邮箱 安全验证（beta）

## 如何使用

Go语言程序, 可直接下载使用 [点此查看发布页](https://github.com/iikira/Baidu-Login/releases)

在 (Linux, MacOS)终端／(Windows)命令提示符 中运行。

Windows 也可双击程序运行。

本程序会开启本地服务器, 默认端口为9090.

程序会列出一些本地服务器网址, 浏览器访问网址即可使用。

## 如何手动编译安装

### 1. 安装Go语言运行环境

* 访问 [Go语言官网](https://golang.org) 下载安装Golang
* 设置GOPATH环境变量

Linux: 
```shell
export GOPATH=/path/to/your/gopath
```
Windows:
```shell
set GOPATH=C:\path\to\your\gopath
```

如果提示找不到 go 命令, 请先设置 PATH , 以Linux为例
```shell
export PATH=$PATH:$GOROOT/bin
```
$GOROOT 即Go语言的安装目录

### 2. 安装

#### 编译安装(需要设置GOPATH环境变量)
```shell
go get -u -v github.com/iikira/Baidu-Login
```
编译生成的文件在GOPATH的bin目录下

#### 手动编译安装(需要设置GOPATH环境变量)

1. 下载程序源码到源码目录
```shell
git clone https://github.com/iikira/Baidu-Login.git
```
2. 安装依赖包

```shell
go get -u -v github.com/iikira/Tieba-Cloud-Sign-Backend/baiduUtil
go get -u -v github.com/GeertJohan/go.rice/rice
go get -u -v github.com/dop251/goja
```

3. 进入源码目录，编译

```shell
cd Baidu-Login
go build
```

## 如何将静态资源打包进程序

强烈建议使用 rice 将 http-files 目录内的文件打包进程序

详情: [https://github.com/GeertJohan/go.rice](https://github.com/GeertJohan/go.rice)

设置好环境变量 GOPATH:
```shell
export GOTATH=/path/to/your/gopath
```

安装依赖包:
```shell
go get github.com/GeertJohan/go.rice
go get github.com/GeertJohan/go.rice/rice
```

最后执行以下命令编译程序:
```shell
go generate
go build
```

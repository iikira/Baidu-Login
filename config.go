package baidulogin

import (
	"github.com/GeertJohan/go.rice"
	"github.com/astaxie/beego/session"
)

var (
	Version = "v1"

	templateFilesBox = rice.MustFindBox("http-files/template")
	libFilesBox      = rice.MustFindBox("http-files/static")

	globalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
	}) // 全局 sessions 管理器
)

func init() {
	go globalSessions.GC()
}

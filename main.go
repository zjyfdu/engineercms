package main

import (
	_ "github.com/3xxx/engineercms/controllers/version"
	_ "github.com/3xxx/engineercms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/3xxx/engineercms/models"
)

func main() {
	// beego.AddFuncMap("dict", dict)
	beego.AddFuncMap("loadtimes", loadtimes)
	beego.AddFuncMap("subsuffix", subsuffix)
	//开启orm调试模式
	orm.Debug = true
	//创建附件目录ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
	os.Mkdir("attachment", os.ModePerm)
	//创建轮播图片存放目录
	os.Mkdir("attachment//carousel", os.ModePerm)
	//自动建表
	orm.RunSyncdb("default", false, true)
	models.InsertUser()
	// insertGroup()
	// models.InsertRole()
	beego.Run()
}

//显示页面加载时间
func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

//去除扩展名
func subsuffix(in string) string {
	fileSuffix := path.Ext(in)
	return strings.TrimSuffix(in, fileSuffix)
}

// GoLang 如何在网页显示当前环境的版本号
func server() {
	http.HandleFunc("/version", version)
	http.ListenAndServe(":8080", nil)
}

func version(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf("%s", out))
}

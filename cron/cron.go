package cron

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func TodayDocSync() error {
	docs, err := models.GetDocs()
	if err != nil {
		beego.Error(err)
		return nil
	}
	nowTime := time.Now()
	todayDate := nowTime.Format("20060102")
	todayDoc := ""
	todayKey := ""

	for _, w := range docs {
		Attachments, _ := models.GetOnlyAttachments(w.Id)
		ext := path.Ext(Attachments[0].FileName)
		if ext != ".xlsx" {
			continue
		}
		if strings.Contains(w.Title, todayDate) {
			todayDoc = w.Title
			todayKey = strconv.FormatInt(Attachments[0].Updated.UnixNano(), 10)
			break
		}
	}
	beego.Info(todayKey, todayDoc)

	url := beego.AppConfig.String("documentserverip") + "/coauthoring/CommandService.ashx"

	//json序列化
	postJson := `{"c":"forcesave","key":"` + todayKey + `"}`
	var jsonStr = []byte(postJson)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}

func Init2() {
	tk := toolbox.NewTask("docSync", beego.AppConfig.String("forcesavecron"), TodayDocSync)
	err := tk.Run()
	if err != nil {
		fmt.Println(err)
	}
	toolbox.AddTask("docSync", tk)
	toolbox.StartTask()
}

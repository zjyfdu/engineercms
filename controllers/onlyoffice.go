package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/3xxx/engineercms/cron"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"

	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

type OnlyController struct {
	beego.Controller
}

type Callback struct {
	Key           string    `json:"key"`
	Status        int       `json:"status"`
	Url           string    `json:"url"`
	Changesurl    string    `json:"changesurl"`
	History       history1  `json:"history"`
	Users         []string  `json:"users"`
	Actions       []action  `json:"actions"`
	Lastsave      time.Time `json:"lastsave"`
	Notmodified   bool      `json:"notmodified"`
	Forcesavetype int       `json:"forcesavetype"`
}

type action struct {
	Type   int   `json:"type"`
	Userid int64 `json:"userid"`
}

type history1 struct {
	ServerVersion string    `json:"serverVersion"`
	Changes       []change  `json:"changes"`
	Created       time.Time `json:"created"`
	ChangesUrl    string    `json:"changesurl"`
	FileUrl       string    `json:"fileurl"`
	Key           string    `json:"key"`
	User          User1     `json:"user"`
	Version       int       `json:"version"`
}

type change struct {
	Created string `json:"created"` //time.Time
	User    User1  `json:"user"`
}

type User1 struct {
	Id   string `json:"id"` //必须大写才能在tpl中显示{{.json}}
	Name string `json:"name"`
}

//构造changesurl结构
type changesurl struct {
	Version    int    `json:"version"`
	ChangesUrl string `json:"changesurl"`
}

type Onlyoffice1 struct {
	Id      int64
	Code    string
	Title   string
	Ext     string
	Created time.Time
	Updated time.Time
}

type OnlyLink struct {
	Id        int64
	Code      string
	Title     string
	Label     string
	End       time.Time
	Principal string
	Uid       int64
	Uname     string
	Created   time.Time
	Updated   time.Time
	Docxlink  []DocxLink
}

type DocxLink struct {
	Id         int64
	Title      string
	Suffix     string
	Permission string
	Created    time.Time
	Updated    time.Time
}

//权限表提交的table中json数据解析成struct
type Rolepermission struct {
	Id         int64
	Name       string `json:"name"`
	Rolenumber string
	Permission string `json:"role"`
}

func cellName(col, row int) string {
	res := ""
	col, row = col+1, row+1
	for col > 0 {
		col--
		res = string(byte(col%26)+'A') + res
		col /= 26
	}
	res += strconv.Itoa(row)
	return res
}

func copyEmptyXlsx(dst string, src string) (err error) {
	rowkeep, _ := beego.AppConfig.Int("rowkeep")
	colkeep, _ := beego.AppConfig.Int("colkeep")
	logs.Info(rowkeep, colkeep)

	f, err := excelize.OpenFile(src)
	if err != nil {
		return err
	}

	nowTime := time.Now()
	nowTimeStr := nowTime.Format("0102")

	sheetnames := f.GetSheetMap()
	fmt.Println(sheetnames)
	for k, v := range sheetnames {
		if k == 1 {
			f.SetSheetName(v, nowTimeStr)
		} else {
			fmt.Println(v)
			f.DeleteSheet(v)
		}
	}

	rows, err := f.GetRows(nowTimeStr)
	for rowid, row := range rows {
		if rowid < rowkeep {
			continue
		}
		for colid, _ := range row {
			if colid < colkeep {
				continue
			}
			f.SetCellDefault(nowTimeStr, cellName(colid, rowid), "")
			f.DeletePicture(nowTimeStr, cellName(colid, rowid))
		}
	}

	// 根据指定路径保存文件
	if err := f.SaveAs(dst); err != nil {
		return err
	}

	return nil
}

//文档管理页面
func (c *OnlyController) Get() {
	//取得客户端用户名

	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	// 复制最近的一篇
	docs, err := models.GetDocs()
	if err != nil {
		beego.Error(err)
	}

	nowTime := time.Now()
	todayDate := nowTime.Format("20060102")
	var yesDate string = ""

	// 上一个日报的日期，用正则搜索
	yesRegex, _ := regexp.Compile(`20\d{2}(0[1-9]|1[0-2])(0[1-9]|[1-2][0-9]|3[0-1])`)

	var todayDoc string = ""
	var yesDoc string = ""
	var ext string = ""

	for _, w := range docs {
		Attachments, err := models.GetOnlyAttachments(w.Id)
		ext = path.Ext(Attachments[0].FileName)
		logs.Info(ext)
		logs.Info(ext)
		logs.Info(ext)
		logs.Info(ext)
		if ext != ".xlsx" {
			continue
		}
		if strings.Contains(w.Title, todayDate) {
			todayDoc = w.Title
			break
		}
		if yesDate == "" && yesRegex.FindString(w.Title) != "" {
			yesDate = yesRegex.FindString(w.Title)
			yesDoc = w.Title
			Attachments, err = models.GetOnlyAttachments(w.Id)
			if err != nil {
				beego.Error(err)
			}
			ext = path.Ext(Attachments[0].FileName)
		}
	}

	// 有昨天的日报，但是没有今天的日报
	if yesDoc != "" && todayDoc == "" {
		todayDoc = strings.Replace(yesDoc, yesDate, todayDate, -1)

		date, err := time.Parse("20060102", todayDate)
		if err != nil {
			beego.Error(err)
		}
		prodId, err := models.AddDoc(todayDoc, todayDoc, "", "复制"+yesDate, date, uid)
		if err != nil {
			beego.Error(err)
		}

		attachmentname := todayDoc + ext

		_, _, err2 := models.AddOnlyAttachment(attachmentname, 0, 0, prodId)
		if err2 != nil {
			beego.Error(err2)
		} else {
			DiskDirectory := "./attachment/onlyoffice"
			todayDoc = DiskDirectory + "/" + todayDoc + ext
			yesfilepath := DiskDirectory + "/" + yesDoc + ext
			beego.Info(todayDoc)
			beego.Info(yesfilepath)

			err = copyEmptyXlsx(todayDoc, yesfilepath)
			if err != nil {
				source, _ := os.Open(yesfilepath)
				defer source.Close()
				destination, _ := os.Create(todayDoc)
				defer destination.Close()
				logs.Info(err)
				_, _ = io.Copy(destination, source)
			} else {
				logs.Info("copy empty succeed")
			}
		}
	}

	c.Data["IsOnlyOffice"] = true
	c.TplName = "onlyoffice/docs.tpl"
}

// 这里的函数才获取文档列表
func (c *OnlyController) GetData() {
	//1.取得客户端用户名
	var err error
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	useridstring := strconv.FormatInt(uid, 10)
	//增加admin，everyone，isme

	var myRes, roleRes [][]string
	if useridstring != "0" {
		myRes = e.GetPermissionsForUser(useridstring)
	}
	myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据

	docs, err := models.GetDocs()
	if err != nil {
		beego.Error(err)
	}

	link := make([]OnlyLink, 0)
	Docxslice := make([]DocxLink, 0)
	for _, w := range docs {
		linkarr := make([]OnlyLink, 1)
		linkarr[0].Id = w.Id
		linkarr[0].Code = w.Code
		linkarr[0].Title = w.Title
		linkarr[0].Label = w.Label
		linkarr[0].End = w.End
		linkarr[0].Principal = w.Principal
		linkarr[0].Uid = w.Uid
		user := models.GetUserByUserId(w.Uid)
		linkarr[0].Uname = user.Nickname
		linkarr[0].Created = w.Created
		linkarr[0].Updated = w.Updated

		Attachments, err := models.GetOnlyAttachments(w.Id)
		if err != nil {
			beego.Error(err)
		}
		//docid——me——1
		for _, v := range Attachments {
			// beego.Info(v.FileName)
			// fileext := path.Ext(v.FileName)
			docxarr := make([]DocxLink, 1)
			docxarr[0].Permission = "1"
			//查询v.Id是否和myres的V1路径后面的id一致，如果一致，则取得V2（权限）
			//查询用户具有的权限
			// beego.Info(useridstring)
			if useridstring != "0" { //如果是登录用户，则设置了权限的文档不能看
				if w.Uid != uid { //如果不是作者本人
					for _, k := range myResall { //所有设置了权限的都不能看
						// beego.Info(k)
						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
							// docxarr[0].Permission = "4"
							docxarr[0].Permission = "3"
						}
					}

					for _, k := range myRes { //如果与登录用户对应上，则赋予权限
						// beego.Info(k)
						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
							docxarr[0].Permission = k[2]
						}
					}
					roles, _ := e.GetRolesForUser(useridstring) //取出用户的所有角色
					for _, w1 := range roles {                  //2018.4.30修改这个bug，这里原先w改为w1
						roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
						for _, k := range roleRes {
							// beego.Info(k)
							if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
								// docxarr[0].Permission = k[2]
								int1, err := strconv.Atoi(k[2])
								if err != nil {
									beego.Error(err)
								}
								int2, err := strconv.Atoi(docxarr[0].Permission)
								if err != nil {
									beego.Error(err)
								}
								if int1 < int2 {
									docxarr[0].Permission = k[2] //按最小值权限
								}
								//补充everyone
							}
						}
					}
				} //如果是用户自己的文档，则permission为1，默认
			} else { //如果用户没登录，则设置了权限的文档不能看
				for _, k := range myResall { //所有设置了权限的不能看
					// beego.Info(k)
					if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
						docxarr[0].Permission = "3"
					}
				}
			}

			docxarr[0].Id = v.Id
			docxarr[0].Title = v.FileName
			if path.Ext(v.FileName) == ".docx" || path.Ext(v.FileName) == ".DOCX" || path.Ext(v.FileName) == ".doc" || path.Ext(v.FileName) == ".DOC" {
				docxarr[0].Suffix = "docx"
			} else if path.Ext(v.FileName) == ".wps" || path.Ext(v.FileName) == ".WPS" {
				docxarr[0].Suffix = "docx"
			} else if path.Ext(v.FileName) == ".XLSX" || path.Ext(v.FileName) == ".xlsx" || path.Ext(v.FileName) == ".XLS" || path.Ext(v.FileName) == ".xls" {
				docxarr[0].Suffix = "xlsx"
			} else if path.Ext(v.FileName) == ".ET" || path.Ext(v.FileName) == ".et" {
				docxarr[0].Suffix = "xlsx"
			} else if path.Ext(v.FileName) == ".pptx" || path.Ext(v.FileName) == ".PPTX" || path.Ext(v.FileName) == ".ppt" || path.Ext(v.FileName) == ".PPT" {
				docxarr[0].Suffix = "pptx"
			} else if path.Ext(v.FileName) == ".DPS" || path.Ext(v.FileName) == ".dps" {
				docxarr[0].Suffix = "pptx"
			} else if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				docxarr[0].Suffix = "pdf"
			} else if path.Ext(v.FileName) == ".txt" || path.Ext(v.FileName) == ".TXT" {
				docxarr[0].Suffix = "txt"
			}
			Docxslice = append(Docxslice, docxarr...)
		}
		linkarr[0].Docxlink = Docxslice
		Docxslice = make([]DocxLink, 0) //再把slice置0
		//无权限的不显示
		//如果permission=4，则不赋值给link
		if linkarr[0].Docxlink[0].Permission != "4" {
			link = append(link, linkarr...)
		}
	}
	c.Data["json"] = link //products
	c.ServeJSON()
}

//协作页面的显示
//补充权限判断
//补充token
func (c *OnlyController) OnlyOffice() {
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}

	//docid——uid——me
	doc, err := models.Getdocbyid(onlyattachment.DocId)
	if err != nil {
		beego.Error(err)
	}

	var useridstring, Permission string
	var myRes, roleRes [][]string

	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	useridstring = strconv.FormatInt(uid, 10)
	//增加admin，everyone，isme

	var usersessionid string //客户端sesssionid
	if islogin {
		usersessionid = c.Ctx.Input.Cookie("hotqinsessionid")
	}
	myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据
	if uid != 0 {                           //无论是登录还是ip查出了用户id
		myRes = e.GetPermissionsForUser(useridstring)
		if doc.Uid == uid { //isme
			Permission = "1"
		} else { //如果是登录用户，则设置了权限的文档根据权限查看
			Permission = "1"
			for _, k := range myResall {
				if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
					Permission = "3"
				}
			}
			for _, k := range myRes {
				if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
					Permission = k[2]
				}
			}

			roles, _ := e.GetRolesForUser(useridstring) //取出用户的所有角色
			for _, w1 := range roles {                  //2018.4.30修改这个bug，这里原先w改为w1
				roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
				for _, k := range roleRes {
					// beego.Info(k)
					if id == path.Base(k[1]) {
						// docxarr[0].Permission = k[2]
						int1, err := strconv.Atoi(k[2])
						if err != nil {
							beego.Error(err)
						}
						int2, err := strconv.Atoi(Permission)
						if err != nil {
							beego.Error(err)
						}
						if int1 < int2 {
							Permission = k[2] //按最小值权限
						}
						//补充everyone权限，如果登录用户权限大于everyone，则用小的
					}
				}
			}
		}
		// c.Data["Uid"] = user.Id
		// userrole = user.Role
	} else { //如果用户没登录，则设置了权限的文档不能看
		Permission = "1"
		for _, k := range myResall { //所有设置了权限的不能看
			if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
				Permission = "3"
				// Permission = "4"
			}
			//如果设置了everyone用户权限，则按everyone的权限
		}
		c.Data["Uname"] = c.Ctx.Input.IP()
		c.Data["Uid"] = c.Ctx.Input.IP()
	}

	if Permission == "1" {
		c.Data["Mode"] = "edit"
		c.Data["Edit"] = true
		c.Data["Review"] = true
		c.Data["Comment"] = true
		c.Data["Download"] = true
		c.Data["Print"] = true
	} else if Permission == "2" {
		c.Data["Mode"] = "edit"
		c.Data["Edit"] = false
		c.Data["Review"] = true
		c.Data["Comment"] = true
		c.Data["Download"] = false
		c.Data["Print"] = false
	} else if Permission == "3" {
		c.Data["Mode"] = "view"
		c.Data["Edit"] = false
		c.Data["Review"] = false
		c.Data["Comment"] = false
		c.Data["Download"] = false
		c.Data["Print"] = false
	} else if Permission == "4" {
		return
	}

	c.Data["Doc"] = onlyattachment
	c.Data["Sessionid"] = usersessionid
	c.Data["attachid"] = idNum
	c.Data["Key"] = strconv.FormatInt(onlyattachment.Updated.UnixNano(), 10)

	//构造[]history
	history, err := models.GetOnlyHistory(onlyattachment.Id)
	if err != nil {
		beego.Error(err)
	}

	onlyhistory := make([]history1, 0)
	onlychanges := make([]change, 0)
	onlychangesurl := make([]changesurl, 0)
	for _, v := range history {
		aa := make([]history1, 1)
		cc := make([]changesurl, 1)
		// aa[0].Created = v.Created
		aa[0].Key = v.HistoryKey
		aa[0].User.Id = strconv.FormatInt(v.UserId, 10) //

		if v.UserId != 0 {
			user := models.GetUserByUserId(v.UserId)
			aa[0].User.Name = user.Nickname
		}
		aa[0].ServerVersion = v.ServerVersion
		aa[0].Version = v.Version
		aa[0].FileUrl = v.FileUrl
		aa[0].ChangesUrl = v.ChangesUrl
		//取得changes
		changes, err := models.GetOnlyChanges(v.HistoryKey)
		if err != nil {
			beego.Error(err)
		}
		for _, v1 := range changes {
			bb := make([]change, 1)
			bb[0].Created = v1.Created
			bb[0].User.Id = v1.UserId
			bb[0].User.Name = v1.UserName
			onlychanges = append(onlychanges, bb...)
		}
		aa[0].Changes = onlychanges
		// aa[0].ChangesUrl = v.ChangesUrl
		aa[0].Created = v.Created

		cc[0].Version = v.Version
		cc[0].ChangesUrl = v.ChangesUrl
		onlychanges = make([]change, 0)
		onlyhistory = append(onlyhistory, aa...)
		onlychangesurl = append(onlychangesurl, cc...)
	}

	c.Data["onlyhistory"] = onlyhistory
	c.Data["changesurl"] = onlychangesurl

	historyversion, err := models.GetOnlyHistoryVersion(onlyattachment.Id)
	if err != nil {
		beego.Error(err)
	}
	var first int
	for _, v := range historyversion {
		if first < v.Version {
			first = v.Version
		}
	}
	c.Data["currentversion"] = first

	if path.Ext(onlyattachment.FileName) == ".docx" || path.Ext(onlyattachment.FileName) == ".DOCX" {
		c.Data["fileType"] = "docx"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".wps" || path.Ext(onlyattachment.FileName) == ".WPS" {
		c.Data["fileType"] = "docx"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".XLSX" || path.Ext(onlyattachment.FileName) == ".xlsx" {
		c.Data["fileType"] = "xlsx"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".ET" || path.Ext(onlyattachment.FileName) == ".et" {
		c.Data["fileType"] = "xlsx"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".pptx" || path.Ext(onlyattachment.FileName) == ".PPTX" {
		c.Data["fileType"] = "pptx"
		c.Data["documentType"] = "presentation"
	} else if path.Ext(onlyattachment.FileName) == ".dps" || path.Ext(onlyattachment.FileName) == ".DPS" {
		c.Data["fileType"] = "pptx"
		c.Data["documentType"] = "presentation"
	} else if path.Ext(onlyattachment.FileName) == ".doc" || path.Ext(onlyattachment.FileName) == ".DOC" {
		c.Data["fileType"] = "doc"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".txt" || path.Ext(onlyattachment.FileName) == ".TXT" {
		c.Data["fileType"] = "txt"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".XLS" || path.Ext(onlyattachment.FileName) == ".xls" {
		c.Data["fileType"] = "xls"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".csv" || path.Ext(onlyattachment.FileName) == ".CSV" {
		c.Data["fileType"] = "csv"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".ppt" || path.Ext(onlyattachment.FileName) == ".PPT" {
		c.Data["fileType"] = "ppt"
		c.Data["documentType"] = "presentation"
	} else if path.Ext(onlyattachment.FileName) == ".pdf" || path.Ext(onlyattachment.FileName) == ".PDF" {
		c.Data["fileType"] = "pdf"
		c.Data["documentType"] = "text"
		c.Data["Mode"] = "view"
	}

	c.Data["DocumentServerIP"] = beego.AppConfig.String("documentserverip")
	c.Data["EntryServerIP"] = beego.AppConfig.String("entryserverip")

	c.TplName = "onlyoffice/onlyoffice.tpl"
}

//cms中查阅office
func (c *OnlyController) OfficeView() {
	//设置响应头——没有作用
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	attachment, err := models.GetAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}
	fileext := path.Ext(attachment.FileName)
	product, err := models.GetProd(attachment.ProductId)
	if err != nil {
		beego.Error(err)
	}

	//根据projid取出路径
	proj, err := models.GetProj(product.ProjectId)
	if err != nil {
		beego.Error(err)
	}

	var projurl string
	if proj.ParentIdPath == "" || proj.ParentIdPath == "$#" {
		projurl = "/" + strconv.FormatInt(proj.Id, 10) + "/"
	} else {
		// projurl = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10) + "/"
		projurl = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10) + "/"
	}
	//由proj id取得url
	fileurl, _, err := GetUrlPath(product.ProjectId)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(fileurl + "/" + attachment.FileName)
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	useridstring := strconv.FormatInt(uid, 10)
	var usersessionid string //客户端sesssionid
	if islogin {
		usersessionid = c.Ctx.Input.Cookie("hotqinsessionid")
		//服务端sessionid怎么取出
		// v := c.GetSession("uname")
		// beego.Info(v.(string))
		if e.Enforce(useridstring, projurl, c.Ctx.Request.Method, fileext) || isadmin {
			// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)//这样写下载的文件名称不对
			// c.Redirect(url+"/"+attachment.FileName, 302)
			// c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
			c.Data["FilePath"] = fileurl + "/" + attachment.FileName
			c.Data["Username"] = username
			c.Data["Ip"] = c.Ctx.Input.IP()
			c.Data["role"] = role
			c.Data["IsAdmin"] = isadmin
			c.Data["IsLogin"] = islogin
			c.Data["Uid"] = uid
			c.Data["Mode"] = "edit"
			c.Data["Edit"] = true    //false
			c.Data["Review"] = true  //false
			c.Data["Comment"] = true //false
			c.Data["Download"] = true
			c.Data["Print"] = true
			c.Data["Doc"] = attachment
			c.Data["AttachId"] = idNum
			c.Data["Key"] = strconv.FormatInt(attachment.Updated.UnixNano(), 10)
			c.Data["Sessionid"] = usersessionid

			if path.Ext(attachment.FileName) == ".docx" || path.Ext(attachment.FileName) == ".DOCX" {
				c.Data["fileType"] = "docx"
				c.Data["documentType"] = "text"
			} else if path.Ext(attachment.FileName) == ".wps" || path.Ext(attachment.FileName) == ".WPS" {
				c.Data["fileType"] = "docx"
				c.Data["documentType"] = "text"
			} else if path.Ext(attachment.FileName) == ".XLSX" || path.Ext(attachment.FileName) == ".xlsx" {
				c.Data["fileType"] = "xlsx"
				c.Data["documentType"] = "spreadsheet"
			} else if path.Ext(attachment.FileName) == ".ET" || path.Ext(attachment.FileName) == ".et" {
				c.Data["fileType"] = "xlsx"
				c.Data["documentType"] = "spreadsheet"
			} else if path.Ext(attachment.FileName) == ".pptx" || path.Ext(attachment.FileName) == ".PPTX" {
				c.Data["fileType"] = "pptx"
				c.Data["documentType"] = "presentation"
			} else if path.Ext(attachment.FileName) == ".dps" || path.Ext(attachment.FileName) == ".DPS" {
				c.Data["fileType"] = "pptx"
				c.Data["documentType"] = "presentation"
			} else if path.Ext(attachment.FileName) == ".doc" || path.Ext(attachment.FileName) == ".DOC" {
				c.Data["fileType"] = "doc"
				c.Data["documentType"] = "text"
			} else if path.Ext(attachment.FileName) == ".txt" || path.Ext(attachment.FileName) == ".TXT" {
				c.Data["fileType"] = "txt"
				c.Data["documentType"] = "text"
			} else if path.Ext(attachment.FileName) == ".XLS" || path.Ext(attachment.FileName) == ".xls" {
				c.Data["fileType"] = "xls"
				c.Data["documentType"] = "spreadsheet"
			} else if path.Ext(attachment.FileName) == ".csv" || path.Ext(attachment.FileName) == ".CSV" {
				c.Data["fileType"] = "csv"
				c.Data["documentType"] = "spreadsheet"
			} else if path.Ext(attachment.FileName) == ".ppt" || path.Ext(attachment.FileName) == ".PPT" {
				c.Data["fileType"] = "ppt"
				c.Data["documentType"] = "presentation"
			} else if path.Ext(attachment.FileName) == ".pdf" || path.Ext(attachment.FileName) == ".PDF" {
				c.Data["fileType"] = "pdf"
				c.Data["documentType"] = "text"
				c.Data["Mode"] = "view"
			}

			u := c.Ctx.Input.UserAgent()
			matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
			if err != nil {
				beego.Error(err)
			}
			if matched == true {
				c.Data["Type"] = "mobile"
			} else {
				c.Data["Type"] = "desktop"
			}
			c.TplName = "onlyoffice/officeview.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//协作页面的保存和回调
//关闭浏览器标签后获取最新文档保存到文件夹
func (c *OnlyController) UrltoCallback() {
	// pk1 := c.Ctx.Input.RequestBody
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}

	var callback Callback
	json.Unmarshal(c.Ctx.Input.RequestBody, &callback)

	beego.Info(callback.Status)
	beego.Info(callback.Status)
	beego.Info(callback.Status)
	beego.Info(callback.Status)
	beego.Info(callback.Status)
	beego.Info(callback.Status)
	beego.Info(callback)
	beego.Info(callback)
	beego.Info(callback)

	if callback.Status == 1 || callback.Status == 4 {
		//•	1 - document is being edited,
		//•	4 - document is closed with no changes,
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else if callback.Status == 2 || callback.Status == 6 {
		//•	2 - document is ready for saving
		//•	6 - document is being edited, but the current document state is saved,
		resp, err := http.Get(callback.Url)
		if err != nil {
			beego.Error(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
		}
		defer resp.Body.Close()
		if err != nil {
			beego.Error(err)
		}

		// 原来的代码有bug，不能用append
		// f, err := os.OpenFile("./attachment/onlyoffice/"+onlyattachment.FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
		f, err := os.Create("./attachment/onlyoffice/" + onlyattachment.FileName)
		defer f.Close()

		if err != nil {
			beego.Error(err)
		}

		_, err = f.Write(body)

		if err != nil {
			beego.Error(err)
		} else {
			err = models.UpdateOnlyAttachment(idNum, "")
			if err != nil {
				beego.Error(err)
			}
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	}
}

//cms中返回值
//没改历史版本问题
func (c *OnlyController) OfficeViewCallback() {
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	attachment, err := models.GetAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}

	product, err := models.GetProd(attachment.ProductId)
	if err != nil {
		beego.Error(err)
	}
	//由proj id取得文件路径
	_, diskdirectory, err := GetUrlPath(product.ProjectId)
	if err != nil {
		beego.Error(err)
	}

	var callback Callback
	json.Unmarshal(c.Ctx.Input.RequestBody, &callback)
	//•	1 - document is being edited,
	//•	4 - document is closed with no changes,
	if callback.Status == 1 || callback.Status == 4 {
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
		//•	2 - document is ready for saving
		//•	6 - document is being edited, but the current document state is saved,
	} else if callback.Status == 2 && callback.Notmodified == false {
		//•	2 - document is ready for saving
		resp, err := http.Get(callback.Url)
		if err != nil {
			beego.Error(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
		}
		defer resp.Body.Close()
		if err != nil {
			beego.Error(err)
		}
		// f, err := os.OpenFile("./attachment/onlyoffice/"+onlyattachment.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		f, err := os.Create(diskdirectory + "/" + attachment.FileName)
		if err != nil {
			beego.Error(err)
		}
		defer f.Close()
		_, err = f.Write(body) //这里直接用resp.Body如何？
		// _, err = f.WriteString(str)
		// _, err = io.Copy(body, f)
		if err != nil {
			beego.Error(err)
		} else {
			//更新附件的时间和changesurl
			err = models.UpdateAttachmentTime(idNum)
			if err != nil {
				beego.Error(err)
			}
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
		//3-document saving error has occurred
		//•	7 - error has occurred while force saving the document.
	} else if callback.Status == 3 || callback.Status == 7 {
		//更新附件的时间和changesurl
		err = models.UpdateAttachmentTime(idNum)
		if err != nil {
			beego.Error(err)
		}
		//更新文档更新时间
		// err = models.UpdateProductTime(product.Id)
		// if err != nil {
		// 	beego.Error(err)
		// }
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	}
}

//批量添加一对一模式
//要避免同名覆盖的严重bug！！！！
func (c *OnlyController) AddOnlyAttachment() {
	//取得客户端用户名
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	var filepath, DiskDirectory, Url string
	err := os.MkdirAll("./attachment/onlyoffice/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	DiskDirectory = "./attachment/onlyoffice/"
	Url = "/attachment/onlyoffice/"

	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	if h != nil {
		//保存附件
		//将附件的编号和名称写入数据库
		_, code, title, _, _, _, _ := Record(h.Filename)
		// filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		if code == "" {
			code = title //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		//存入成果数据库
		//如果编号重复，则不写入，只返回Id值。
		//根据id添加成果code, title, label, principal, content string, projectid int64
		prodlabel := c.Input().Get("prodlabel")
		prodprincipal := c.Input().Get("prodprincipal")
		inputdate := c.Input().Get("proddate")
		// beego.Info(inputdate)
		var t1, end time.Time
		// var convdate1, convdate2 string
		const lll = "2006-01-02"
		if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
			t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
			if err != nil {
				beego.Error(err)
			}
			end = t1
		} else { //如果取系统时间，则需要加8小时
			date := time.Now()
			convdate := date.Format(lll)
			date, err = time.Parse(lll, convdate)
			if err != nil {
				beego.Error(err)
			}
			end = date
		}
		// beego.Info(end)
		prodId, err := models.AddDoc(code, title, prodlabel, prodprincipal, end, uid)
		if err != nil {
			beego.Error(err)
		}
		//改名，替换文件名中的#和斜杠
		title = strings.Replace(title, "#", "号", -1)
		title = strings.Replace(title, "/", "-", -1)

		filepath = DiskDirectory + "/" + h.Filename
		attachmentname := h.Filename

		_, _, err2 := models.AddOnlyAttachment(attachmentname, 0, 0, prodId)
		if err2 != nil {
			beego.Error(err2)
		} else {
			//存入文件夹
			//判断文件是否存在，如果不存在就写入
			if _, err := os.Stat(filepath); err != nil {
				// beego.Info(err)
				if os.IsNotExist(err) {
					err = c.SaveToFile("file", filepath) //存文件
					if err != nil {
						beego.Error(err)
					}
					c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": h.Filename, "original": h.Filename, "url": Url + "/" + h.Filename}
					c.ServeJSON()
				}
			} else {
				// beego.Info(err)
				c.Data["json"] = map[string]interface{}{"state": "error"}
				c.ServeJSON()
			}

		}
	}
}

//协作页面下载的文档，采用绝对路径型式
func (c *OnlyController) DownloadDoc() {
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	if strings.Contains(filePath, "?hotqinsessionid=") {
		filePathtemp := strings.Split(filePath, "?")
		filePath = filePathtemp[0]
	}
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
}

//文档管理页面下载文档
func (c *OnlyController) Download() {
	cron.TodayDocSync()
	time.Sleep(2 * time.Second)
	v := c.GetSession("uname")
	if v != nil {
		uname := v.(string)
		beego.Info(uname)
	}
	docid := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	filePath := "attachment/onlyoffice/" + attachments[0].FileName
	c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
}

//文档管理页面下载文档
func (c *OnlyController) Sync() {
	// cron.TodayDocSync()
	// time.Sleep(1 * time.Second)
	v := c.GetSession("uname")
	if v != nil {
		uname := v.(string)
		beego.Info(uname)
	}
	docid := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}

	todayKey := strconv.FormatInt(attachments[0].Updated.UnixNano(), 10)
	url := beego.AppConfig.String("documentserverip") + "/coauthoring/CommandService.ashx"

	postJson := `{"c":"forcesave","key":"` + todayKey + `"}`
	var jsonStr = []byte(postJson)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	c.Data["json"] = "ok"
	c.ServeJSON()
}

//编辑成果信息
func (c *OnlyController) UpdateDoc() {
	id := c.Input().Get("pid")
	code := c.Input().Get("code")
	title := c.Input().Get("title")
	label := c.Input().Get("label")
	principal := c.Input().Get("principal")
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	inputdate := c.Input().Get("proddate")
	var t1, end time.Time
	const lll = "2006-01-02"
	if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
		t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
		if err != nil {
			beego.Error(err)
		}
		end = t1
		// t1 = printtime.Add(+time.Duration(hours) * time.Hour)
	} else { //如果取系统时间，则需要加8小时
		date := time.Now()
		convdate := date.Format(lll)
		date, err = time.Parse(lll, convdate)
		if err != nil {
			beego.Error(err)
		}
		end = date
	}
	//根据id添加成果
	err = models.UpdateDoc(idNum, code, title, label, principal, end)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//删除成果，包含成果里的附件。删除附件用attachment中的
func (c *OnlyController) DeleteDoc() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		attachments, err := models.GetOnlyAttachments(idNum)
		if err != nil {
			beego.Error(err)
		}
		for _, w := range attachments {
			//取得附件的成果id
			attach, err := models.GetOnlyAttachbyId(w.Id)
			if err != nil {
				beego.Error(err)
			}
			path := "./attachment/onlyoffice/" + attach.FileName
			//删除附件
			err = os.Remove(path)
			if err != nil {
				beego.Error(err)
			}
			//删除附件数据表
			err = models.DeleteOnlyAttachment(w.Id)
			if err != nil {
				beego.Error(err)
			}
		}
		err = models.DeleteDoc(idNum) //删除成果数据表
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//onlyoffice权限管理
//添加用户和角色的权限
//先删除这个文档id下所有permission，再添加新的。
func (c *OnlyController) Addpermission() {
	// roleids := c.GetString("roleids")
	// rolearray := strings.Split(roleids, ",")
	// userids := c.GetString("userids")
	// userarray := strings.Split(userids, ",")
	ids := c.GetString("ids")
	// beego.Info(ids)
	tt := []byte(ids)
	var rolepermission []Rolepermission
	json.Unmarshal(tt, &rolepermission)
	// beego.Info(rolepermission)
	docid := c.GetString("docid")
	//id转成64位
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件——这里只取第一个
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	suf := ".*"
	var success bool
	//先删除这个文档所有的permission
	e.RemoveFilteredPolicy(1, "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10))
	//再添加permission
	for _, v1 := range rolepermission {
		// beego.Info(v1.Id)
		if v1.Rolenumber != "" { //存储角色id
			success = e.AddPolicy("role_"+strconv.FormatInt(v1.Id, 10), "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10), v1.Permission, suf)
		} else { //存储用户id
			success = e.AddPolicy(strconv.FormatInt(v1.Id, 10), "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10), v1.Permission, suf)
		}
		//这里应该用AddPermissionForUser()，来自casbin\rbac_api.go
	}
	if success == true {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "wrong"
	}
	c.ServeJSON()
}

//查询一个文档，哪些用户和角色拥有什么样的权限
//用casbin的内置方法，不应该用查询数据库方法
func (c *OnlyController) Getpermission() {
	docid := c.GetString("docid")
	//id转成64位
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	// var users []beegoormadapter.CasbinRule
	rolepermission := make([]Rolepermission, 0)
	for _, w := range attachments {
		// o := orm.NewOrm()
		// qs := o.QueryTable("casbin_rule")
		// _, err = qs.Filter("PType", "p").Filter("v1", "/onlyoffice/"+strconv.FormatInt(w.Id, 10)).All(&users)
		// if err != nil {
		// 	beego.Error(err)
		// }
		users := e.GetFilteredPolicy(1, "/onlyoffice/"+strconv.FormatInt(w.Id, 10))
		// beego.Info(users)
		for _, v := range users {
			rolepermission1 := make([]Rolepermission, 1)
			if strings.Contains(v[0], "role_") { //是角色
				// beego.Info(v.V0)
				roleid := strings.Replace(v[0], "role_", "", -1)
				//id转成64位
				roleidNum, err := strconv.ParseInt(roleid, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				// beego.Info(roleidNum)
				role := models.GetRoleByRoleId(roleidNum)
				// beego.Info(role)
				rolepermission1[0].Id = roleidNum
				rolepermission1[0].Name = role.Rolename
				rolepermission1[0].Rolenumber = role.Rolenumber
				rolepermission1[0].Permission = v[2]
				// rolepermission = append(rolepermission, rolepermission1...)
			} else { //是用户
				// rolepermission1 := make([]Rolepermission, 1)
				//id转成64位
				uidNum, err := strconv.ParseInt(v[0], 10, 64)
				if err != nil {
					beego.Error(err)
				}
				user := models.GetUserByUserId(uidNum)
				rolepermission1[0].Id = uidNum
				rolepermission1[0].Name = user.Nickname
				rolepermission1[0].Permission = v[2]
			}
			rolepermission = append(rolepermission, rolepermission1...)
		}
	}
	c.Data["json"] = rolepermission
	c.ServeJSON()
}

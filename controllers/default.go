package controllers

import (
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var id string
	navid1 := beego.AppConfig.String("navigationid1")
	navid2 := beego.AppConfig.String("navigationid2")
	navid3 := beego.AppConfig.String("navigationid3")
	navid4 := beego.AppConfig.String("navigationid4")
	navid5 := beego.AppConfig.String("navigationid5")
	navid6 := beego.AppConfig.String("navigationid6")
	navid7 := beego.AppConfig.String("navigationid7")
	navid8 := beego.AppConfig.String("navigationid8")
	navid9 := beego.AppConfig.String("navigationid9")
	index := beego.AppConfig.String("defaultindex")
	beego.Info(index)
	switch index {
	case "IsNav1":
		id = navid1
		c.Redirect("/project/"+id, 301)
	case "IsNav2":
		id = navid2
		// beego.Info(id)
		c.Redirect("/project/"+id, 301)
	case "IsNav3":
		id = navid3
		c.Redirect("/project/"+id, 301)
	case "IsNav4":
		id = navid4
		c.Redirect("/project/"+id, 301)
	case "IsNav5":
		id = navid5
		c.Redirect("/project/"+id, 301)
	case "IsNav6":
		id = navid6
		c.Redirect("/project/"+id, 301)
	case "IsNav7":
		id = navid7
		c.Redirect("/project/"+id, 301)
	case "IsNav8":
		id = navid8
		c.Redirect("/project/"+id, 301)
	case "IsNav9":
		id = navid9
		c.Redirect("/project/"+id, 301)
	case "IsProject":
		c.Redirect("/project", 301)
	case "IsOnlyOffice":
		c.Redirect("/onlyoffice", 301)
	case "IsDesignGant", "IsConstructGant":
		c.Redirect("/projectgant", 301)
	case "IsLogin":
		c.Redirect("/login", 301)
		//登录后默认跳转……
	case "IsIndex":
		c.Redirect("/index", 301)
	default:
		c.Redirect("/index", 301)
	}
}

//文档
func (c *MainController) Getecmsdoc() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "ecmsdoc.tpl"
}

func (c *MainController) Getmeritmsdoc() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "meritmsdoc.tpl"
}

func (c *MainController) Gethydrowsdoc() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "hydrowsdoc.tpl"
}

//api
func (c *MainController) Getecmsapi() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "ecmsapi.tpl"
}

func (c *MainController) Getmeritmsapi() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "meritmsapi.tpl"
}

func (c *MainController) Test() {
	// beego.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "test1.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "test.tpl"
	}
}

func (c *MainController) Slide() {
	c.TplName = "slide.tpl"
}

func (c *MainController) Postdata() {
	const lll = "2006-01-02"
	date := time.Now()
	convdate := string(date.Format(lll))

	f, _, err := c.GetFile("uploadfile")
	// beego.Info(h) //这里 filename是路径，所以不能以filename作为保存的文件名。坑！！
	defer f.Close()

	if err != nil {
		beego.Error(err)
	} else {
		c.SaveToFile("uploadfile", "./static/upload/"+convdate+".data") // 保存位置在 static/upload, 没有文件夹要先创建
		c.Ctx.WriteString("ok")
	}
}

func (c *MainController) Register() {
	// flash := beego.NewFlash()
	token := c.Input().Get("token")
	//是否重复提交
	if c.IsSubmitAgain(token) {
		c.Redirect("/registerpage", 302)
		return
	}
}

func (c *MainController) IsSubmitAgain(token string) bool {
	cotoken := c.Ctx.GetCookie("token")
	if token == "" || len(token) == 0 || token != cotoken || strings.Compare(cotoken, token) != 0 {
		return true
	}
	return false
}

func (c *MainController) Pdf() {
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	if uid == 0 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}

	var Url string
	var pNum int64
	id := c.Input().Get("id")
	// beego.Info(id)
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	p := c.Input().Get("p")
	// beego.Info(p)

	var prods []*models.Product
	var projid int64
	if p == "" { //id是附件或成果id
		attach, err := models.GetAttachbyId(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得成果——进行排序
		prod, err := models.GetProd(attach.ProductId)
		if err != nil {
			beego.Error(err)
		}
		//由侧栏id取得所有成果和所有成果的附件pdf
		prods, err = models.GetProducts(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(prod.ProjectId)
		//由proj id取得url
		// Url, _, err = GetUrlPath(prod.ProjectId)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// beego.Info(Url)
		projid = prod.ProjectId
	} else { //id是侧栏目录id
		//由侧栏id取得所有成果和所有成果的附件pdf——进行排序
		prods, err = models.GetProducts(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由proj id取得url
		Url, _, err = GetUrlPath(idNum)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(Url)
		// projid = idNum
		//id转成64为
		pNum, err = strconv.ParseInt(p, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	// var count int64
	Attachments := make([]*models.Attachment, 0)
	for _, v := range prods {
		//根据成果id取得所有附件数量
		// count1, err := models.GetAttachmentsCount(v.Id)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// count = count + count1
		Attachments1, err := models.GetAttachments(v.Id) //只返回Id?没意义，全部返回
		if err != nil {
			beego.Error(err)
		}
		for _, v := range Attachments1 {
			if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				Attachments = append(Attachments, Attachments1...)
			}
		}
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	// link := make([]AttachmentLink, 0)
	// for _, v := range Attachments {
	// 	if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
	// 		linkarr := make([]AttachmentLink, 1)
	// 		linkarr[0].Id = v.Id
	// 		linkarr[0].Title = v.FileName
	// 		linkarr[0].FileSize = v.FileSize
	// 		linkarr[0].Downloads = v.Downloads
	// 		linkarr[0].Created = v.Created
	// 		linkarr[0].Updated = v.Updated
	// 		linkarr[0].Link = Url + "/" + v.FileName
	// 		link = append(link, linkarr...)
	// 	}
	// }
	// c.Data["json"] = link
	// pdfs, err := models.GetAllPdfs(idNum, false)
	// if err != nil {
	// 	beego.Error(err)
	// }
	count := len(Attachments)
	count1 := strconv.Itoa(count)
	count2, err := strconv.ParseInt(count1, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// cnt, err := o.QueryTable("user").Count()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// sets this.Data["paginator"] with the current offset (from the url query param)
	postsPerPage := 1
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
	// fmt.Println(*c.Ctx)
	// beego.Info(c.Ctx)
	// beego.Info(paginator.Offset())   0
	// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
	// beego.Info(p.Offset())   0
	// fetch the next 20 posts
	//根据paginator.Offset()序列，取得附件
	// Attachments := make([]*models.Attachment, 0)
	// for _, v := range prod {
	// 	//根据成果id取得所有附件id
	// 	Attachments1, err := models.GetAttachments(v.Id) //只返回Id?没意义，全部返回
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	Attachments = append(Attachments, Attachments1...)
	// }
	//查询当前附件id所在位置
	// var p1, PdfLink string
	// for i, v := range Attachments {
	// 	if v.Id == idNum {
	// 		p1 = strconv.Itoa(i)
	// 		PdfLink = Url + "/" + v.FileName
	// 		break
	// 	}
	// }
	// pdfs, err = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//这里根据上面取得的分页topics，取出这页的成果对应的所有项目
	// slice1 := make([]models.Category, 0)
	// for _, v := range topics {
	// 	tid := strconv.FormatInt(v.Id, 10)
	// 	category, err := models.GetTopicProj(tid)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// beego.Info(category.Title)
	// 	aa := make([]models.Category, 1)
	// 	aa[0].Author = category.Title//这句出错，不知何故runtime error: invalid memory address or nil pointer dereference
	// 	slice1 = append(slice1, aa...)
	// }
	// c.Data["Categories"] = slice1
	c.Data["paginator"] = paginator
	if p == "" {
		var p1 string
		for i, v := range Attachments {
			if v.Id == idNum {
				p1 = strconv.Itoa(i + 1)
				// PdfLink = Url + "/" + v.FileName
				break
			}
		}
		c.Redirect("/pdf?p="+p1+"&id="+strconv.FormatInt(projid, 10), 302)
	} else {
		PdfLink := Url + "/" + Attachments[pNum-1].FileName
		// beego.Info(PdfLink)
		c.Data["PdfLink"] = PdfLink
		c.TplName = "web/viewer.html"
	}
	// logs := logs.NewLogger(1000)
	// logs.SetLogger("file", `{"filename":"log/test.log"}`)
	// logs.EnableFuncCallDepth(true)
	// logs.Info(c.Ctx.Input.IP() + " " + "ListAllTopic")
	// logs.Close()
	// count, _ := models.M("logoperation").Alias(`op`).Field(`count(op.id) as count`).Where(where).Count()
	// if count > 0 {
	// 	pagesize := 10
	// 	p := tools.NewPaginator(this.Ctx.Request, pagesize, count)
	// 	log, _ := models.M("logoperation").Alias(`op`).Where(where).Limit(strconv.Itoa(p.Offset()), strconv.Itoa(pagesize)).Order(`op.id desc`).Select()
	// 	this.Data["data"] = log
	// 	this.Data["paginator"] = p
	// }
}

// @Title dowload wx pdf
// @Description get wx pdf by id
// @Param id path string  true "The id of pdf"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /wxpdf/:id [get]
func (c *MainController) WxPdf() {
	// var openID string
	// openid := c.GetSession("openID")
	// beego.Info(openid)
	// if openid == nil {

	// } else {
	// 	openID = openid.(string)
	// }
	openID := c.GetSession("openID")
	beego.Info(openID)
	if openID != nil {
		user, err := models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		} else {
			useridstring := strconv.FormatInt(user.Id, 10)
			// 判断用户是否具有权限。
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

			product, err := models.GetProd(attachment.ProductId)
			if err != nil {
				beego.Error(err)
			}
			//根据projid取出路径
			proj, err := models.GetProj(product.ProjectId)
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(err.Error())
			}
			var projurl string
			if proj.ParentIdPath == "" || proj.ParentIdPath == "$#" {
				projurl = "/" + strconv.FormatInt(proj.Id, 10) + "/"
			} else {
				projurl = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10) + "/"
			}
			//由proj id取得url
			fileurl, _, err := GetUrlPath(product.ProjectId)
			if err != nil {
				beego.Error(err)
			}
			fileext := path.Ext(attachment.FileName)
			if e.Enforce(useridstring, projurl, c.Ctx.Request.Method, fileext) {
				c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
			}
		}
	} else {
		c.Data["json"] = "未查到openID"
		c.ServeJSON()
	}
}

// @Title dowload wx standardpdf
// @Description get wx standardpdf by id
// @Param id path string  true "The id of standardpdf"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /wxstandardpdf/:id [get]
func (c *MainController) WxStandardPdf() {
	// id := c.Input().Get("id")
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//根据id取得规范的路径
	standard, err := models.GetStandard(idNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(standard.Route)
	fileurl := strings.Replace(standard.Route, "/attachment/", "attachment/", -1)
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, standard.Route)
	c.Ctx.Output.Download(fileurl)
}

//升级数据库
func (c *MainController) UpdateDatabase() {
	beego.Info("ok")
	err1, err2, err3, err4, err5, err6, err7 := models.UpdateDatabase()
	if err1 != nil {
		beego.Error(err1)
	}
	if err2 != nil {
		beego.Error(err2)
	}
	if err3 != nil {
		beego.Error(err3)
	}
	if err4 != nil {
		beego.Error(err4)
	}
	if err5 != nil {
		beego.Error(err5)
	}
	if err6 != nil {
		beego.Error(err6)
	}
	if err7 != nil {
		beego.Error(err7)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//删除数据表和字段测试
func (c *MainController) ModifyDatabase() {
	err := models.ModifyDatabase()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = err
	c.ServeJSON()
}

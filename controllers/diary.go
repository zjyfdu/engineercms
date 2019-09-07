package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	// "encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	// "net"
	// "net/http"
	// "net/url"
	// "path"
	"strconv"
	"strings"
	// "time"
)

// CMSWXDIARY API
type DiaryController struct {
	beego.Controller
}

// @Title post wx diary by catalogId
// @Description post diary by projectid
// @Param title query string  true "The title of diary"
// @Param diarydate query string  true "The diarydate of diary"
// @Param diaryactivity query string  true "The diaryactivity of diary"
// @Param diaryweather query string  true "The diaryweather of diary"
// @Param content query string  true "The content of diary"
// @Param skey query string  true "The skey of user"
// @Success 200 {object} models.AddDiary
// @Failure 400 Invalid page supplied
// @Failure 404 Diary not found
// @router /addwxdiary [post]
//向设代日记id下添加微信小程序文章_珠三角设代plus用_这个是文字图片分开方式，用下面这个
func (c *DiaryController) AddWxDiary() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	beego.Info(openID)
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		} else {
			beego.Info(user)
			pid := beego.AppConfig.String("wxdiaryprojectid") //"26159"
			title := c.Input().Get("title")
			diarydate := c.Input().Get("diarydate")

			array := strings.Split(diarydate, "-")
			//当月天数
			const base_format = "2006-01-02"
			year := array[0]
			month := array[1]
			if len(month) == 1 {
				month = "0" + month
			}
			day := array[2]
			if len(day) == 1 {
				day = "0" + day
			}
			// diarydate2, err := time.Parse(base_format, year+"-"+month+"-"+day)
			// if err != nil {
			// 	beego.Error(err)
			// }
			diarydate2 := year + "-" + month + "-" + day
			// beego.Info(diarydate2)
			diaryactivity := c.Input().Get("diaryactivity")
			diaryweather := c.Input().Get("diaryweather")

			content := c.Input().Get("content")
			content = "<p style='font-size: 16px;'>分部：" + diaryactivity + "；</p><p style='font-size: 16px;'>天气：" + diaryweather + "；</p><p style='font-size: 16px;'>记录：" + user.Nickname + "；</p>" + content //<span style="font-size: 18px;">这个字体到底是多大才好看</span>

			//id转成64为
			pidNum, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//添加日志
			aid, err := models.AddDiary(title, content, diarydate2, pidNum, user.Id)
			if err != nil {
				beego.Error(err)
				c.Data["json"] = map[string]interface{}{"status": 0, "info": "ERR", "id": aid}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
				c.ServeJSON()
			}
		}
	} else {
		return
	}
	// //根据pid查出项目id
	// proj, err := models.GetProj(pidNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
	// parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	// patharray := strings.Split(parentidpath1, "-")
	// topprojectid, err := strconv.ParseInt(patharray[0], 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// code := time.Now().Format("2006-01-02 15:04")
	// code = strings.Replace(code, "-", "", -1)
	// code = strings.Replace(code, " ", "", -1)
	// code = strings.Replace(code, ":", "", -1)
	// //根据项目id添加成果code, title, label, principal, content string, projectid int64
	// Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
	// if err != nil {
	// 	beego.Error(err)
	// }
}

// @Title get wx diaries list
// @Description get diaries by page
// @Param page query string  true "The page for diaries list"
// @Param skey query string  true "The skey for user"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxdiaries [get]
//小程序取得所有文章列表，分页_珠三角设代用
func (c *DiaryController) GetWxdiaries() {
	// id := c.Ctx.Input.Param(":id")
	id := beego.AppConfig.String("wxdiaryprojectid") //"26159" //25002珠三角设代日记id26159
	// wxsite := beego.AppConfig.String("wxreqeustsite")
	limit := "10"
	// limit1, err := strconv.ParseInt(limit, 10, 64)
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("page")
	// page1, err := strconv.ParseInt(page, 10, 64)
	page1, err := strconv.Atoi(page)
	if err != nil {
		beego.Error(err)
	}

	var idNum int64
	//id转成64为
	idNum, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// var offset int64
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	// diaries, err := models.GetWxDiaries(idNum, limit1, offset)
	diaries, err := models.GetWxDiaries2(idNum, limit1, offset)
	if err != nil {
		beego.Error(err)
	}
	beego.Info(diaries)

	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "diaries": diaries}
	c.ServeJSON()
}

// @Title get wx diary by diaryId
// @Description get diary by diaryid
// @Param id path string  true "The id of diary"
// @Param skey path string  true "The skey of user"
// @Success 200 {object} models.GetDiary
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxdiary/:id [get]
//根据id查看一篇微信文章
func (c *DiaryController) GetWxDiary() {
	var err error
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	wxsite := beego.AppConfig.String("wxreqeustsite")

	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(idNum)
	Diary, err := models.GetDiary(idNum)
	if err != nil {
		beego.Error(err)
	}

	content := strings.Replace(Diary.Content, "/attachment/", wxsite+"/attachment/", -1)
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	// hours := 8
	// const lll = "2006-01-02 15:04"
	// diarytime := Diary.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)
	wxDiary := &models.Diary{
		Id:        Diary.Id,
		Title:     Diary.Title,
		Diarydate: Diary.Diarydate,
		Content:   content, //Diary.Content,
		// LeassonType: 1,
		Views:   Diary.Views,
		Created: Diary.Created,
		// Updated:     diarytime,
	}
	c.Data["json"] = wxDiary
	c.ServeJSON()
}

// @Title post wx diary by diaryid
// @Description post diary by diaryid
// @Param id query string true "The id of diary"
// @Param title query string true "The title of diary"
// @Param content query string true "The content of diary"
// @Success 200 {object} models.AddDiary
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /updatewxdiary [post]
//编辑设代日记id下微信小程序文章_珠三角设代plus用_editor方式
func (c *DiaryController) UpdateWxDiary() {
	// pid := beego.AppConfig.String("wxcatalogid") //"26159"
	//hotqinsessionid携带过来后，用下面的方法获取用户登录存储在服务端的session
	openid := c.GetSession("openID")
	if openid == nil {
		return
	}

	id := c.Input().Get("id")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	//将content中的http://ip/去掉
	wxsite := beego.AppConfig.String("wxreqeustsite")
	content = strings.Replace(content, wxsite, "", -1)
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//取得文章
	_, err = models.GetDiary(idNum)
	if err != nil {
		beego.Error(err)
	} else {
		//更新文章
		err = models.UpdateDiary(idNum, title, content)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERR", "id": id}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": id}
			c.ServeJSON()
		}
	}
}

// @Title post wx diary by diaryId
// @Description post diary by catalogid
// @Param id query string  true "The id of diary"
// @Success 200 {object} models.Adddiary
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /deletewxdiary [post]
//根据id删除_没删除文章中的图片
func (c *DiaryController) DeleteWxDiary() {
	var openID string
	openid := c.GetSession("openID")
	beego.Info(openid.(string))
	if openid == nil {
		c.Data["json"] = "没有注册，未查到openid"
		c.ServeJSON()
	} else {
		openID = openid.(string)
		user, err := models.GetUserByOpenID(openID)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = "未查到openid对应的用户"
			c.ServeJSON()
		} else {
			//判断是否具备admin角色
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			uid := strconv.FormatInt(user.Id, 10)
			roleid := strconv.FormatInt(role.Id, 10)
			if e.HasRoleForUser(uid, "role_"+roleid) {
				id := c.Input().Get("id")
				//id转成64为
				idNum, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				err = models.DeleteDiary(idNum)
				if err != nil {
					beego.Error(err)
					c.Data["json"] = "delete wrong"
					c.ServeJSON()
				} else {
					c.Data["json"] = "ok"
					c.ServeJSON()
				}
			} else {
				c.Data["json"] = "不是管理员"
				c.ServeJSON()
			}
		}
	}
}

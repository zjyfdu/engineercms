package controllers

import (
	// "bytes"

	"github.com/astaxie/beego"

	// "image/png"

	"os"
	"path"

	// "hydrocms/models"

	"strconv"
	"time"

	"github.com/3xxx/engineercms/models"
)

// CMSWX froala API
type FroalaController struct {
	beego.Controller
}

// @Title post wx artile img by catalogId
// @Description post article img by catalogid
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwximg [post]
//微信wx添加文章里的图片上传_独立上传图片模式
func (c *FroalaController) UploadWxImg() {
	//解析表单
	pid := beego.AppConfig.String("wxcatalogid") //"26159" //c.Input().Get("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx artile img by catalogId
// @Description post article img by catalogid
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwxeditorimg [post]
//微信wx添加文章里的图片上传_小程序富文本里的上传图片
func (c *FroalaController) UploadWxEditorImg() {
	//解析表单
	pid := beego.AppConfig.String("wxcatalogid") //"26159" //c.Input().Get("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx artile img by catalogId
// @Description post article img by catalogid
// @Param id query string  true "The id of project"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwximgs/:id [post]
//微信wx添加文章里的图片上传——这个没有用上，但这个id更通用
func (c *FroalaController) UploadWxImgs() {
	//解析表单
	pid := c.Ctx.Input.Param(":id")
	// pid := beego.AppConfig.String("wxcatalogid") //"26159" //c.Input().Get("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx user avatar
// @Description post user avatar
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwxavatar [post]
//微信wx添加用户头像上传
func (c *FroalaController) UploadWxAvatar() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	err = os.MkdirAll("./static/avatar/", 0777)                          //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = "./static/avatar/" + newname
		Url := "/static/avatar/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "photo": "", "title": "", "original": ""}
			c.ServeJSON()
		} else {
			filesize, _ = FileSize(path)
			filesize = filesize / 1000.0
			_, err = models.AddUserAvator(user.Id, Url+newname)
			if err != nil {
				beego.Error(err)
			}
			wxsite := beego.AppConfig.String("wxreqeustsite")
			// c.Data["json"] = map[string]interface{}{"errNo": 1, "msg": "success", "photo": wxsite + Url + newname, "title": newname, "original": newname}
			// c.ServeJSON()
			c.Ctx.WriteString(wxsite + Url + newname)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"errNo": 0, "state": "ERROR", "photo": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

//添加wiki里的图片上传
func (c *FroalaController) UploadWikiImg() {
	//保存上传的图片
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	var filesize int64
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/wiki/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	path1 := "./attachment/wiki/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
	Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
	err = c.SaveToFile("file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		beego.Error(err)
	}
	filesize, _ = FileSize(path1)
	filesize = filesize / 1000.0
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
	c.ServeJSON()
}

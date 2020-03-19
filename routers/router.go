// @APIVersion 1.0.0
// @Title EngineerCMS API
// @Description ECMS has every tool to get any job done, so codename for the new ECMS APIs.
// @Contact 504284@qq.com
package routers

import (
	"github.com/3xxx/engineercms/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//运行跨域请求
	//在http请求的响应流头部加上如下信息
	//rw.Header().Set("Access-Control-Allow-Origin", "*")
	// beego.InsertFilter("?=/article/*", beego.BeforeRouter, FilterFunc)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Router("/test", &controllers.MainController{}, "*:Test")

	//升级数据库
	beego.Router("/updatedatabase", &controllers.MainController{}, "*:UpdateDatabase")
	//删除数据表和字段测试
	beego.Router("/modifydatabase", &controllers.MainController{}, "*:ModifyDatabase")

	beego.Router("/url-to-callback", &controllers.OnlyController{}, "*:UrltoCallback")

	beego.Router("/onlyoffice", &controllers.OnlyController{}, "get:Get")
	//table获取所有数据给上面界面使用
	beego.Router("/onlyoffice/data", &controllers.OnlyController{}, "*:GetData")
	//添加一个文档
	beego.Router("/onlyoffice/addattachment", &controllers.OnlyController{}, "post:AddOnlyAttachment")
	//在onlyoffice中打开文档协作
	beego.Router("/onlyoffice/:id:string", &controllers.OnlyController{}, "*:OnlyOffice")

	//删除
	beego.Router("/onlyoffice/deletedoc", &controllers.OnlyController{}, "*:DeleteDoc")
	//修改
	beego.Router("/onlyoffice/updatedoc", &controllers.OnlyController{}, "*:UpdateDoc")
	//onlyoffice页面下载doc
	beego.Router("/attachment/onlyoffice/*", &controllers.OnlyController{}, "get:DownloadDoc")
	//文档管理页面下载doc
	beego.Router("/onlyoffice/download/:id:string", &controllers.OnlyController{}, "get:Download")
	//文档管理页面下载doc
	beego.Router("/onlyoffice/sync/:id:string", &controllers.OnlyController{}, "get:Sync")
	// beego.Router("/onlyoffice/changes", &controllers.OnlyController{}, "post:ChangesUrl")
	//*****onlyoffice document权限
	//添加用户和角色权限
	beego.Router("/onlyoffice/addpermission", &controllers.OnlyController{}, "post:Addpermission")
	//取得文档的用户和角色权限列表
	beego.Router("/onlyoffice/getpermission", &controllers.OnlyController{}, "get:Getpermission")

	beego.Router("/role/test", &controllers.RoleController{}, "*:Test")
	beego.Router("/1/slide", &controllers.MainController{}, "*:Slide")
	beego.Router("/postdata", &controllers.MainController{}, "*:Postdata")
	//文档
	beego.Router("/doc/ecms", &controllers.MainController{}, "get:Getecmsdoc")
	beego.Router("/doc/meritms", &controllers.MainController{}, "get:Getmeritmsdoc")
	beego.Router("/doc/hydrows", &controllers.MainController{}, "get:Gethydrowsdoc")

	//api接口
	beego.Router("/api/ecms", &controllers.MainController{}, "get:Getecmsapi")
	beego.Router("/api/meritms", &controllers.MainController{}, "get:Getmeritmsapi")
	//根据app.conf里的设置，显示首页
	beego.Router("/", &controllers.MainController{}, "get:Get")
	//显示首页
	beego.Router("/index", &controllers.IndexController{}, "*:GetIndex")
	//首页放到onlyoffice
	// beego.Router("/", &controllers.OnlyController{}, "get:Get")
	beego.Router("/pdf", &controllers.MainController{}, "*:Pdf")
	//显示右侧页面框架
	beego.Router("/index/user", &controllers.IndexController{}, "*:GetUser")
	//这里显示用户查看主人日程
	beego.Router("/calendar", &controllers.IndexController{}, "*:Calendar")

	//用户修改自己密码
	beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")
	//用户产看自己的table中数据填充
	beego.Router("/usermyself", &controllers.UserController{}, "get:Usermyself")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	//页面登录提交用户名和密码
	beego.Router("/post", &controllers.LoginController{}, "post:Post")
	//弹框登录提交用户名和密码
	beego.Router("/loginpost", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")
	beego.Router("/roleerr", &controllers.UserController{}, "*:Roleerr") //显示权限不够

	//项目列表界面
	beego.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	beego.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")

	//侧栏懒加载下级
	beego.Router("/project/getprojcate", &controllers.ProjController{}, "*:GetProjCate")
	//添加项目，应该是project/addproj,delproj,updateproj
	beego.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")
	//添加项目，根据项目模板
	beego.Router("/project/addprojtemplet", &controllers.ProjController{}, "*:AddProjTemplet")

	//修改项目
	beego.Router("/project/updateproject", &controllers.ProjController{}, "*:UpdateProject")
	//删除项目
	beego.Router("/project/deleteproject", &controllers.ProjController{}, "*:DeleteProject")
	//项目时间轴
	beego.Router("/project/:id([0-9]+)/gettimeline", &controllers.ProjController{}, "get:ProjectTimeline")
	beego.Router("/project/:id([0-9]+)/timeline", &controllers.ProjController{}, "get:Timeline")

	//根据项目id进入一个具体项目的侧栏
	beego.Router("/project/:id([0-9]+)", &controllers.ProjController{}, "*:GetProject")
	//进入项目日历
	beego.Router("/project/:id([0-9]+)/getcalendar", &controllers.ProjController{}, "*:GetCalendar")
	//取得日历数据
	beego.Router("/project/:id([0-9]+)/calendar", &controllers.ProjController{}, "*:Calendar")
	//添加日历
	beego.Router("/project/:id([0-9]+)/calendar/addcalendar", &controllers.ProjController{}, "*:AddCalendar")
	//修改
	beego.Router("/project/calendar/updatecalendar", &controllers.ProjController{}, "*:UpdateCalendar")
	//删除
	beego.Router("/project/calendar/deletecalendar", &controllers.ProjController{}, "*:DeleteCalendar")
	//拖曳事件
	beego.Router("/project/calendar/dropcalendar", &controllers.ProjController{}, "*:DropCalendar")
	//resize事件
	beego.Router("/project/calendar/resizecalendar", &controllers.ProjController{}, "*:ResizeCalendar")
	//日历中传照片
	beego.Router("/project/calendar/uploadimage", &controllers.ProjController{}, "*:UploadImage")

	//点击侧栏，根据id返回json数据给导航条
	beego.Router("/project/navbar/:id:string", &controllers.ProjController{}, "*:GetProjNav")

	//根据项目侧栏id显示这个id下的成果界面——作废，用上面GetProject界面
	beego.Router("/project/:id:string/:id:string", &controllers.ProdController{}, "*:GetProjProd")

	//给上面那个页面提供table所用的json数据
	beego.Router("/project/products/:id:string", &controllers.ProdController{}, "*:GetProducts")
	//点击项目名称，显示这个项目下所有成果
	beego.Router("/project/allproducts/:id:string", &controllers.ProjController{}, "*:GetProjProducts")
	//这个对应上面的页面进行表格内数据填充，用的是prodcontrollers中的方法
	beego.Router("/project/products/all/:id:string", &controllers.ProdController{}, "*:GetProjProducts")
	//给上面那个页面提供项目同步ip的json数据
	beego.Router("/project/synchproducts/:id:string", &controllers.ProdController{}, "*:GetsynchProducts")

	//对外提供同步成果接口json数据
	beego.Router("/project/providesynchproducts", &controllers.ProdController{}, "*:ProvidesynchProducts")

	//向项目侧栏id下添加成果_这个没用到，用下面addattachment
	// beego.Router("/project/product/addproduct", &controllers.ProdController{}, "post:AddProduct")
	//编辑成果信息
	beego.Router("/project/product/updateproduct", &controllers.ProdController{}, "post:UpdateProduct")
	//删除成果
	beego.Router("/project/product/deleteproduct", &controllers.ProdController{}, "post:DeleteProduct")
	//取得成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/attachment/:id:string", &controllers.AttachController{}, "*:GetAttachments")
	//取得同步成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/synchattachment", &controllers.AttachController{}, "*:GetsynchAttachments")
	//提供接口：同步成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/providesynchattachment", &controllers.AttachController{}, "*:ProvideAttachments")

	//取得成果中所有附件中——包含pdf，非pdf，文章
	beego.Router("/project/product/allattachments/:id:string", &controllers.AttachController{}, "*:GetAllAttachments")

	//取得成果中所有附件的pdf列表
	beego.Router("/project/product/pdf/:id:string", &controllers.AttachController{}, "*:GetPdfs")
	//取得同步成果中所有pdf列表
	beego.Router("/project/product/synchpdf", &controllers.AttachController{}, "*:GetsynchPdfs")
	//提供同步成果中所有pdf列表
	beego.Router("/project/product/providesynchpdf", &controllers.AttachController{}, "*:ProvidePdfs")

	//向成果里添加附件：批量一对一模式
	beego.Router("/project/product/addattachment", &controllers.AttachController{}, "post:AddAttachment")
	//dwg写入服务器
	beego.Router("/project/product/savedwgfile", &controllers.AttachController{}, "post:SaveDwgfile")
	//新建dwg文件
	beego.Router("/project/product/newdwg", &controllers.AttachController{}, "post:NewDwg")

	//向成果里添加附件：多附件模式
	beego.Router("/project/product/addattachment2", &controllers.AttachController{}, "post:AddAttachment2")
	//编辑附件列表：向成果里追加附件：多附件模式
	beego.Router("/project/product/updateattachment", &controllers.AttachController{}, "post:UpdateAttachment")
	//删除附件列表中的附件
	beego.Router("/project/product/deleteattachment", &controllers.AttachController{}, "post:DeleteAttachment")

	//附件下载"/attachment/*", &controllers.AttachController{}
	// beego.InsertFilter("/attachment/*", beego.BeforeRouter, controllers.ImageFilter)
	//根据附件绝对地址下载
	beego.Router("/attachment/*", &controllers.AttachController{}, "get:Attachment")
	//根据附件id号，判断权限下载
	beego.Router("/downloadattachment", &controllers.AttachController{}, "get:DownloadAttachment")

	//上面用attachment.ImageFilter是不行的，必须是package.func
	//首页轮播图片的权限
	beego.Router("/attachment/carousel/*.*", &controllers.AttachController{}, "get:GetCarousel")

}

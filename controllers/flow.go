package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	// "strings"
	// "testing"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/js-ojus/flow"
	// "github.com/3xxx/meritms/models"
	// _ "github.com/mattn/go-sqlite3"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	"log"
	"strconv"
	// "time"
)

// Flowtest API
type FlowController struct {
	beego.Controller
}

func (c *FlowController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "flow.tpl"
}

// var db *sql.DB

// func init() {
// 	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
// 	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	// flow.RegisterDB(tdb)
// 	if tdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db = tdb
// 	tx, _ := db.Begin()
// 	db.Close()
// }

// @Title show wf page
// @Description show workflow page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /workflow [get]
// 页面
func (c *FlowController) WorkFlow() {
	c.TplName = "index.html"
}

// @Title post wf doctype...
// @Description post workflowdoctype..
// @Param name query string  true "The name of doctype"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtype [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowType() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	beego.Info(name)
	jsoninfo := c.GetString("name1") //获取formdata
	beego.Info(jsoninfo)
	//定义流程类型
	_, err := flow.DocTypes.New(tx, name) //"图纸设计流程"
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

//后端分页的数据结构
type doctypelist struct {
	Doctype []*flow.DocType `json:"doctypes"`
	Page    int64           `json:"page"`
	Total   int             `json:"total"` //string或int64都行！
}

// @Title get wf doctypelist...
// @Description get workflowdoctype..
// @Param page query string false "The page of doctype"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtypelist [get]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowTypeList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	doctype, err := flow.DocTypes.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	}

	arr, err := flow.DocTypes.List(0, 0)
	if err != nil {
		beego.Error(err)
	}
	list := doctypelist{doctype, page1, len(arr)}
	// tx.Commit()
	c.Data["json"] = list
	c.ServeJSON()
}

// @Title post wf doctype...
// @Description post workflowdoctype..
// @Param name query string false "The name of doctype"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtypeupdate [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowTypeUpdate() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	dtid := c.Input().Get("dtid")
	//pid转成64为
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//定义流程类型
	err = flow.DocTypes.Rename(tx, flow.DocTypeID(dtID), name) //"图纸设计流程"
	if err != nil {
		fmt.Println(err)
	}

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf doctype...
// @Description post workflowdoctype..
// @Param name query string  true "The name of doctype"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtypedelete [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowTypeDelete() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义流程类型
	_, err := flow.DocTypes.New(tx, name) //"图纸设计流程"
	if err != nil {
		fmt.Println(err)
	}

	tx.Commit() //这个必须要！！！！！！

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf docstate...
// @Description post workflowdocstate..
// @Param name query string  true "The name of docstate"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowstate [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowState() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义流程状态
	_, err := flow.DocStates.New(tx, name) //"设计中..."
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	// dsID2, err := flow.DocStates.New(tx, "校核中...")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dsID3, err := flow.DocStates.New(tx, "审查中...")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// flow.DocStates.New(tx, "批准中...")
	// flow.DocStates.New(tx, "申报中...")
	// flow.DocStates.New(tx, "评估中...")
	// flow.DocStates.New(tx, "审批中...")
	c.Data["json"] = "ok"
	c.ServeJSON()
}

// @Title post wf docstate...
// @Description post workflowdocstate..
// @Param page query string false "The page of docstate"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowstatelist [get]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowStateList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	docstate, err := flow.DocStates.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	}
	// tx.Commit()
	c.Data["json"] = docstate
	c.ServeJSON()
}

// @Title post wf docaction...
// @Description post workflowdocaction..
// @Param name query string  true "The name of docaction"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowaction [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowAction() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	reconfirm := c.Input().Get("reconfirm")
	beego.Info(reconfirm)
	var reconfirm1 bool
	if reconfirm == "true" {
		reconfirm1 = true
	} else {
		reconfirm1 = false
	}
	//定义流程动作类型
	_, err := flow.DocActions.New(tx, name, reconfirm1) //"设计完成后提交"改变状态设计中...为校核中...
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

// @Title get wf docaction...
// @Description get workflowdocaction..
// @Param page query string true "The page of docaction"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowactionlist [get]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowActionList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	docstate, err := flow.DocActions.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = docstate
	c.ServeJSON()
}

// @Title post wf transition...
// @Description post transition..
// @Param name query string  true "The name of transition"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtransition [post]
// 管理员定义流程类型doctype、流程状态state、流程节点node、
// 流程流向transition，输入doctype、docstate1、docaction、docstate2
func (c *FlowController) FlowTransition() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	// dtid := c.Input().Get("dtid")
	dtid := c.GetString("dtid")
	// beego.Info(dtid)
	//pid转成64为
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsid1 := c.Input().Get("dsid1")
	dsID1, err := strconv.ParseInt(dsid1, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	daid := c.Input().Get("daid")
	daID, err := strconv.ParseInt(daid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsid2 := c.Input().Get("dsid2")
	dsID2, err := strconv.ParseInt(dsid2, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//添加流程规则1:oldstate1 action1 newstate2
	err = flow.DocTypes.AddTransition(tx, flow.DocTypeID(dtID), flow.DocStateID(dsID1), flow.DocActionID(daID), flow.DocStateID(dsID2))
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	//添加流程规则2:oldstate2 action2 newstate3
	// err = flow.DocTypes.AddTransition(tx, flow.DocTypeID(dtID1), flow.DocStateID(dsID2), flow.DocActionID(daID2), flow.DocStateID(dsID3))
	// if err != nil {
	// 	beego.Error(err)
	// }
}

//后端分页的数据结构
type transitionlist struct {
	Transisions []*flow.Transitionstruct `json:"transitions"`
	Page        int64                    `json:"page"`
	Total       int                      `json:"total"` //string或int64都行！
}

// @Title get wf transition...
// @Description post transition..
// @Param page query string  true "The page of transition"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowtransitionlist [get]
// 展示doctype下from docstate可能的transion
func (c *FlowController) FlowTransitionList() {
	// dtid := c.Input().Get("dtid")
	// dtID, err := strconv.ParseInt(dtid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// dsid := c.Input().Get("dsid")
	// dsID, err := strconv.ParseInt(dsid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// transisions, err := flow.DocTypes.Transitions(flow.DocTypeID(dtID), flow.DocStateID(dsID))
	// if err != nil {
	// 	beego.Error(err)
	// }
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	transisions, err := flow.DocTypes.TransitionsList(offset, limit1)
	if err != nil {
		beego.Error(err)
	}
	arr, err := flow.DocTypes.TransitionsList(0, 0)
	if err != nil {
		beego.Error(err)
	}
	list := transitionlist{transisions, page1, len(arr)}
	c.Data["json"] = list
	c.ServeJSON()
}

// @Title post wf Workflow...
// @Description post Workflow..
// @Param name query string  true "The name of Workflow"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowworkflow [post]
// 管理员定义流程Workflow
// 输入doctype和初始action
func (c *FlowController) FlowWorkflow() {
	// func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)//注册驱动
	// orm.RegisterModel(new(Model))//注册 model
	// orm.RegisterDataBase("default", "mysql", "test:123456@/test?charset=utf8",30,30)//注册默认数据库
	//orm.RegisterDataBase("default", "mysql", "test:@/test?charset=utf8")//密码为空格式
	// }
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name") //图纸设计-三级校审流程
	dtid := c.Input().Get("dtid")
	//pid转成64为
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsid := c.Input().Get("dsid")
	dsID, err := strconv.ParseInt(dsid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//定义流程类型doctype下的唯一流程workflow
	workflowID, err := flow.Workflows.New(tx, name, flow.DocTypeID(dtID), flow.DocStateID(dsID)) //初始状态是“设计中...”——校核——审查——完成
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		beego.Info(workflowID)
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	// workflowID2, err := flow.Workflows.New(tx, "图纸设计-二级校审流程", dtID1, dsID1) //初始状态是“设计中...”-“校核”——完成
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// beego.Info(workflowID2)
	//定义合同评审下的流程类型：部门合同流程，总院合同流程
	//略
}

// @Title post wf Workflow...
// @Description post Workflow..
// @Param page query string  true "The page of Workflow"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowworkflowlist [get]
// 管理员定义流程Workflow
// 输入doctype和初始action
func (c *FlowController) FlowWorkflowList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	workflows, err := flow.Workflows.List(offset, limit1)
	c.Data["json"] = workflows
	c.ServeJSON()
}

// @Title post wf AccessContext...
// @Description post AccessContext..
// @Param name query string  true "The name of AccessContext"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowaccesscontext [post]
// 管理员定义流程AccessContext
// 流程命名空间
func (c *FlowController) FlowAccessContext() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	//定义用户、组、角色、权限集合
	_, err := flow.AccessContexts.New(tx, name) //"Context"
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

// @Title get wf AccessContext...
// @Description post AccessContext..
// @Param page query string  true "The page of AccessContext"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowaccesscontextlist [get]
// 管理员定义流程AccessContext
// 流程命名空间
func (c *FlowController) FlowAccessContextList() {
	prefix := c.Input().Get("prefix")
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	accesscontexts, err := flow.AccessContexts.List(prefix, offset, limit1)
	c.Data["json"] = accesscontexts
	c.ServeJSON()
}

// @Title post wf Node...
// @Description post Node..
// @Param name query string  true "The name of Node"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flownode [post]
// 管理员定义流程Node
// 流程node，输入doctype，docstate1，access，workflow，name和nodetype
// A `Node` each has to be defined for each document state of the workflow,
// except the final state. Please look at `_Workflows.AddNode`.
func (c *FlowController) FlowNode() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	dtid := c.Input().Get("dtid")
	//pid转成64为
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsid := c.Input().Get("dsid")
	dsID, err := strconv.ParseInt(dsid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	acid := c.Input().Get("acid")
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	var flownodetype flow.NodeType
	nodetype := c.Input().Get("nodetype") // flow.NodeTypeBegin
	switch nodetype {
	case "begin":
		flownodetype = flow.NodeTypeBegin
	case "end":
		flownodetype = flow.NodeTypeEnd
	case "linear":
		flownodetype = flow.NodeTypeLinear
	case "branch":
		flownodetype = flow.NodeTypeBranch
	case "joinany":
		flownodetype = flow.NodeTypeJoinAny
	case "joinall":
		flownodetype = flow.NodeTypeJoinAll
		// NodeTypeEnd = "end"
		// // NodeTypeLinear : one incoming, one outgoing
		// NodeTypeLinear = "linear"
		// // NodeTypeBranch : one incoming, two or more outgoing
		// NodeTypeBranch = "branch"
		// // NodeTypeJoinAny : two or more incoming, one outgoing
		// NodeTypeJoinAny = "joinany"
		// // NodeTypeJoinAll : two or more incoming, one outgoing
		// NodeTypeJoinAll = "joinall"
	}
	//根据doctypeid获得workflow
	workflow, err := flow.Workflows.GetByDocType(flow.DocTypeID(dtID))
	//定义流程类型workflow下的具体每个节点node，用户对文件执行某个动作（event里的action）后，会沿着这些节点走
	// AddNode maps the given document state to the specified node.  This
	// map is consulted by the workflow when performing a state transition
	// of the system.nodeID1
	_, err = flow.Workflows.AddNode(tx, flow.DocTypeID(dtID), flow.DocStateID(dsID), flow.AccessContextID(acID), workflow.ID, name, flownodetype)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID2, accessContextID1, workflowID1, "图纸设计-三级校审流程-校核", flow.NodeTypeLinear)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// _, err = flow.Workflows.AddNode(tx, dtID1, dsID3, accessContextID1, workflowID1, "图纸设计-三级校审流程-审查", flow.NodeTypeEnd)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

// @Title get wf Node...
// @Description post Node..
// @Param page query string  true "The page of Node"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flownodelist [get]
// 管理员定义流程Node
func (c *FlowController) FlowNodeList() {
	// var offset, limit int64
	// limit = 5
	// page := c.Input().Get("page")
	// page1, err := strconv.ParseInt(page, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	workflowid := c.Input().Get("workflowid")
	workflowid1, err := strconv.ParseInt(workflowid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// if page1 <= 1 {
	// 	offset = 0
	// } else {
	// 	offset = (page1 - 1) * limit
	// }
	// nodes, err := flow.Nodes.NodeList(offset, limit)
	nodes, err := flow.Nodes.List(flow.WorkflowID(workflowid1))
	c.Data["json"] = nodes
	c.ServeJSON()
}

// @Title post wf user...
// @Description post user..
// @Param firstname query string  true "The firstname of user"
// @Param lastname query string  true "The lastname of user"
// @Param email query string  true "The email of user"
// @Param active query string  true "The active of user"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowuser [post]
// 管理员定义流程user
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowUser() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	first_name := c.Input().Get("firstname")
	last_name := c.Input().Get("lastname")
	email := c.Input().Get("email")
	active := c.Input().Get("active")
	var err error
	var res sql.Result
	if active == "true" {
		res, err = tx.Exec("INSERT INTO users_master(first_name, last_name, email, active) VALUES(?, ?, ?, ?)", first_name, last_name, email, 1)
	} else {
		res, err = tx.Exec("INSERT INTO users_master(first_name, last_name, email, active) VALUES(?, ?, ?, ?)", first_name, last_name, email, 0)
	}
	//定义用户-组-角色-权限关系
	// res, err := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES(` + first_name + `, ` + last_name + `, ` + email + `, ` + active1 + `)`)
	if err != nil {
		log.Fatalf("%v\n", err)
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		uid, _ := res.LastInsertId()
		uID1 := flow.UserID(uid)
		_, err = flow.Groups.NewSingleton(tx, uID1)
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}

	// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-2', 'email2@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ = res.LastInsertId()
	// uID2 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID2)

	// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-3', 'email3@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ = res.LastInsertId()
	// uID3 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID3)

	// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
	// 		VALUES('秦', '晓川-4', 'email4@example.com', 1)`)
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// uid, _ = res.LastInsertId()
	// uID4 := flow.UserID(uid)
	// _, err = flow.Groups.NewSingleton(tx, uID4)
}

// @Title get wf user...
// @Description post user..
// @Param page query string  true "The page of user"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowuserlist [get]
// 管理员定义流程user
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowUserList() {
	prefix := c.Input().Get("prefix")
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	nodes, err := flow.Users.List(prefix, offset, limit1)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = nodes
		c.ServeJSON()
	}
}

// @Title post wf Group...
// @Description post Group..
// @Param name query string  true "The name of Group"
// @Param grouptype query string  true "The type of Group"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgroup [post]
// 管理员定义流程Group
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowGroup() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	grouptype := c.Input().Get("grouptype")
	//注意：单人组自动建立！！！
	_, err := flow.Groups.New(tx, name, grouptype) //).(flow.GroupID)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	// gID2 := fatal1(flow.Groups.New(tx, "校核人员组", "G")).(flow.GroupID)
}

// @Title get wf Group...
// @Description post Group..
// @Param page query string  true "The page of Group"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgrouplist [get]
// 管理员定义流程Group
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowGroupList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	nodes, err := flow.Groups.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = nodes
		c.ServeJSON()
	}
}

// @Title post wf GroupUser...
// @Description post Group..
// @Param name query string  true "The name of GroupUser"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowusergroup [post]
// 管理员定义流程GroupUser
// 将users加入group
func (c *FlowController) FlowUserGroup() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	// uid := c.Input().Get("uid")
	// beego.Info(uid)
	uid := make([]string, 0, 2)
	c.Ctx.Input.Bind(&uid, "uid") //ul ==[str array]
	// uid := c.GetStrings("uid")
	beego.Info(uid)
	uID, err := strconv.ParseInt(uid[0], 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gid := c.Input().Get("gid")
	//pid转成64为
	gID, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = flow.Groups.AddUser(tx, flow.GroupID(gID), flow.UserID(uID))
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

type GroupUsers struct {
	Id    flow.GroupID
	Group *flow.Group
	Users []*flow.User
}

// @Title get wf GroupUsers...
// @Description post Group..
// @Param name query string  true "The name of GroupUser"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgroupuserslist [get]
// 查询Group下的所有Users
func (c *FlowController) FlowGroupUsersList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	groupusers := make([]GroupUsers, 0)
	groups, err := flow.Groups.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	} else {
		for _, v := range groups {
			groupusersarr := make([]GroupUsers, 1)
			users, err := flow.Groups.Users(flow.GroupID(v.ID))
			if err != nil {
				beego.Error(err)
			}
			groupusersarr[0].Id = v.ID
			groupusersarr[0].Group = v
			groupusersarr[0].Users = users
			groupusers = append(groupusers, groupusersarr...)
		}
	}
	c.Data["json"] = groupusers
	c.ServeJSON()
}

// @Title post wf Role...
// @Description post Role..
// @Param name query string  true "The name of Role"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowrole [post]
// 管理员定义流程Role
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowRole() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	name := c.Input().Get("name")
	_, err := flow.Roles.New(tx, name)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	// roleID2 := fatal1(flow.Roles.New(tx, "校核人员角色")).(flow.RoleID)
}

// @Title get wf Role...
// @Description post Role..
// @Param page query string  true "The page of Role"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowrolelist [get]
// 查询所有role
func (c *FlowController) FlowRoleList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	roles, err := flow.Roles.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = roles
		c.ServeJSON()
	}
}

// @Title post wf Permission...
// @Description post Permission..
// @Param name query string  true "The name of Permission"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowpermission [post]
// 管理员定义流程Permission
// 流程动作action、流程流向transition、流程事件event
func (c *FlowController) FlowPermission() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	roleid := c.Input().Get("roleid")
	//pid转成64为
	roleID, err := strconv.ParseInt(roleid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dtid := c.Input().Get("dtid")
	//pid转成64为
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据用户选择的动作
	daid := c.Input().Get("daid")
	daID, err := strconv.ParseInt(daid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	actions := []flow.DocActionID{flow.DocActionID(daID)} //[]flow.DocActionID{daID1, daID2, daID3, daID4}

	//给角色role赋予action权限
	err = flow.Roles.AddPermissions(tx, flow.RoleID(roleID), flow.DocTypeID(dtID), actions)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	// fatal0(flow.Roles.AddPermissions(tx, roleID2, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4, daID5, daID6, daID7}))
}

type RolePermission struct {
	RoleID flow.RoleID
	Action map[string]struct {
		DocTypeID flow.DocTypeID
		Actions   []*flow.DocAction
	}
}

type Actions struct {
	DocTypeID flow.DocTypeID
	Actions   []*flow.DocAction
}

// @Title get wf Permission...
// @Description post Permission..
// @Param page query string  true "The page of Permission"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowrolepermissionlist [get]
// 查询role和对应对应的permission
func (c *FlowController) FlowRolePermissionList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	// rolepermission := make([]RolePermission, 0)
	rolepermission := make([]flow.RolePermission, 0)
	roles, err := flow.Roles.List(offset, limit1)
	if err != nil {
		beego.Error(err)
	} else {
		for _, v := range roles {
			rolepermissionarr := make([]flow.RolePermission, 1)
			// rolepermissions, err := flow.Roles.Permissions(flow.RoleID(v.ID))
			rolepermissions, err := flow.Roles.PermissionsList(flow.RoleID(v.ID))
			if err != nil {
				beego.Error(err)
			}
			rolepermissionarr[0] = rolepermissions
			if rolepermissions.TypeAction.DocTypeID == 0 {
				flowaction := make([]*flow.DocAction, 0, 1)
				var da flow.DocAction
				da.ID = 0
				flowaction = append(flowaction, &da)
				rolepermissionarr[0].TypeAction.Actions = flowaction
			}
			// rolepermissionarr[0].Action = rolepermissions
			rolepermission = append(rolepermission, rolepermissionarr...)
		}
	}
	c.Data["json"] = rolepermission
	c.ServeJSON()
}

// @Title post wf GroupRole...
// @Description post GroupRole..
// @Param name query string  true "The name of GroupRole"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgrouprole [post]
// 管理员定义流程GroupRole
// 来自accesscontext
func (c *FlowController) FlowGroupRole() {
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, _ := db.Begin()
	db.Close()

	acid := c.Input().Get("acid")
	//pid转成64为
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gid := c.Input().Get("gid")
	//pid转成64为
	gID, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	roleid := c.Input().Get("roleid")
	//pid转成64为
	roleID, err := strconv.ParseInt(roleid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//给用户组group赋予角色role
	err = flow.AccessContexts.AddGroupRole(tx, flow.AccessContextID(acID), flow.GroupID(gID), flow.RoleID(roleID))
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
	//将group和role加到accesscontext里——暂时不理解
	// err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID2, roleID2)
	// if err != nil {
	// 	beego.Error(err) //UNIQUE constraint failed: wf_ac_group_roles.ac_id已修补
	// }
}

// @Title get wf GroupRole...
// @Description get GroupRole..
// @Param page query string  true "The page of GroupRole"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgrouprolelist [get]
// 查询group的角色role-来自accesscontext
func (c *FlowController) FlowGroupRoleList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	// accesscontexts, err := flow.AccessContexts.List(prefix, offset, limit)
	//groups,err:=
	// var gids []flow.GroupID
	// groupid := flow.GroupID(11)
	// gids = append(gids, groupid)
	// groupid = flow.GroupID(12)
	// gids = append(gids, groupid)
	// groupid = flow.GroupID(13)
	// gids = append(gids, groupid)
	// grouproles, err := flow.AccessContexts.GroupRoles(flow.AccessContextID(1), gids, offset, limit)
	grouproles, err := flow.AccessContexts.GroupRolesList(offset, limit1)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = grouproles
	c.ServeJSON()
}

// @Title post wf document
// @Description post document
// @Param dtid query string  true "The doctypeid of document"
// @Param acid query string  true "The accesscontext of document"
// @Param gid query string  true "The groupid of Group"
// @Param name query string  true "The name of document"
// @Param data query string  true "The data of document"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdoc [post]
func (c *FlowController) FlowDoc() {
	//连接数据库
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}

	//查询预先定义的doctype流程类型
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	acid := c.Input().Get("acid")
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gid := c.Input().Get("gid")
	gID, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	name := c.Input().Get("docname")
	data := c.Input().Get("docdata")

	//查询预先定义的流程类型workflow，这个相当于doctype下面再分很多种流程
	//比如doctype为图纸设计流程，下面可以分为二级校审流程，三级校审流程，四级校审流程
	//查询预先定义的doctype流程类型

	//查询context——这个应该是管理用户-组-权限的
	//查询预先定义的doctype流程类型

	//开始为具体一个文件设立流程-此处是新建一个文件。对于旧文件应该怎么操作来着？
	//document根据doctype取得唯一workflow的state作为document的state
	docNewInput := flow.DocumentsNewInput{
		DocTypeID:       flow.DocTypeID(dtID),       //属于图纸设计类型的流程
		AccessContextID: flow.AccessContextID(acID), //所有用户权限符合这个contex的要求
		GroupID:         flow.GroupID(gID),          //groupId,初始状态下的用户组，必须是个人用户组（一个用户也可以成为一个独特的组，因为用户无法赋予角色，所以必须将用户放到组里）
		Title:           name,                       //这个文件的名称
		Data:            data,                       //文件的描述
	}
	// flow.Documents.New(tx, &docNewInput)
	_, err = flow.Documents.New(tx, &docNewInput)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

// @Title get wf doclist
// @Description get workflow doclist
// @Param dtid query string  true "The id of doctype"
// @Param page query string  true "The page of doc"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdoclist [get]
// 文件列表页，水平显示每个文件的状态
func (c *FlowController) FlowDocList() {
	//查询预先定义的doctype流程类型
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// acid := c.Input().Get("acid")
	// acID, err := strconv.ParseInt(acid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// gid := c.Input().Get("gid")
	// gID, err := strconv.ParseInt(gid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// dsid := c.Input().Get("dsid")
	// dsID, err := strconv.ParseInt(dsid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// documentslistinput := flow.DocumentsListInput{
	// 	DocTypeID:       flow.DocTypeID(dtID),       // Documents of this type are listed; required
	// 	AccessContextID: flow.AccessContextID(acID), // Access context from within which to list; required
	// 	GroupID:         flow.GroupID(gID),          // 本人所在的组List documents created by this (singleton) group
	// 	DocStateID:      flow.DocStateID(dsID),      // 忽略List documents currently in this state
	// 	//CtimeStarting:   time.Now(), // List documents created after this time
	// 	//CtimeBefore:     time.Now(), // List documents created before this time
	// 	//TitleContains:   string,     // List documents whose title contains the given text; expensive operation
	// 	//RootOnly:        bool,       // List only root (top-level) documents
	// }
	// documents, err := flow.Documents.List(&documentslistinput, offset, limit)
	//列出符合条件的events
	// DocEventsListInput指定一组筛选条件来缩小文档清单。
	// type DocEventsListInput struct {
	// 	DocTypeID                   // Events on documents of this type are listed
	// 	AccessContextID             // Access context from within which to list
	// 	GroupID                     // List events created by this (singleton) group
	// 	DocStateID                  // List events acting on this state
	// 	CtimeStarting   time.Time   // List events created after this time
	// 	CtimeBefore     time.Time   // List events created before this time
	// 	Status          EventStatus // List events that are in this state of application
	// }
	// DocEvents, err := flow.DocEvents.List(DocEventsListInput)
	// if err != nil {
	// 	beego.Error(err)
	// }

	var offset, limit1, page1 int64
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	documents, err := flow.Documents.DocumentList(flow.DocTypeID(dtID), offset, limit1)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = documents
		c.ServeJSON()
	}
}

// @Title get wf doclist
// @Description get workflow doclist
// @Param dtid query string  true "The id of doctype"
// @Param page query string  true "The page of doc"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdoctransitionlist [get]
// 文件列表页，显示doc和文档可能的走向——不要正确，暂停
// func (c *FlowController) FlowDocTransitionList() {
// 	// var tx *sql.Tx
// 	var err error
// 	var offset, limit1, page1 int64
// 	limit := c.Input().Get("limit")
// 	if limit == "" {
// 		limit1 = 0
// 	} else {
// 		limit1, err = strconv.ParseInt(limit, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}
// 	page := c.Input().Get("page")
// 	if page == "" {
// 		limit1 = 0
// 		page1 = 1
// 	} else {
// 		page1, err = strconv.ParseInt(page, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}
// 	if page1 <= 1 {
// 		offset = 0
// 	} else {
// 		offset = (page1 - 1) * limit1
// 	}
// 	//查询预先定义的doctype流程类型
// 	//文件所在目录的doctype
// 	dtid := c.Input().Get("dtid")
// 	dtID, err := strconv.ParseInt(dtid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//文件所在的目录的accesscontext
// 	acid := c.Input().Get("acid")
// 	acID, err := strconv.ParseInt(acid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//登录用户所处的用户组
// 	gid := c.Input().Get("gid")
// 	gID, err := strconv.ParseInt(gid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//文件所处某个状态的目录，如果不填，列出所有状态的？
// 	dsid := c.Input().Get("dsid")
// 	dsID, err := strconv.ParseInt(dsid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	documentslistinput := flow.DocumentsListInput{
// 		DocTypeID:       flow.DocTypeID(dtID),       // Documents of this type are listed; required
// 		AccessContextID: flow.AccessContextID(acID), // Access context from within which to list; required
// 		GroupID:         flow.GroupID(gID),          // 本人所在的组List documents created by this (singleton) group
// 		DocStateID:      flow.DocStateID(dsID),      // 可以忽略List documents currently in this state
// 		//CtimeStarting:   time.Now(), // List documents created after this time
// 		//CtimeBefore:     time.Now(), // List documents created before this time
// 		//TitleContains:   string,     // List documents whose title contains the given text; expensive operation
// 		//RootOnly:        bool,       // List only root (top-level) documents
// 	}
// 	documents, err := flow.Documents.List(&documentslistinput, offset, limit1)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//列出每个文档下一个transition——没这个针对具体一个文档的方式
// 	for _, v := range documents {
// 		TransitionMap, err := flow.DocTypes.Transitions(flow.DocTypeID(dtID), v.State.ID)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}

// 	if err != nil {
// 		beego.Error(err)
// 		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
// 		c.ServeJSON()
// 	} else {
// 		c.Data["json"] = docevents
// 		c.ServeJSON()
// 	}
// }

// A `DocEvent` has to be constructed构造 to signal信号 to the workflow
// that a `DocAction` has been performed执行 by either the user or a system event.
// Document states are changed by `Workflow.ApplyEvent`, to which this event has to be fed.
// In turn, the workflow loads the node corresponding to the document state specified in the event,
// and applies the event to the node with the specified document.
// The node applies the given document action to transition the document into
// the target state of the action. These transitions are defined for document types.
// Please see `Transition` and `TransitionMap` in the file `doctype.go`.

// myDocument, err := Documents.Get(tx, DocTypeID(1), DocumentID(1))
// retMap, err := DocTypes.Transitions(DocTypeID(1), myDocument.State.ID)
// Yes, once you have the transitions defined, this query gives you the set of possible next states,
// together with which actions would lead to such transitions.

// @Title post wf event
// @Description get docevent
// @Param name query string  true "The name of event"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdocevent [post]
// 添加events：
func (c *FlowController) FlowEvent() {
	//连接数据库
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}

	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	docid := c.Input().Get("docid")
	docID, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	dsid := c.Input().Get("dsid")
	dsID, err := strconv.ParseInt(dsid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	daid := c.Input().Get("daid")
	daID, err := strconv.ParseInt(daid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gid := c.Input().Get("gid")
	gID, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	text := c.Input().Get("text")
	beego.Info(text)
	//建立好document，循环建立events，根据哪个来建立？
	//根据document的Doctypes.Transitions获取state和action
	//循环建立events，然后展示给客户端
	//用户点开这个文件，根据文件的状态，list出所有这个状态的events，比如前进，后退等
	docEventInput := flow.DocEventsNewInput{
		DocTypeID:   flow.DocTypeID(dtID), //flow.DocTypeID(1),
		DocumentID:  flow.DocumentID(docID),
		DocStateID:  flow.DocStateID(dsID),  //document state must be this state，文档的现状状态
		DocActionID: flow.DocActionID(daID), //Action performed by `Group`; required,由用户组执行的操作
		GroupID:     flow.GroupID(gID),      //Group (user) who performed the action that raised this event; required，执行引发此事件的操作的组(用户)
		Text:        text,                   //Any comments or notes; required，
	}
	beego.Info(docEventInput)
	_, err = flow.DocEvents.New(tx, &docEventInput)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx.Commit() //这个必须要！！！！！！
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

// @Title get wf eventlist
// @Description get doceventlist
// @Param page query string  true "The page of event"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdoceventlist [get]
// 查询events：
func (c *FlowController) FlowEventList() {
	// var tx *sql.Tx
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	acid := c.Input().Get("acid")
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gid := c.Input().Get("gid")
	gID, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsid := c.Input().Get("dsid")
	dsID, err := strconv.ParseInt(dsid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//用户点开这个文件，根据文件的状态，list出所有这个状态的events，比如前进，后退等
	//doctypeid从哪来？所有操作都带doctype吧，因为当前走的流程都属于这个doctype下的
	docEventListInput := flow.DocEventsListInput{
		DocTypeID:       flow.DocTypeID(dtID),       // Events on documents of this type are listed
		AccessContextID: flow.AccessContextID(acID), // Access context from within which to list
		GroupID:         flow.GroupID(gID),          // List events created by this (singleton) group
		DocStateID:      flow.DocStateID(dsID),      // List events acting on this state
		// CtimeStarting:   time.Time,             // List events created after this time
		// CtimeBefore:     time.Time,             // List events created before this time
		Status: flow.EventStatusAll, // EventStatusAll,List events that are in this state of application
	}

	var offset, limit1, page1 int64
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	myDocEvent, err := flow.DocEvents.List(&docEventListInput, offset, limit1)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = myDocEvent
		c.ServeJSON()
	}
}

// @Title get wf document
// @Description get document
// @Param dtid query string  true "The id of doctype"
// @Param acid query string  true "The id of accesscontext"
// @Param gid query string  false "The id of group"
// @Param dsid query string  false "The id of docstate"
// @Param page query string  true "The page of doc"
// @Param limit query string  false "The limit page of doc"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdocumentlist [get]
// 1.列表显示文档
// 2.点击一个具体文档——显示详情——显示actions
func (c *FlowController) FlowDocumentList() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//查询预先定义的doctype流程类型
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	acid := c.Input().Get("acid")
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//查出登录用户所在group
	// uses.GroupsOf(uid)
	// SingletonGroupOf(uid)
	// gid := c.Input().Get("gid")
	// if gid != "" {
	// 	gID, err := strconv.ParseInt(gid, 10, 64)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	// dsid := c.Input().Get("dsid")
	// if gid != "" {
	// 	dsID, err := strconv.ParseInt(dsid, 10, 64)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	documentslistinput := flow.DocumentsListInput{
		DocTypeID:       flow.DocTypeID(dtID),       // Documents of this type are listed; required
		AccessContextID: flow.AccessContextID(acID), // Access context from within which to list; required
		// GroupID:         flow.GroupID(gID),          // List documents created by this (singleton) group
		// DocStateID:      flow.DocStateID(dsID),      // 忽略List documents currently in this state
		//CtimeStarting:   time.Now(), // List documents created after this time
		//CtimeBefore:     time.Now(), // List documents created before this time
		//TitleContains:   string,     // List documents whose title contains the given text; expensive operation
		//RootOnly:        bool,       // List only root (top-level) documents
	}
	documents, err := flow.Documents.List(&documentslistinput, offset, limit1)

	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = documents
		c.ServeJSON()
	}
}

type DocumentDetail struct {
	// DocTypeId    flow.DocTypeID
	Document *flow.Document
	Action   []flow.DocAction
	History  []*flow.DocEventsHistory
	Text     string
}

// @Title get wf document details
// @Description get documentdetail
// @Param dtid query string  true "The id of doctype"
// @Param docid query string  true "The id of doc"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowdocumentdetail [get]
// 2.点击一个具体文档——显示详情——显示actions
func (c *FlowController) FlowDocumentDetail() {
	var tx *sql.Tx
	//查询预先定义的doctype流程类型
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	docid := c.Input().Get("docid")
	docID, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	document, err := flow.Documents.Get(tx, flow.DocTypeID(dtID), flow.DocumentID(docID))
	if err != nil {
		beego.Error(err)
	}
	//列出actions
	//列出符合条件的actions
	TransitionMap, err := flow.DocTypes.Transitions(flow.DocTypeID(dtID), document.State.ID)
	if err != nil {
		beego.Error(err)
	}
	//列出符合要求（有权限）的接受groups
	// beegon.Info(TransitionMap[document.State.ID].Transitions[].Upon.ID)
	// {
	//   "8": {
	//     "From": {
	//       "ID": 8,
	//       "Name": "校核中..."
	//     },
	//     "Transitions": {
	//       "12": {
	//         "Upon": {
	//           "ID": 12,
	//           "Name": "审查",
	//           "Reconfirm": false
	//         },
	//         "To": {
	//           "ID": 9,
	//           "Name": "审查中..."
	//         }
	//       },
	//       "14": {
	//         "Upon": {
	//           "ID": 14,
	//           "Name": "评估",
	//           "Reconfirm": true
	//         },
	//         "To": {
	//           "ID": 7,
	//           "Name": "设计中..."
	//         }
	//       }
	//     }
	//   }
	// }
	//非数组模式
	// var documentdetail DocumentDetail
	// for key, value := range TransitionMap[document.State.ID].Transitions {
	// 	beego.Info(key)
	// 	beego.Info(value.Upon.ID)
	// 	documentdetailarr := make([]flow.DocAction, 1)
	// 	documentdetailarr[0].ID = value.Upon.ID
	// 	documentdetailarr[0].Name = value.Upon.Name
	// 	documentdetail.Action = append(documentdetail.Action, documentdetailarr...)
	// }
	// documentdetail.Document = document
	//数组模式
	documentdetail := make([]DocumentDetail, 1)
	if _, ok := TransitionMap[document.State.ID]; ok {
		//存在
		for _, value := range TransitionMap[document.State.ID].Transitions {
			// beego.Info(key)
			// beego.Info(value.Upon.ID)
			documentdetailarr := make([]flow.DocAction, 1)
			documentdetailarr[0].ID = value.Upon.ID
			documentdetailarr[0].Name = value.Upon.Name
			documentdetail[0].Action = append(documentdetail[0].Action, documentdetailarr...)
		}
	}
	documentdetail[0].Document = document
	//查出历史记录
	// docEventListInput := flow.DocEventsListInput{
	// 		DocTypeID:       flow.DocTypeID(dtID),       // Events on documents of this type are listed
	// 		AccessContextID: flow.AccessContextID(acID), // Access context from within which to list
	// 		GroupID:         flow.GroupID(gID),          // List events created by this (singleton) group
	// 		DocStateID:      flow.DocStateID(dsID),      // List events acting on this state
	// 		// CtimeStarting:   time.Time,             // List events created after this time
	// 		// CtimeBefore:     time.Time,             // List events created before this time
	// 		Status: flow.EventStatusAll, // EventStatusAll,List events that are in this state of application
	// 	}
	// myDocEvent, err := flow.DocEvents.List(&docEventListInput, offset, limit1)
	doceventshistory, err := flow.DocEvents.DocEventsHistory(flow.DocTypeID(dtID), flow.DocumentID(docID))
	if err != nil {
		beego.Error(err)
	}
	documentdetail[0].History = doceventshistory
	//action按钮
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = documentdetail
		c.ServeJSON()
	}
}

// @Title post wf next
// @Description post workflow next
// @Param dtid query string  true "The id of doctype"
// @Param daid query string  true "The id of action"
// @Param docid query string  true "The id of document"
// @Param gid query string  true "The id of group"
// @Param text query string  false "The text of apply"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flownext [post]
// FlowDocAction列出了文档和动作，用户点击action，则这里进行修改docstate
func (c *FlowController) FlowNext() {
	// var tx *sql.Tx //用这个nil，后面就不用commit了吧，都在flow里commit了。
	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
	if tdb == nil {
		log.Fatal("given database handle is `nil`")
	}
	db := tdb
	tx, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}
	tx1, err := db.Begin()
	if err != nil {
		beego.Error(err)
	}
	// GetByDocType从数据库检索请求的工作流的详细信息。
	// 注意:此方法检索工作流的主要信息。组成此工作流的节点的信息必须单独获取。
	// jsoninfo := c.GetString("docaction") //获取formdata
	// beego.Info(jsoninfo)
	// {"Document":
	// 	{
	// 		"ID":19,
	// 		"DocType":{"ID":3,"Name":"图纸设计"},
	// 		"Path":"",
	// 		"AccessContext":{"ID":1},
	// 		"DocState":{"ID":8,"Name":"校核中..."},
	// 		"Group":{"ID":11,"Name":"设计人员组","GroupType":""},
	// 		"Ctime":"2019-01-11T21:42:50Z",
	// 		"Title":"厂房布置图",
	// 		"Data":"设计、制图: 秦晓川1, 校核:秦晓川2"},
	// 		"Action":[
	// 			{"ID":14,"Name":"评估","Reconfirm":false},
	// 			{"ID":12,"Name":"审查","Reconfirm":false}
	// 		]
	// 	}
	// }
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	daid := c.Input().Get("daid")
	daID, err := strconv.ParseInt(daid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	docid := c.Input().Get("docid")
	docID, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据docid取得document
	document, err := flow.Documents.Get(tx, flow.DocTypeID(dtID), flow.DocumentID(docID))
	if err != nil {
		beego.Error(err)
	}
	//根据document取得workflow
	myWorkflow, err := flow.Workflows.GetByDocType(document.DocType.ID)
	if err != nil {
		beego.Error(err)
	}

	//当前用户所在的用户组
	singletongroup, err := flow.Users.SingletonGroupOf(flow.UserID(8))
	if err != nil {
		beego.Error(err)
	}
	//接受用户组
	gid := make([]string, 0, 2)
	c.Ctx.Input.Bind(&gid, "gid")
	// beego.Info(gid)
	var groupIds []flow.GroupID
	for _, v := range gid {
		gID, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		groupIds = append(groupIds, flow.GroupID(gID))
	}

	text := c.Input().Get("text")
	if text == "" {
		text = "no comments"
	}
	//建立event
	docEventInput := flow.DocEventsNewInput{
		DocTypeID:   flow.DocTypeID(dtID),
		DocumentID:  flow.DocumentID(docID),
		DocStateID:  document.State.ID,      //document state must be this state，文档的现状状态
		DocActionID: flow.DocActionID(daID), //Action performed by `Group`; required,由用户组执行的操作
		GroupID:     singletongroup.ID,      //Group (user) who performed the action that raised this event; required，执行引发此事件的操作的组(用户)
		Text:        text,                   //Any comments or notes; required，
	}
	// beego.Info(docEventInput)
	deID, err := flow.DocEvents.New(tx, &docEventInput)
	if err != nil {
		beego.Error(err)
	} else {
		tx.Commit()
	}
	myDocEvent, err := flow.DocEvents.Get(flow.DocEventID(deID))
	if err != nil {
		beego.Error(err)
	} else {
		beego.Info(myDocEvent)
	}

	newDocStateId, err := myWorkflow.ApplyEvent(tx1, myDocEvent, groupIds)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "写入失败!"}
		c.ServeJSON()
	} else {
		tx1.Commit() //用这个nil，后面就不用commit了吧，都在flow里commit了。
		fmt.Println("newDocStateId=", newDocStateId, err)
		c.Data["json"] = map[string]interface{}{"err": nil, "data": "写入成功!"}
		c.ServeJSON()
	}
}

// @Title get user mailbox
// @Description get usermailbox
// @Param uid query string true "The id of user"
// @Param page query string true "The page of mailbox"
// @Param limit query string false "The limit page of mailbox"
// @Param unread query string false "The unread of mailbox"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowusermailbox [get]
// 1.列表显示用户邮件
func (c *FlowController) FlowUserMailbox() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//查询预先定义的doctype流程类型
	uid := c.Input().Get("uid")
	uID, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	unread := c.Input().Get("unread")
	var unreadbool bool
	if unread != "" {
		if unread == "true" {
			unreadbool = true
		} else {
			unreadbool = false
		}
	} else {
		unreadbool = false
	}

	notification, err := flow.Mailboxes.ListByUser(flow.UserID(uID), offset, limit1, unreadbool)

	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = notification
		c.ServeJSON()
	}
}

// @Title get group mailbox
// @Description get groupmailbox
// @Param gid query string true "The id of group"
// @Param page query string true "The page of mailbox"
// @Param limit query string false "The limit page of mailbox"
// @Param unread query string false "The unread of mailbox"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /flowgroupmailbox [get]
// 1.列表显示用户邮件
func (c *FlowController) FlowGroupMailbox() {
	var offset, limit1, page1 int64
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 0
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		limit1 = 0
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//查询预先定义的doctype流程类型
	gid := c.Input().Get("gid")
	gID, err := strconv.ParseInt(gid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	unread := c.Input().Get("unread")
	var unreadbool bool
	if unread != "" {
		if unread == "true" {
			unreadbool = true
		} else {
			unreadbool = false
		}
	} else {
		unreadbool = false
	}

	notification, err := flow.Mailboxes.ListByGroup(flow.GroupID(gID), offset, limit1, unreadbool)

	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	} else {
		c.Data["json"] = notification
		c.ServeJSON()
	}
}

// fatal1 expects a value and an error value as its arguments.
func fatal1(val1 interface{}, err error) interface{} {
	if err != nil {
		fmt.Println("%v", err)
	}
	return val1
}

// error0 expects only an error value as its argument.
func error0(err error) error {
	if err != nil {
		fmt.Println("%v", err)
	}
	return err
}

// error1 expects a value and an error value as its arguments.
func error1(val1 interface{}, err error) interface{} {
	if err != nil {
		fmt.Println("%v", err)
		return nil
	}
	return val1
}

// fatal0 expects only an error value as its argument.
func fatal0(err error) {
	if err != nil {
		fmt.Println("%v", err)
	}
}

// Document Type : docType1
// Document States : [
// 	docState1,
// 	docState2,
// 	docState3,
// 	docState4 // for example
// ]
// Document Actions : [
// 	docAction12,
// 	docAction23,
// 	docAction34 // for the above document states
// ]
// Document Type State Transitions : [
// 	docState1 --docAction12--> docState2,
// 	docState2 --docAction23--> docState3,
// 	docState3 --docAction34--> docState4,
// ]
// Access Contexts : [
// 	accCtx1, accCtx2 // for example
// ]
// Workflow : {
// 	Name : wFlow1,
// 	Initial State : docState1
// }
// Nodes : [
// node1: {
// 	Document Type : docType1,
// 	Workflow : wFlow1,
// 	Node Type : NodeTypeBegin, // note this
// 	From State : docState1,
// 	Access Context : accCtx1,
// },
// node2: {
// 	Document Type : docType1,
// 	Workflow : wFlow1,
// 	Node Type : NodeTypeLinear, // note this
// 	From State : docState2,
// 	Access Context : accCtx2, // a different context
// },
// node3: {
// 	Document Type : docType1,
// 	Workflow : wFlow1,
// 	Node Type : NodeTypeEnd, // note this
// 	From State : docState3,
// 	Access Context : accCtx1,
// },
// ]

// With the above setup, you can dispatch document events to the workflow appropriately.
// With each event, the workflow moves along, as defined.

// * When you create new documents, their events need **_not_** be created at the same time.
// Events should be created and applied in response to user actions (or system events).
// * Whether you list all possible actions in a list of documents is up to you.
// In my opinion, the possible actions - for a given user - on a given document,
// are best made available in the detailed view of that document (not in a document list view).
// * Yes, when the user clicks on an action button, we can create the corresponding event instance,
// and apply it to the document's workflow.
// * With a `DocEventID`, you can fetch the corresponding event instance using `DocEvents.Get`

// 当您创建新文档时，它们的事件不需要同时创建。事件应该创建并应用于响应用户操作(或系统事件)。
// *是否在文档列表中列出所有可能的操作由您决定。
// 我认为，对于给定用户在给定文档上的可能操作，最好在该文档的详细视图中提供(而不是在文档列表视图中)。
// *是的，当用户单击操作按钮时，我们可以创建相应的事件实例，并将其应用于文档的工作流。
// *使用“DocEventID”，您可以使用“DocEvents.Get”获取相应的事件实例

// Initialise DB connection.
// func init() {
// gt = t

// Connect to the database.travis
// driver, connStr := "mysql", "root:root@/flow"
// tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// flow.RegisterDB(tdb)

// if tdb == nil {
// 	log.Fatal("given database handle is `nil`")
// }
// db = tdb

// return nil
// }

// func RegisterDB(sdb *sql.DB) error {
// 	if sdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db = sdb

// 	return nil
// }wflist []*flow.DocState

// func (c *FlowController) testaddflow() {
// 	// driver, connStr := "mysql", "root:root@/flow"
// 	// tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	// if tdb == nil {
// 	// 	log.Fatal("given database handle is `nil`")
// 	// }
// 	// db := tdb
// 	driver, connStr := "sqlite3", "database/meritms.db"
// 	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	if tdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db := tdb
// 	tx, err := db.Begin()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// var tx *sql.Tx
// 	defer tx.Rollback()

// 	// docType1, err := flow.DocTypes.New(tx, "EXAM:COMMON")
// 	// DocTypes.New(tx, "Stor Request")
// 	// DocTypes.New(tx, "Compute Request")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// docState1, err := flow.DocStates.New(tx, "Init") //初始化
// 	// flow.DocStates.New(tx, "EntrustCreate")          //委托创建
// 	// flow.DocStates.New(tx, "EntrustApproved")        //委托审核
// 	// flow.DocStates.New(tx, "SampleHandon")           //样本交接
// 	// flow.DocStates.New(tx, "PrepareFinish")          //制备完成
// 	// flow.DocStates.New(tx, "PrepareApproved")        //制备审核
// 	// flow.DocStates.New(tx, "TaskAssign")             //任务分配
// 	// flow.DocStates.New(tx, "DataApproved")           //数据录入审核
// 	// flow.DocStates.New(tx, "ReportGen")              //报告生成
// 	// flow.DocStates.New(tx, "ReportApproved")         //报告审核

// 	// docActionID1, _ := flow.DocActions.New(tx, "CreateEntrust", false)  //创建委托
// 	// docActionID2, _ := flow.DocActions.New(tx, "ApproveEntrust", false) //审核委托
// 	// docActionID3, _ := flow.DocActions.New(tx, "HandonSample", false)   //提交样本
// 	// docActionID4, _ := flow.DocActions.New(tx, "FinishPrepare", true)   //完成制备
// 	// docActionID5, _ := flow.DocActions.New(tx, "ApprovePrepare", true)  //审核制备
// 	// docActionID6, _ := flow.DocActions.New(tx, "AssignTask", false)     //分配任务
// 	// docActionID7, _ := flow.DocActions.New(tx, "ApproveData", false)    //审核数据
// 	// docActionID8, _ := flow.DocActions.New(tx, "GenReport", false)      //生成报告
// 	// docActionID9, _ := flow.DocActions.New(tx, "ApproveReport", true)   //审核报告

// 	// workflowId, _ := flow.Workflows.New(tx, "Examination", docType1, docState1)

// 	// flow.Workflows.SetActive(tx, workflowId, true)

// 	//创建Docments
// 	// contextId, _ := flow.AccessContexts.New(tx, "Context")
// 	// groupId, err := flow.Groups.New(tx, "Examination", "G")
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }
// 	// resUser, _ := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
// 	// 	VALUES('admin', 'dashoo', 'admin@dashoo.com', 1)`)
// 	// uid, _ := resUser.LastInsertId()
// 	// userID1 := flow.UserID(uid)

// 	// flow.Groups.AddUser(tx, groupId, 5)
// 	// roleID1, _ := flow.Roles.New(tx, "administrator")
// 	flow.Roles.AddPermissions(tx, 5, 4, []flow.DocActionID{8, 9,
// 		10, 11, 12, 13, 14, 15, 16})

// 	docNewInput := flow.DocumentsNewInput{
// 		DocTypeID:       4, //docType1,
// 		AccessContextID: 1, //contextId,
// 		GroupID:         7,
// 		Title:           "entrust flow",
// 		Data:            "eid: 111, entrustNo: 2222",
// 	}
// 	// flow.Documents.New(tx, &docNewInput)

// 	_, err = flow.Documents.New(tx, &docNewInput)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// Documents.setState(tx, docType1, documentID1, docState2, contextId)
// 	tx.Commit()
// 	// wflist, err = flow.DocStates.List(0, 0)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }
// 	// return wflist
// }

// func (c *FlowController) flownext() {
// 	// tx, _ := flow.db.Begin()
// 	var tx *sql.Tx
// 	defer tx.Rollback()

// 	/*docNewInput := DocumentsNewInput {
// 	  	DocTypeID: docType1,
// 	  	AccessContextID: contextId,
// 	  	GroupID: groupId,
// 	  	Title: "entrust flow",
// 	  	Data: "eid: 111, entrustNo: 2222",
// 	  }
// 	  documentID1, err := Documents.New(tx, &docNewInput)
// 	  fmt.Println("documentID1=", documentID1, err)*/

// 	docEventInput := flow.DocEventsNewInput{
// 		DocTypeID:   flow.DocTypeID(4),
// 		DocumentID:  flow.DocumentID(1),
// 		DocStateID:  flow.DocStateID(9),
// 		DocActionID: flow.DocActionID(2),
// 		GroupID:     flow.GroupID(1),
// 		Text:        "开始审批",
// 	}
// 	groupIds := []flow.GroupID{flow.GroupID(1)}
// 	myWorkflow, err := flow.Workflows.Get(flow.WorkflowID(3))
// 	docEvent1, err := flow.DocEvents.New(tx, &docEventInput)
// 	tx.Commit()
// 	myDocEvent, err := flow.DocEvents.Get(docEvent1)
// 	newDocStateId, err := myWorkflow.ApplyEvent(tx, myDocEvent, groupIds)
// 	tx.Commit()
// 	fmt.Println("newDocStateId=", newDocStateId, err)
// }

//2019-01-11测试成功，作为保留
// func (c *FlowController) FlowGetDocTypeByName() {
// 	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
// 	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	if tdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db := tdb

// 	tx, err := db.Begin()
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// db.Close()
// 	//查询doctype
// 	dtID1, err := flow.DocTypes.GetByName("图纸设计")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	beego.Info(dtID1)
// 	// err = flow.DocTypes.AddTransition(tx, dtID1.ID, 7, 10, 8)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	myWorkflow, err := flow.Workflows.GetByName("图纸设计流程")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	beego.Info(myWorkflow)
// 	//查询context
// 	accessContextID1, err := flow.AccessContexts.List("Context", 0, 0)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	beego.Info(accessContextID1[0].ID)
// 	beego.Info(flow.GroupID(1))
// 	docNewInput := flow.DocumentsNewInput{
// 		DocTypeID:       dtID1.ID,
// 		AccessContextID: accessContextID1[0].ID,
// 		GroupID:         11, //groupId,
// 		Title:           "厂房布置图",
// 		Data:            "设计、制图: 秦晓川1, 校核: 秦晓川2",
// 	}
// 	// flow.Documents.New(tx, &docNewInput)
// 	DocumentID1, err := flow.Documents.New(tx, &docNewInput)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// tx.Commit() //new后面一定要跟commit啊
// 	beego.Info(DocumentID1)

// 	dsID1, err := flow.DocStates.GetByName("设计中...")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	beego.Info(dsID1)
// 	dsID2, err := flow.DocStates.GetByName("校核中...")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	beego.Info(dsID2)
// 	// _, err = flow.Workflows.AddNode(tx, dtID1.ID, dsID1.ID, accessContextID1[0].ID, myWorkflow.ID, "图纸三角校审流程-设计", flow.NodeTypeEnd)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	daID2, err := flow.DocActions.GetByName("提交设计")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	beego.Info(daID2)
// 	beego.Info(flow.GroupID(12))
// 	// docEventInput := flow.DocEventsNewInput{
// 	// 	DocTypeID:   dtID1.ID, //flow.DocTypeID(1),
// 	// 	DocumentID:  DocumentID1,
// 	// 	DocStateID:  dsID1.ID, //document state must be this state
// 	// 	DocActionID: daID2.ID, //flow.DocActionID(2),
// 	// 	GroupID:     12,
// 	// 	Text:        "校核",
// 	// }

// 	// docEventID1, err := flow.DocEvents.New(tx, &docEventInput)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }
// 	// // tx.Commit() //一个函数里只能有一个commit！！！！
// 	// beego.Info(docEventID1)
// 	myDocEvent, err := flow.DocEvents.Get(16)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	beego.Info(myDocEvent)
// 	// myWorkflow, err := flow.Workflows.Get(workflowId.ID)
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	//给出接受的组groupids
// 	groupIds := []flow.GroupID{flow.GroupID(13)}
// 	beego.Info(groupIds)
// 	newDocStateId, err := myWorkflow.ApplyEvent(tx, myDocEvent, groupIds)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	tx.Commit() //一个函数里只能有一个commit！！！！
// 	fmt.Println("newDocStateId=", newDocStateId, err)

// 	// wflist, err := flow.DocTypes.GetByName("Compute Request")
// 	// if err != nil {
// 	// 	beego.Error(err)
// 	// }

// 	c.Data["json"] = "OK" //wflist
// 	c.ServeJSON()
// }

// daID2, err := flow.DocActions.New(tx, "校核完成后提交", false)
// if err != nil {
// 	fmt.Println(err)
// }
// daID3, err := flow.DocActions.New(tx, "审查完成后提交", false)
// if err != nil {
// 	fmt.Println(err)
// }
// daID4, err := flow.DocActions.New(tx, "核定完成后提交", true)
// if err != nil {
// 	fmt.Println(err)
// }
// daID5, err := flow.DocActions.New(tx, "编制完成后提交", true)
// if err != nil {
// 	fmt.Println(err)
// }
// daID6, err := flow.DocActions.New(tx, "审批完成后提交", false)
// if err != nil {
// 	fmt.Println(err)
// }
// daID7, err := flow.DocActions.New(tx, "立项完成后提交", false)
// if err != nil {
// 	fmt.Println(err)
// }

// //添加流程规则1:oldstate1 action1 newstate2
// err = flow.DocTypes.AddTransition(tx, dtID1, dsID1, daID1, dsID2)
// if err != nil {
// 	beego.Error(err)
// }
// //添加流程规则2:oldstate2 action2 newstate3
// err = flow.DocTypes.AddTransition(tx, dtID1, dsID2, daID2, dsID3)
// if err != nil {
// 	beego.Error(err)
// }

// //定义流程类型doctype下的唯一流程workflow
// workflowID1, err := flow.Workflows.New(tx, "图纸设计-三级校审流程", dtID1, dsID1) //初始状态是“设计中...”——校核——审查——完成
// if err != nil {
// 	fmt.Println(err)
// }
// beego.Info(workflowID1)
// // workflowID2, err := flow.Workflows.New(tx, "图纸设计-二级校审流程", dtID1, dsID1) //初始状态是“设计中...”-“校核”——完成
// // if err != nil {
// // 	fmt.Println(err)
// // }
// // beego.Info(workflowID2)
// //定义合同评审下的流程类型：部门合同流程，总院合同流程
// //略

// //定义用户、组、角色、权限集合
// accessContextID1, err := flow.AccessContexts.New(tx, "Context")
// if err != nil {
// 	beego.Error(err)
// }

// //定义流程类型workflow下的具体每个节点node，用户对文件执行某个动作（event里的action）后，会沿着这些节点走
// // AddNode maps the given document state to the specified node.  This
// // map is consulted by the workflow when performing a state transition
// // of the system.nodeID1
// _, err = flow.Workflows.AddNode(tx, dtID1, dsID1, accessContextID1, workflowID1, "图纸设计-三级校审流程-设计", flow.NodeTypeBegin)
// if err != nil {
// 	fmt.Println(err)
// }
// _, err = flow.Workflows.AddNode(tx, dtID1, dsID2, accessContextID1, workflowID1, "图纸设计-三级校审流程-校核", flow.NodeTypeLinear)
// if err != nil {
// 	fmt.Println(err)
// }
// _, err = flow.Workflows.AddNode(tx, dtID1, dsID3, accessContextID1, workflowID1, "图纸设计-三级校审流程-审查", flow.NodeTypeEnd)
// if err != nil {
// 	fmt.Println(err)
// }
// //定义用户-组-角色-权限关系
// res, err := tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
// 		VALUES('秦', '晓川-1', 'email1@example.com', 1)`)
// if err != nil {
// 	log.Fatalf("%v\n", err)
// }
// uid, _ := res.LastInsertId()
// uID1 := flow.UserID(uid)
// _, err = flow.Groups.NewSingleton(tx, uID1)

// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
// 		VALUES('秦', '晓川-2', 'email2@example.com', 1)`)
// if err != nil {
// 	log.Fatalf("%v\n", err)
// }
// uid, _ = res.LastInsertId()
// uID2 := flow.UserID(uid)
// _, err = flow.Groups.NewSingleton(tx, uID2)

// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
// 		VALUES('秦', '晓川-3', 'email3@example.com', 1)`)
// if err != nil {
// 	log.Fatalf("%v\n", err)
// }
// uid, _ = res.LastInsertId()
// uID3 := flow.UserID(uid)
// _, err = flow.Groups.NewSingleton(tx, uID3)

// res, err = tx.Exec(`INSERT INTO users_master(first_name, last_name, email, active)
// 		VALUES('秦', '晓川-4', 'email4@example.com', 1)`)
// if err != nil {
// 	log.Fatalf("%v\n", err)
// }
// uid, _ = res.LastInsertId()
// uID4 := flow.UserID(uid)
// _, err = flow.Groups.NewSingleton(tx, uID4)

// gID1 := fatal1(flow.Groups.New(tx, "设计人员组", "G")).(flow.GroupID)
// gID2 := fatal1(flow.Groups.New(tx, "校核人员组", "G")).(flow.GroupID)
// fatal0(flow.Groups.AddUser(tx, gID1, uID1))
// fatal0(flow.Groups.AddUser(tx, gID1, uID2))
// fatal0(flow.Groups.AddUser(tx, gID1, uID3))

// fatal0(flow.Groups.AddUser(tx, gID2, uID2))
// fatal0(flow.Groups.AddUser(tx, gID2, uID3))
// fatal0(flow.Groups.AddUser(tx, gID2, uID4))
// roleID1 := fatal1(flow.Roles.New(tx, "设计人员角色")).(flow.RoleID)
// roleID2 := fatal1(flow.Roles.New(tx, "校核人员角色")).(flow.RoleID)
// //给角色role赋予action权限
// fatal0(flow.Roles.AddPermissions(tx, roleID1, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4}))
// fatal0(flow.Roles.AddPermissions(tx, roleID2, dtID1, []flow.DocActionID{daID1, daID2, daID3, daID4, daID5, daID6, daID7}))
// //给用户组group赋予角色role
// err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID1, roleID1)
// if err != nil {
// 	beego.Error(err)
// }
// //将group和role加到accesscontext里——暂时不理解
// err = flow.AccessContexts.AddGroupRole(tx, accessContextID1, gID2, roleID2)
// if err != nil {
// 	beego.Error(err) //UNIQUE constraint failed: wf_ac_group_roles.ac_id已修补
// }

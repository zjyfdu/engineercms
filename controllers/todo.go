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
	// "strings"
	// "fmt"
	// "reflect"
	// "sort"
	// "time"
)

// CMSTODO API
type TodoController struct {
	beego.Controller
}

// @Title post todo
// @Description post todo
// @Param name query string true "The name for todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /create [post]
func (c *TodoController) Create() {
	var userid int64
	openID := c.GetSession("openID")
	if openID == nil {
		// return false
		userid = 0
		//     this.SetSession("asta", int(1))
		//     this.Data["num"] = 0
	} else {
		user, err := models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
		userid = user.Id
	}

	name := c.Input().Get("name")

	todoid, err := models.TodoCreate(name, userid)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"message": "写入数据库出错"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"message": "ok", "todoid": todoid}
		c.ServeJSON()
	}
}

// @Title get todolist
// @Description get tolist
// @Param page query string false "The page of todo"
// @Param limit query string false "The size of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /gettodo [get]
//取出所有未完成待办
func (c *TodoController) GetTodo() {
	var offset, limit1, page1 int
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 20
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		// limit1 = 10
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	todos, err := models.GetTodoUser(limit1, offset)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = todos //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title update todolist
// @Description update tolist
// @Param todoid query string false "The id of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /updatetodo [post]
//取出所有未完成待办
func (c *TodoController) UpdateTodo() {
	id := c.Input().Get("todoid")
	//pid转成64为
	todoid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateTodo(todoid)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"message": "更新数据库错误"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"todoid": id, "message": "ok"}
		c.ServeJSON()
	}
}

// @Title delete todolist
// @Description delete tolist
// @Param todoid query string false "The id of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /deletetodo [post]
//取出所有未完成待办
func (c *TodoController) DeleteTodo() {
	id := c.Input().Get("todoid")
	//pid转成64为
	todoid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteTodo(todoid)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"message": "删除数据错误"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"todoid": id, "message": "ok"}
		c.ServeJSON()
	}
}

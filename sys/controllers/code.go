package controllers

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"sys/models"
)

type CodeController struct {
	beego.Controller
}

func (c *CodeController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
}

// GetOne ...
// @Title Get One
// @Description get Code by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Code
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CodeController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	fmt.Println(idStr)
	id,_ := strconv.Atoi(idStr)
	v, err := models.GetCodeById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @router / [get]
func (c *CodeController) GetAll() {
	var fields []string
	var sortby []string
	var order  []string
	var query  = make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v!= "" {
		fields = strings.Split(v,",")
	}

	//int 默认是int64
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}

	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}

	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v,",")
	}

	if v := c.GetString("order"); v != "" {
		order = strings.Split(v,",")
	}

	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			//将query内容分割成2个子串,多余不再分割
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error:invalid query key/value paid")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query,fields,sortby,order,offset,limit)
	if err != nil {
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = l
	}
	c.ServeJSON()
}




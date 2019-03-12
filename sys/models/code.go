package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Code struct {
	Id 			int
	Channel 	uint8
	Order   	string
	TicketId 	int
	Mobile  	string
	Code    	string
	Status      uint8
	SendStatus  uint8
	SendResult  string
	CreateTime  int
	UpdateTime  int
	SendTime    int
	UserId      int
}

//注册model
func init() {
	orm.RegisterModel(new(Code))
}

//指定表名
func (c *Code) TableName() string {
	return "a_box_code"
}

//获取单条数据
func GetCodeById(id int) (c *Code, err error) {
	o := orm.NewOrm()
	c = &Code{Id: id}
	//Read 默认通过查询主键赋值，可以使用指定的字段进行查询
	if err = o.Read(c); err == nil {
		return c, nil
	}
	return nil, err
}

//获取多条数据
func GetAllCode(query map[string]string, fields []string, sortby []string, order []string, offset int64,
	limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Code))

	for k, v := range query {
		// 字段的分隔符号使用双下划线 __，除了描述字段
		k = strings.Replace(k,".","__",-1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v =="true" || v == "1"))
		}else{
			qs = qs.Filter(k, v)
		}
	}

	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {

		} else if len(sortby) != len(order) && len(order) == 1 {

		} else if len(sortby) != len(order) && len(order) != 1 {

		}
	}

}


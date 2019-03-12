package init

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"time"
)

func initDb(){
	//	注册数据库
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("MysqlDsn"))

	//  开启日志记录
	if beego.BConfig.RunMode != "pro" {
		var (
			w io.Writer
			err error
		)

		sqlLog := "logs/sql_"+time.Now().Format("20060102")+".log"
		w, err = os.OpenFile(sqlLog,os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_RDWR, 0666)
		if err != nil {
			w = nil
			fmt.Println("open sql log failure. errmsg:"+err.Error())
			os.Exit(1)
		}
		orm.Debug = true
		orm.DebugLog = orm.NewLog(w)
	}
}

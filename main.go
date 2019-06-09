package main

import (
	"app"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "homestead:secret@tcp(192.168.10.10:3306)/october?charset=utf8")
}

func main() {

	app.Run()
}

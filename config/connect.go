package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DbConn *sql.DB
	err    error
)

func init(){
	//连接数据库
	constr := "root:141592@tcp(127.0.0.1:3306)/cat_miniapp?charset=utf8"
	//打开连接
	db ,err := sql.Open("mysql",constr) //返回mysql实例db
	if err != nil {
		log.Panic(err.Error())
		return
	}
	DbConn = db
}



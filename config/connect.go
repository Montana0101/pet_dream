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

func init() {
	//连接数据库
	constr := "montana:@141592qb@tcp(rm-bp1076k66l7izw64nmo.mysql.rds.aliyuncs.com:3306)/pet_wechat?charset=utf8mb4"
	//打开连接
	db, err := sql.Open("mysql", constr) //返回mysql实例db
	if err != nil {
		log.Panic(err.Error())
		return
	}
	DbConn = db
}

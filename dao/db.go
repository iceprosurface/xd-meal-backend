package dao

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	userName = os.Getenv("DATABASE_USERNAME")
	password = os.Getenv("DATABASE_PASSWORD")
	ip       = os.Getenv("DATABASE_IP")
	port     = os.Getenv("DATABASE_PORT")
	dbName   = os.Getenv("DATABASE_NAME")
)

//Db数据库连接池
var DB *sql.DB

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
}

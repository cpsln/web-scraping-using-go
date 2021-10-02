package dbconnection


import (
	// "fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)



func ConnectDatabase() {
	// var err error
	orm.RegisterDriver("mysql", orm.DRMySQL)

	connectionString := "root:root@tcp(localhost:3306)/testdb?charset=utf8"

	orm.RegisterDataBase("default", "mysql", connectionString)
}

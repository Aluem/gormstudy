package gormstudy

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetBiReportDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// interpolateParams
		DSN: "root:123456@tcp(10.1.69.186:3306)/boss_bi_report?charset=utf8&interpolateParams=true&parseTime=True&loc=Local", // DSN data source name
		// DSN: "root:123456@tcp(10.1.69.186:3306)/boss_bi_report?charset=utf8&parseTime=True&loc=Local", // DSN data source name

	}), &gorm.Config{
		// PrepareStmt: true,
	})
	if err != nil {
		return nil, fmt.Errorf("GetBiReportDb gorm open err:%v", err)
	}

	return db, nil
}

// 定义一个全局对象db
var gdb *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(10.1.69.186:3306)/boss_bi_report?charset=utf8&parseTime=True&loc=Local&interpolateParams=true"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	gdb, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = gdb.Ping()
	if err != nil {
		return err
	}
	return nil
}

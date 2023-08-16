package gormstudy

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetBiReportDb() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(10.1.69.186:3306)/boss_bi_report?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("GetBiReportDb gorm open err:%v", err)
	}

	return db, nil
}

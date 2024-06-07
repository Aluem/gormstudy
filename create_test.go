package gormstudy

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestSQLExpressionMap(t *testing.T) {

	type DeclaringModelsTest struct {
		Id     int `gorm:"primaryKey"`
		Delete int64
	}

	var datas []map[string]interface{}

	for i := 0; i < 50000; i++ {
		data := make(map[string]interface{})
		data["Name"] = fmt.Sprintf("test%v", i)

		datas = append(datas, data)
	}

	// Create from map
	err := db.Table("declaring_models_test").Debug().CreateInBatches(&datas, 500).Error
	if err != nil {
		t.Fatal(err)
	}
	// INSERT INTO `declaring_models_test` (`Delete`,`Name`) VALUES (unix_timestamp('2020-12-22 00:00:00'),'jinzhu')

	fmt.Println("success")

}

type CurTime struct {
	Time string
}

// Scan implements the sql.Scanner interface
func (loc *CurTime) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	loc.Time = fmt.Sprintf("%v", v)
	return nil
}

func (loc CurTime) GormDataType() string {
	return "geometry"
}

func (loc CurTime) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "unix_timestamp(?)",
		Vars: []interface{}{fmt.Sprintf("%s", loc.Time)},
	}
}

func TestSQLExpressionStruct(t *testing.T) {
	type DeclaringModelsTest struct {
		Id     int `gorm:"primaryKey"`
		Delete CurTime
	}

	db, err := GetBiReportDb()
	if err != nil {
		t.Fatal(err)
	}

	createM := DeclaringModelsTest{Delete: CurTime{
		Time: "2020-12-22 00:00:00",
	}}
	err = db.Table("declaring_models_test").Debug().Create(&createM).Error
	if err != nil {
		t.Fatal(err)
	}

	var m DeclaringModelsTest
	err = db.Table("declaring_models_test").Debug().First(&m, createM.Id).Error
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("get:", m)

	fmt.Println("success")

}

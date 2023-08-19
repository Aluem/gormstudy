package gormstudy

import (
	"context"
	"fmt"
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Test1(t *testing.T) {
	db, err := GetBiReportDb()
	if err != nil {
		t.Fatal(err)
	}

	var id int64
	err = db.Debug().Table("glaze_querys").Select("id").Where("id=1").Find(&id).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("success")
	fmt.Println("id=", id)
}

func TestCreatingTracking(t *testing.T) {
	type DeclaringModelsTest struct {
		Id        int
		Name      string
		CreatedAt int64
		UpdatedAt int64
		DeletedAt int64

		Create int64 `gorm:"autoCreateTime"`
		Update int64 `gorm:"autoUpdateTime"`
		Delete int64 `gorm:"autoDeleteTime"`
	}

	db, err := GetBiReportDb()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Debug().Table("declaring_models_test").Create(&DeclaringModelsTest{Name: "zhangsan"}).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("success")

}

func TestEmbeddedStruct(t *testing.T) {
	type Author struct {
		Name string
	}

	type DeclaringModelsTest struct {
		Author Author `gorm:"embedded;embeddedPrefix:author_"`
	}

	db, err := GetBiReportDb()
	if err != nil {
		t.Fatal(err)
	}

	err = db.Debug().Table("declaring_models_test").Create(&DeclaringModelsTest{Author: Author{
		Name: "aluemhe",
	}}).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("success")

}

func TestSQLExpressionMap(t *testing.T) {

	type DeclaringModelsTest struct {
		Id     int `gorm:"primaryKey"`
		Delete int64
	}

	db, err := GetBiReportDb()
	if err != nil {
		t.Fatal(err)
	}

	// Create from map
	db.Table("declaring_models_test").Debug().Create(map[string]interface{}{
		"Name":   "jinzhu",
		"Delete": clause.Expr{SQL: "unix_timestamp(?)", Vars: []interface{}{"2020-12-22 00:00:00"}},
	})
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

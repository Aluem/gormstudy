package gormstudy

import (
	"fmt"
	"testing"
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

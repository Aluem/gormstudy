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

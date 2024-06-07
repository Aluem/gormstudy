package gormstudy

import (
	"fmt"
	"testing"

	"gorm.io/gorm/clause"
)

func TestUpdateSelect(t *testing.T) {

	db, err := GetBiReportDb()
	if err != nil {
		t.Fatal(err)
	}

	d := DeclaringModelsTest{Name: "new_name2", CreatedAt: 2}
	// Create from map
	// err = db.Model(&DeclaringModelsTest{}).Where("id = 19").Updates(&d).Error
	err = db.Model(&DeclaringModelsTest{}).Debug().Clauses(clause.Returning{Columns: []clause.Column{{Name: "created_at"}}}).Where("id = 19").Updates(&d).Error
	// err = db.Model(&DeclaringModelsTest{}).Where("id = 19").Select("name").Updates(&DeclaringModelsTest{Name: "new_name", CreatedAt: 1}).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("d.CreatedAt:", d.CreatedAt)

	fmt.Println("success")

}

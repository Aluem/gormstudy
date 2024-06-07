package gormstudy

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db, _ = GetBiReportDb()
}

type DeclaringModelsTest struct {
	Id        int
	Name      string
	CreatedAt int64 // `gorm:"column:created_at"`
}

func (DeclaringModelsTest) TableName() string {
	return "declaring_models_test"
}

func TestGolangDb(t *testing.T) {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	sqlStr := "select id from declaring_models_test where id in("
	var args []any
	var ps []string
	for i := 0; i < 30000; i++ {
		ps = append(ps, "?")
		args = append(args, i)
	}
	sqlStr += strings.Join(ps, ",")
	sqlStr += ")"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	rows, err := gdb.Query(sqlStr, args...)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var alb DeclaringModelsTest
		if err := rows.Scan(&alb.Id); err != nil {
			t.Fatal(err)
		}
		fmt.Printf("id:%d \n", alb.Id)
	}
	if err := rows.Err(); err != nil {
		t.Fatal(err)
	}

}

func TestInject(t *testing.T) {
	// 测试是否有sql注入风险
	fmt.Println()
	var d []DeclaringModelsTest
	err := db.Debug().Where("id in(?)", "1') or name = 't1t1' or id in ('1").Find(&d).Error // 占位符替换，没有风险
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("res:", d)

	err = db.Debug().Where(fmt.Sprintf("id in(%v)", "1) or name = 't1t1' or id in (1")).Find(&d).Error // 直接拼接，有风险
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("res2:", d)
}

func TestFirstOrInit(t *testing.T) {
	fmt.Println(12232)
	var d []DeclaringModelsTest

	now := time.Now().Unix()
	var ids []string
	for i := 0; i < 30000; i++ {
		ids = append(ids, strconv.Itoa(i))
	}
	// err := db.Debug().Where("id in(?)", strings.Join(ids, ",")).Find(&d).Error
	// err := db.Debug().Where("id in(?)", ids).Find(&d).Error
	err := db.Debug().Where("id in(?)", ids).Find(&d).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("查询用时1：", time.Now().Unix()-now, "s")
	fmt.Println(d)

	ids = nil
	for i := 30000; i < 60000; i++ {
		ids = append(ids, strconv.Itoa(i))
	}

	err = db.Debug().Where("id in(?)", ids).Find(&d).Error
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("查询用时2：", time.Now().UnixMilli()-now, "ms")
	fmt.Println(d)

}

func Test102(t *testing.T) {
	var arr []*int
	for i := 0; i < 10; i++ {
		j := i
		arr = append(arr, &j)
	}

	var arr2 []*int
	for _, i := range arr {
		arr2 = append(arr2, i)
	}

	fmt.Println(arr2)
}

func Test03(t *testing.T) {
	err := setModel(db).Where("id=2").Update("name", "t1t1").Error
	if err != nil {
		t.Fatal(err)
	}

	err = db.Where("id=3").Update("name", "t1t1").Error
	if err != nil {
		t.Fatal(err)
	}
}

func setModel(db *gorm.DB) *gorm.DB {
	return db.Model(&DeclaringModelsTest{})
}

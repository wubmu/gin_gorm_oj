package test

import (
	"fmt"
	"gin_gorm_oj/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

//
//  TestGormTet
//  @Description: 测试GORM
//  @param t
//
//func TestModelsGormTest(t *testing.T) {
//	models.GetProblemList()
//}
func TestGormTet(t *testing.T) {

	//	测试GORM
	//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.ProblemBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		fmt.Printf("Problem ==> %v \n\n", v)
	}

}

//func main() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//}

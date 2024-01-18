package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 16:46
func main() {
	dsn := "admin:123456@tcp(192.168.1.52:3306)/yami_shops?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tz_", // 定义表前缀，没有全局定义时，生成时需要输写完整表名
			SingularTable: true,  // 统一采用单数的时候定义结构名
		},
	})

	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dao",                                                     // 定义 dao 文件输出目录
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "internal/entity",                                                  // 定义 model 文件输出目录
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

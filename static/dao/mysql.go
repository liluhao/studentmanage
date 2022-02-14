package dao

import (
	"bluebell/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var AdminTable, StudentTable, DB *gorm.DB

func MysqlInit() {
	conn, err := gorm.Open(mysql.New(mysql.Config{
		//需要使用者在本地创建一个名字叫做”student“的数据库；414524需要使用者更改为自己本地Mysql的密码
		DSN:               "root:414524@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // love表将是love，不再是loves，即可成功取消表明被加s
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	AdminTable = conn
	StudentTable = conn
	err = AdminTable.AutoMigrate(&model.Admin{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = StudentTable.AutoMigrate(&model.Student{})
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

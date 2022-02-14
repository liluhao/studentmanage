package model

import (
	_ "gorm.io/driver/mysql"
)

//学生表
type Student struct {
	Name     string
	Password string
	Sex      string
}

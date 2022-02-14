package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"studentmanage/dao"
	"studentmanage/model"
)

var students []model.Student
var student model.Student
var admins []model.Admin
var admin model.Admin

func Create(c *gin.Context) {
	admin.Name = c.PostForm("adminname")
	admin.Password = c.PostForm("adminpassword")
	res := dao.AdminTable.Create(&admin)
	if res.Error != nil {
		c.String(http.StatusOK, "注册管理员失败")
	} else {
		c.String(http.StatusOK, "注册管理员成功")
	}
}

func AdminLogin(c *gin.Context) {
	admin.Name = c.PostForm("adminname")
	admin.Password = c.PostForm("adminpassword")
	res := dao.StudentTable.Where("name=? AND password=?", admin.Name, admin.Password).First(&admin) //如果没匹配到的话,会打印出错误的
	if res.Error != nil {
		c.String(http.StatusOK, "登录管理员失败")
	} else {
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		c.Redirect(http.StatusMovedPermanently, "/admin/manage")
	}
}

func AddAdmin(c *gin.Context) {
	admin.Name = c.PostForm("adminname")
	admin.Password = c.PostForm("adminpassword")
	res := dao.AdminTable.Create(&admin)
	if res.Error != nil {
		c.String(http.StatusOK, "插入管理员失败")
	} else {
		c.String(http.StatusOK, "插入管理员成功")
	}

}
func DelAdmin(c *gin.Context) {
	admin.Name = c.PostForm("adminname")
	res := dao.AdminTable.Unscoped().Where("name = ?", admin.Name).Delete(&admin) //如果没匹配到的话,不会打印出错误的
	if res.Error != nil || res.RowsAffected == 0 {
		c.String(http.StatusOK, "删除指定管理员失败")
	} else {
		c.String(http.StatusOK, "删除指定管理员成功")
	}
}

func UpdateAdmin(c *gin.Context) {
	admin.Name = c.PostForm("adminname")
	admin.Password = c.PostForm("adminpassword")

	res := dao.AdminTable.Model(&model.Admin{}).Where("name LIKE ?", admin.Name).Update("password", admin.Password)
	if res.Error != nil || res.RowsAffected == 0 {
		c.String(http.StatusOK, "更新管理员密码失败")
	} else {
		c.String(http.StatusOK, "更新管理员密码成功")
	}
}

func AddStudent(c *gin.Context) {
	student.Name = c.PostForm("studentname")
	student.Password = c.PostForm("studentpassword")
	student.Sex = c.PostForm("studentsex")
	res := dao.StudentTable.Create(&student)
	if res.Error != nil {
		c.String(http.StatusOK, "注册学生失败")
	} else {
		c.String(http.StatusOK, "注册学生成功")
	}
}

func DelStudent(c *gin.Context) {
	student.Name = c.PostForm("studentname")
	res := dao.StudentTable.Unscoped().Where("name = ?", student.Name).Delete(&student) //如果没匹配到的话,不会打印出错误的
	if res.Error != nil || res.RowsAffected == 0 {
		c.String(http.StatusOK, "删除学生失败")
	} else {
		c.String(http.StatusOK, "删除学生成功")
	}

}

func UpdateStudent(c *gin.Context) {
	student.Name = c.PostForm("studentname")
	student.Password = c.PostForm("studentpassword")

	res := dao.StudentTable.Model(&model.Student{}).Where("name LIKE ?", student.Name).Update("password", student.Password) //如果没匹配到的话,不会打印出错误的
	if res.Error != nil || res.RowsAffected == 0 {
		c.String(http.StatusOK, "更新学生密码失败")
	} else {
		c.String(http.StatusOK, "更新学生密码成功")
	}
}

func GetThisStudent(c *gin.Context) {
	Name := c.PostForm("studentname")
	res := dao.StudentTable.Where("name=?", Name).First(&students) //如果没匹配到的话,不会打印出错误的
	if res.Error != nil || res.RowsAffected == 0 {
		c.String(http.StatusOK, "获取学生信息失败")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Name": Name,
			"Data": students,
		})
	}
}

func GetStudent(c *gin.Context) {
	res := dao.StudentTable.Find(&students)
	if res.Error != nil {
		c.String(http.StatusOK, "获取所有信息失败")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Data": students,
		})
	}
}

func StudentLogin(c *gin.Context) {
	Name := c.PostForm("studentname")
	Password := c.PostForm("studentpassword")
	res := dao.StudentTable.Where("name=? AND password=?", Name, Password).First(&student)
	if res.Error != nil {
		c.String(http.StatusOK, "登良失败")
	} else {
		c.String(http.StatusOK, "登录成功")
	}
}

func AddStudentAgain(c *gin.Context) {
	student.Name = c.PostForm("studentname")
	student.Password = c.PostForm("studentpassword")
	student.Sex = c.PostForm("studentsex")
	res := dao.StudentTable.Create(&student)
	if res.Error != nil {
		c.String(http.StatusOK, "注册学生失败")
	} else {
		c.String(http.StatusOK, "注册学生成功")
	}
}

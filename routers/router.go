package routers

import (
	"github.com/gin-gonic/gin"
	"studentmanage/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("static/login.html", "static/manage.html") //加载模板文件
	r.GET("/web", controller.Login)
	r.POST("/admin/create", controller.Create)
	r.POST("/admin/login", controller.AdminLogin)
	adminGroup := r.Group("admin", controller.Cookie)
	{
		adminGroup.POST("/addAdmin", controller.AddAdmin)
		adminGroup.POST("/delAdmin", controller.DelAdmin)
		adminGroup.POST("/updateAdmin", controller.UpdateAdmin)
		adminGroup.GET("/manage", controller.Manage)

		adminGroup.POST("/addStudent", controller.AddStudent)
		adminGroup.POST("/delStudent", controller.DelStudent)
		adminGroup.POST("/updateStudent", controller.UpdateStudent)
		adminGroup.GET("/getStudent", controller.GetStudent)
		adminGroup.POST("/getThisStudent", controller.GetThisStudent)
	}
	studentGroup := r.Group("student")
	{
		studentGroup.POST("/studentLogin", controller.StudentLogin)
		studentGroup.POST("/addStudent", controller.AddStudentAgain)
	}
	return r
}

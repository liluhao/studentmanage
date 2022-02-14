package main

import (
	"studentmanage/dao"
	"studentmanage/routers"

	_ "github.com/gin-gonic/gin"
)

func main() {
	dao.MysqlInit()
	r := routers.SetupRouter()
	r.Run(":9000")
}

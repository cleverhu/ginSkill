package main

import (
	"ginSkill/src/common"
	_ "ginSkill/src/dbs"
	"ginSkill/src/handlers"
	_ "ginSkill/src/validators"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.New()
	r.Use(common.ErrorHandler())

	//查询用户列表
	r.GET("/users", handlers.UserList)

	//查询单个用户
	r.GET("/users/:id", handlers.UserDetail)

	//新增用户
	r.PUT("/users", handlers.AddUser)

	//修改用户
	r.POST("/users", handlers.UpdateUser)

	//删除用户
	r.DELETE("/users/:id", handlers.DeleteUser)

	log.Fatal(r.Run(":80"))
}

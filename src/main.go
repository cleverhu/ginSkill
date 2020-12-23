package main

import (
	"fmt"
	"ginSkill/src/common"
	_ "ginSkill/src/dbs"
	"ginSkill/src/handlers"
	"ginSkill/src/models/UserModel"
	"ginSkill/src/services/UserService"
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

func testInsert() {
	for i := 0; i < 100; i++ {
		u := UserModel.New().
			Mutate(UserModel.WithUsername("user" + fmt.Sprintf("%d", i))).
			Mutate(UserModel.WithPassword("password" + fmt.Sprintf("%d", i))).
			Mutate(UserModel.WithEmail("email" + fmt.Sprintf("%d", i) + "@qq.com"))
		UserService.Insert(u)
	}
}

func testDel() {
	for i := 0; i < 20; i++ {
		b := UserService.DeleteByID(i)
		fmt.Println(b)
	}
}

func testQuery(id int) {
	u := UserService.QueryByID(id)
	fmt.Println(u)
	fmt.Println(u.Username)
}

func testPageQuery(page, size int) {
	us := UserService.QueryPage(page, size)
	fmt.Println(us)
	for _, v := range us {
		fmt.Println(v)
	}
}

func testUpdate() {
	u := UserModel.New().
		Mutate(UserModel.WithUserID(22)).
		Mutate(UserModel.WithUsername("changed22"))
	fmt.Println(u)
	fmt.Println(UserService.Update(u))

	fmt.Println(UserService.QueryByID(u.ID))
}

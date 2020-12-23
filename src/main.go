package main

import (
	"fmt"
	"ginSkill/src/models/UserModel"
	"ginSkill/src/services/UserService"
)

func main() {
	//testDel()
	//testInsert()
	//testQuery(100)
	//testPageQuery(1, 100)
	//testUpdate()



	//r := gin.New()
	//
	//r.GET("/", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{"user": "user"})
	//})
	//r.PUT("/", func(ctx *gin.Context) {
	//	u := UserModel.New()
	//	err := ctx.ShouldBindJSON(&u)
	//
	//})
	//
	//
	//
	//log.Fatal(r.Run(":80"))
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

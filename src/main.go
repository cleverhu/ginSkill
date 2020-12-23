package main

import (
	"fmt"
	"ginSkill/src/models/UserModel"
)

func main() {
	u := UserModel.New().
		Mutate(UserModel.WithUserID(1)).
		Mutate(UserModel.WithPassword("123456")).
		Mutate(UserModel.WithUsername("Lily"))
	fmt.Println(u)
	//r := gin.New()
	//r.GET("/", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{"user": "user"})
	//})
	//
	//log.Fatal(r.Run(":80"))
}

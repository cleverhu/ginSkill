package handlers

import (
	"ginSkill/src/data/Getter"
	"ginSkill/src/data/Setter"
	"ginSkill/src/models/UserModel"
	"ginSkill/src/result"
	"github.com/gin-gonic/gin"
	"strconv"
)

//取用户列表
func UserList(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}

	sizeStr := ctx.DefaultQuery("size", "10")
	size, _ := strconv.Atoi(sizeStr)
	if size <= 0 {
		size = 10
	}

	R(ctx)("query users success", "10000", Getter.UserGetter.GetUserListByPage(page, size))(OK)
}

func UserDetail(ctx *gin.Context) {
	id := &struct {
		Id int `uri:"id" binding:"required,gt=10"`
	}{}
	result.Result(ctx.ShouldBindUri(id)).Unwrap()
	R(ctx)("userDetail", "10001", Getter.UserGetter.GetUserByID(id.Id).Unwrap())(OK)
}

func DeleteUser(ctx *gin.Context) {
	id := &struct {
		Id int `uri:"id" binding:"required,gt=10"`
	}{}
	result.Result(ctx.ShouldBindUri(id)).Unwrap()
	R(ctx)("deleteUser", "10002", Setter.UserSetter.DeleteUserByID(id.Id).Unwrap())(OK)

}

func AddUser(ctx *gin.Context) {
	u := UserModel.New()
	result.Result(ctx.ShouldBindJSON(u)).Unwrap()
	R(ctx)("addUser", "10003", Setter.UserSetter.AddUser(u).Unwrap())(OK)
}

func UpdateUser(ctx *gin.Context) {
	u := UserModel.New()
	result.Result(ctx.ShouldBindJSON(u)).Unwrap()
	R(ctx)("updateUser", "10004", Setter.UserSetter.UpdateUser(u).Unwrap())(OK)
}

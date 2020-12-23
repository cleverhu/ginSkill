package Setter

import (
	"fmt"
	"ginSkill/src/dbs"
	"ginSkill/src/models/UserModel"
	"ginSkill/src/result"
	"github.com/gin-gonic/gin"
	"time"
)

var UserSetter IUserSetter

func init() {

	UserSetter = NewUserSetterImpl()
}

type IUserSetter interface {
	DeleteUserByID(id int) *result.ErrorResult
	AddUser(u *UserModel.UserModelImpl) *result.ErrorResult
	UpdateUser(u *UserModel.UserModelImpl) *result.ErrorResult
}

type UserSetterImpl struct {
}

func NewUserSetterImpl() *UserSetterImpl {
	return &UserSetterImpl{}
}

func (this *UserSetterImpl) DeleteUserByID(id int) *result.ErrorResult {
	u := UserModel.New(UserModel.WithUserID(id))
	if dbs.Orm.First(&u).RecordNotFound() == false && dbs.Orm.Delete(&u).RowsAffected == 1 {
		return &result.ErrorResult{
			Err:  nil,
			Data: gin.H{"success": true, "user": u},
		}
	} else {
		return &result.ErrorResult{
			Err:  fmt.Errorf("delete failed,not found id:%d", id),
			Data: nil,
		}
	}
}

func (this *UserSetterImpl) AddUser(u *UserModel.UserModelImpl) *result.ErrorResult {
	u.UpdateTime = time.Now()
	if dbs.Orm.First(UserModel.New(), u.ID).RecordNotFound() == true && dbs.Orm.Save(&u).RowsAffected == 1 {
		return &result.ErrorResult{
			Err:  nil,
			Data: gin.H{"success": true, "user": u},
		}
	} else {
		return &result.ErrorResult{
			Err:  fmt.Errorf("add user failed,user is existed"),
			Data: nil,
		}
	}
}

func (this *UserSetterImpl) UpdateUser(u *UserModel.UserModelImpl) *result.ErrorResult {
	if dbs.Orm.First(UserModel.New(), u.ID).RecordNotFound() == false {
		dbs.Orm.Table("t_user").Where("u_id = ?", u.ID).Update(&u)
		return &result.ErrorResult{
			Err:  nil,
			Data: gin.H{"success": true, "user": u},
		}
	} else {
		return &result.ErrorResult{
			Err:  fmt.Errorf("update user failed,user not existes"),
			Data:nil,
		}
	}
}

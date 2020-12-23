package test

import (
	"fmt"
	"ginSkill/src/models/UserModel"
	"ginSkill/src/result"
	"github.com/gin-gonic/gin"
)

func TestUs(us []*UserModel.UserModelImpl) *result.ErrorResult {
	if len(us) != 10 {
		return &result.ErrorResult{
			Err:  fmt.Errorf("test error"),
			Data: nil,
		}
	} else {
		return &result.ErrorResult{
			Err:  nil,
			Data: gin.H{"message": "test success"},
		}
	}
}

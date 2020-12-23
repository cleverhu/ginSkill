package common

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				ctx.JSON(400, gin.H{"message": e})
			}
		}()
		ctx.Next()
	}
}

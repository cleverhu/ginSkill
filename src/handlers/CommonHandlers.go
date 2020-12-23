package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

type JSONResult struct {
	Message string
	Code    string
	Result  interface{}
}

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return gin.H{
				"message": "",
				"code":    "",
				"result":  nil,
			}
		},
	}
}

type ResultFunc func(msg, code string, result interface{}) func(output Output)
type Output func(ctx *gin.Context, v interface{})

func R(ctx *gin.Context) ResultFunc {
	return func(msg, code string, result interface{}) func(output Output) {
		r := ResultPool.Get().(gin.H)
		defer ResultPool.Put(r)
		r["message"] = msg
		r["code"] = code
		r["result"] = result
		//ctx.JSON(400, r)
		return func(output Output) {
			output(ctx, r)
		}
	}
}

func OK(ctx *gin.Context, v interface{}) {
	ctx.JSON(200, v)
}

func Error(ctx *gin.Context, v interface{}) {
	ctx.JSON(400, v)
}

func OK2String(ctx *gin.Context, v interface{}) {
	ctx.String(200, fmt.Sprintf("%v", v))
}

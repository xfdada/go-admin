package response

import (
	"github.com/gin-gonic/gin"
	"go-admin/global/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// 成功返回
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

//失败返回
func (r *Response) ToError(err *errcode.Error) {
	response := gin.H{"code": err.Codes(), "msg": err.Msgs()}
	r.Ctx.JSON(err.StatusCode(), response)
}

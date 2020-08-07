package democontroller

import (
	"net/http"

	"gaea/app/service/demoservice"
	"gaea/utils"
	"github.com/gin-gonic/gin"
)

//accept http request
func GaeaDemo(ctx *gin.Context) {
	goCtx := utils.TransferToContext(ctx)
	param := ctx.PostForm("param")
	ret, err := demoservice.DoFun(goCtx, param)
	if err != nil {
		resp := utils.Error(err)
		ctx.JSON(http.StatusOK, resp)
	} else {
		resp := utils.Success(ret)
		ctx.JSON(http.StatusOK, resp)
	}
}

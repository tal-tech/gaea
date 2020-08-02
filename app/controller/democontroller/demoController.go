package democontroller

import (
	"net/http"

	"gaea/app/service/demoservice"
	"git.100tal.com/wangxiao_go_lib/xesGoKit/xesgin"
	logger "git.100tal.com/wangxiao_go_lib/xesLogger"
	"github.com/gin-gonic/gin"
)

//accept http request
func GaeaDemo(ctx *gin.Context) {
	goCtx := xesgin.TransferToContext(ctx)
	param := ctx.PostForm("param")
	ret, err := demoservice.DoFun(goCtx, param)
	if err != nil {
		resp := xesgin.Error(err)
		ctx.JSON(http.StatusOK, resp)
	} else {
		resp := xesgin.Success(ret)
		ctx.JSON(http.StatusOK, resp)
	}
}

//accept websocket request
func WebSocketDemo(ctx *gin.Context) {
	goCtx := xesgin.TransferToContext(ctx)
	ws, err := xesgin.GetWebSocketConn(ctx)
	if err != nil {
		logger.Ex(goCtx, "WebSocketDemo", "GetWebSocketConn err:%v", err)
		resp := xesgin.Error(err)
		ctx.JSON(http.StatusOK, resp)
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, e := ws.ReadMessage()
		if e != nil {
			err = e
			logger.Ex(goCtx, "WebSocketDemo", "ReadMessage err:%v", err)
			break
		}
		if string(message) == "hello" {
			message = []byte("hello, i am gaea.")
		}
		//写入ws数据
		e = ws.WriteMessage(mt, message)
		if e != nil {
			err = e
			logger.Ex(goCtx, "WebSocketDemo", "WriteMessage err:%v", err)
			break
		}
	}
	if err != nil {
		resp := xesgin.Error(err)
		ctx.JSON(http.StatusOK, resp)
	}
}

# Gaea
## Introduction
`Gaea` is a Gin-based web framework. In actual work, there are a series of additional engineering issues that need to be resolved before the project is applied to the production environment.
Otherwise, the portability of the system will be poor, and developers will not be able to focus on business development.
Incorporated into a complete set of solutions: dependency management, configuration management, compilation and deployment, monitoring & alarms, and support for one-click quick construction of web applications. If you are considering writing a web server in Golang, then Gaea is undoubtedly your best choice!

## Quick Start
Build & Run
```golang
//Recommended  $GOPATH/src  as your workspace
$ cd $GOPATH/src/

//clone the framework to local
$ git clone https://github.com/tal-tech/gaea.git

//Will use makefile to compile and generate binary files to the bin directory
$ make 
```

## Example
1.Config server port
```golang
//conf/conf.ini
[HTTP]
port = 9898
;...
```

2.Add router
```golang
//app/router/router.go is the file that manage all URI
func RegisterRouter(router *gin.Engine) {
	entry := router.Group("/demo", middleware.PerfMiddleware(), middleware.XesLoggerMiddleware())
	entry.GET("/test", democontroller.GaeaDemo)
}
```

4.Controller (mvc programming style)
```golang
//app/router/
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
```

5.Try it!
```
curl -X POST http://127.0.0.1/demo/test -d "param=1"
```

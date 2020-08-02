# Gaea
## Introduction
Gaea是基于Gin的Web框架,由于在实际工作中，项目应用到生产环境之前，还有一些列额外的工程化问题需要解决，
否则系统的可移植性较差、开发人员无法把精力专注于业务开发，
融入了一整套解决方案：依赖管理、配置管理、 编译部署、监控&报警，支持一键快速搭建Web应用，如果您正在考虑用Golang编写web server，那么Gaea 无疑是您最优的选择！

## Directory
```
gaea/
├── app                             #项目工程目录
│   ├── router                      #router目录
│   │   └── router.go               #router文件
│   ├── service                     #service目录，业务逻辑
│   │   └── demoservice             #demoservice目录
│   │       └── demoservice.go      #demoservice文件，业务实现逻辑
│   ├── controller                  #contronller目录，控制层
│   │   └── democontroller          #democontroller目录
│   │       └── DemoController.go   #democontroller实现
├── bin                             #可执行文件目录
│   └── gaea                        #可执行文件
├── conf                            #配置文件目录
│   ├── conf.ini                    #配置文件
│   └── log.xml                     #日志配置文件
├── delve.sh    
├── main                            #main目录
│   └── main.go                     #进程启动,http服务配置代码
├── version                         #版本目录
│   └── version.go                  #版本文件
├── Makefile                        #编译文件 make
├── mock                            #mock数据目录
├── run-tests.sh                    #单元测试脚本
├── utils                           #工具目录
│   ├── common.go                   #公共工具文件
│   ├── mock.go                     #mock数据处理文件
│   └── struct.go                   #公共schema文件
└── version                         #服务git版本执行文件
```
## Quick Start
编译和运行
```golang
cd $GOPATH/src/gaea   //进入到工作目录
make //会使用makefile编译并生成二进制文件到bin目录
make dev //会使用makefile编译并生成二进制文件到bin目录，并copy测试配置文件
```

## Example
1.服务端口配置
```golang
//conf/conf.ini
[HTTP]
port = 9898  //服务端口配置
```

2.绑定路由
```golang
//app/router/router.go文件里统一管理路由
func RegisterRouter(router *gin.Engine) {
	entry := router.Group("/demo", middleware.PerfMiddleware(), middleware.XesLoggerMiddleware())
	entry.GET("/test", democontroller.GaeaDemo)
}
```

4.路由方法实现
```golang
//app/router/ xxxxcontroller文件夹里统一编写路由实现方法
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

1.请求示例
```
curl -X POST http://127.0.0.1/demo/test -d "param=1"
```

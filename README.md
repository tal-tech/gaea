## 简介
<p align="center">
 <a href="https://tal-tech.github.io/gaea-doc/" target="_blank">
     <img src="https://i.ibb.co/PN1rS28/11.png?raw=true"  alt="Gaea Logo" align=center />
 </a> 
</p>

Gaea是一基于`Gin`的Web框架。 在实际工作中，在将项目应用于生产环境之前，还需要解决一系列其他工程问题,
否则，系统的可移植性将很差，并且开发人员将无法专注于业务开发。
集成到一套完整的解决方案中：依赖关系管理，配置管理，编译和部署，监视和警报，并支持一键式快速构建Web应用程序。 如果您正在考虑用Golang编写Web服务器，那么Gaea无疑是您的最佳选择！
## Document
[Documentation](https://tal-tech.github.io/gaea-doc)

[中文文档](https://www.yuque.com/tal-tech/gaea)


## 安装

### 安装脚手架
通过 [rigger](https://github.com/tal-tech/rigger)脚手架可一键创建Gaea模板的api项目

### 生成框架

此处以"myproject"为项目名称
```shell
$ rigger new api myproject
正克隆到 '/winshare/go/src/myproject'...
myprojec项目已创建完成, 使用:
 cd /winshare/go/src/myproject && rigger build 

```


### 编译

```golang
//Will use makefile to compile and generate binary files to the bin directory
$ cd gaea
$ make
```

### help
```go
$ ./bin/myproject --help
Usage of ./bin/myproject:
  -c string
    	指定ini配置文件，以程序的二进制文件(myproject)为相对目录，正确的相对目录加载方式： -c ../conf/conf_xxx.ini； 默认为加载 ../conf/conf.ini
  -p string
        配置文件的前缀设置，用于以绝对路径形式加载，如  -c conf.ini -p /usr/pathto/myproject/conf
  -cfg string
    	json config path (default "conf/config.json")		
  -extended string
    	扩展参数，程序未定义使用用途，用户可自行处理
  -f	foreground
  -m	mock 开关	
  -s string
    	start or stop
  -usr1 string
    	user defined flag -usr1
  -usr2 string
    	user defined flag -usr2
  -usr3 string
    	user defined flag -usr3
  -usr4 string
    	user defined flag -usr4
  -usr5 string
    	user defined flag -usr5
  -v	version
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

### 运行
```go
$ rigger start
2019/11/06 21:50:03 CONF INIT,path:../conf/conf.ini
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:    export GIN_MODE=release
 - using code:    gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /demo/test                --> myproject/app/controller/democontroller.MyXesGoDemo (3 handlers)
2019/11/06 21:50:03 [overseer master] run
2019/11/06 21:50:04 [overseer master] starting /winshare/go/src/myproject/bin/myproject
2019/11/06 21:50:04 CONF INIT,path:../conf/conf.ini
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:    export GIN_MODE=release
 - using code:    gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /demo/test                --> myproject/app/controller/democontroller.MyXesGoDemo (3 handlers)
2019/11/06 21:50:04 [overseer slave#1] run
2019/11/06 21:50:04 [overseer slave#1] start program
```

### 测试
```go
$ curl  http://127.0.0.1:9898/demo/test
{"code":0,"data":{"ret1":"Welcome to use myproject!"},"msg":"ok","stat":1}
```

----------
至此，我们已经通过Gaea搭建了一个web服务！
接下来我们进一步学习Gaea 相关配置、特性、组件应用等主题


## 各大主流框架路由性能对比

<p align="center">
 <a href="https://tal-tech.github.io/gaea-doc/" target="_blank">
     <img src="https://i.ibb.co/TB64zTd/jianjie-xingneng.png"  alt="Gaea Logo" align=center />
 </a> 
</p>


(图片来自于网络)


### Gaea性能压测
Gaea框架相比于原生Gin 影响性能的点其实是全部集中在中间件上，因为每次http请求都会跑一边，所以，观测下在各个中间开启时对整体性能的影响情况
### 压测条件

| 条件 |  值 |
| ---- | ---- |
| 系统 |  virtualbox 虚拟机上 centos7 |
| 内存| 1GB |
|CPU| 单核|
|请求数量| 10万|
|并发数量|100|
|传输数据|{"code":0,"data":"hell world","msg":"ok","stat":1}|
 
### 压测结果

<p align="center">
 <a href="https://tal-tech.github.io/gaea-doc/" target="_blank">
     <img src="https://i.ibb.co/Wyjr9Zs/perf.png"  alt="Gaea Logo" align=center />
 </a> 
</p>

从图中我们可以明显看出：
* Gaea的默认配置会带来一定的性能耗损，大约30%
* 其中`Logger`中间件在各个中间件影响性能比重最大，其它中间件几乎可以忽略不计

在实际项目应用中，当`Logger` 中间件是瓶颈点时，我们可以关闭它，毕竟请求日志在网关层也会记录！

此外，对[日志库](https://github.com/tal-tech/loggerX)使用建议日志级别设定为`WARNING`

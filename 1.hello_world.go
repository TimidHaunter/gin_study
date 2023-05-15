package main

import "github.com/gin-gonic/gin"

func main() {
	//创建一个默认的路由
	router := gin.Default()

	//绑定路由规则和路由函数，访问/index地址的路由，由对应的函数去处理
	//回调函数
	// router.GET("/index", func(context *gin.Context) {
	// 	context.String(200, "Hello World!")
	// })

	//普通函数
	router.GET("/index", Index)

	//启动监听，gin会把web服务运行在本机0.0.0.0的8081端口
	router.Run("0.0.0.0:8081")
}

func Index(context *gin.Context) {
	//响应字符串
	context.String(200, "Gin 正在运行!")
}

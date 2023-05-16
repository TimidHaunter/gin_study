//go:build ignore

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//加载html模板
	// router.LoadHTMLGlob("templates/*")

	//静态资源
	router.StaticFile("golang.png", "static/images/golang.png")

	//响应字符串
	router.GET("/", func(context *gin.Context) {
		//状态码 HTTP status
		context.String(http.StatusOK, "你好")
	})

	router.GET("/notFound", func(context *gin.Context) {
		context.String(404, "notFound")
	})

	router.GET("/internalServerError", func(context *gin.Context) {
		context.String(500, "Internal Server Error")
	})

	router.GET("/string", _string)

	//响应json原始
	router.GET("/json", _json)
	//响应json结构体
	router.GET("/jsonStruct", _jsonStruct)
	//响应json-map
	router.GET("/jsonMap", _jsonMap)
	//响应xml
	router.GET("/xml", _xml)
	//响应yaml
	router.GET("/yaml", _yaml)
	//返回html
	// router.GET("/html", _html)
	//重定向
	router.GET("/baidu", _redirect)

	router.Run(":8081")
}

func _string(context *gin.Context) {
	context.String(http.StatusOK, "this is string")
}

func _json(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "this is a json", "status": http.StatusOK})
}

func _jsonStruct(context *gin.Context) {
	type UserInfo struct {
		//json 修改字段名称
		UserName string `json:"user_name"`
		Age      int    `json:"age"`
		//password 字段需要过滤掉
		Password string `json:"-"`
	}
	user := UserInfo{"Yintian", 31, "123456"}
	context.JSON(http.StatusOK, user)
}

func _jsonMap(context *gin.Context) {
	userMap := map[string]string{
		"user_name": "Xuxu",
		"age":       "28",
	}
	context.JSON(200, userMap)
}

func _xml(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{"message": "this is a xml", "status": http.StatusOK})
}

func _yaml(context *gin.Context) {
	context.YAML(http.StatusOK, gin.H{"message": "this is a yaml", "status": http.StatusOK})
}

// gin.H{} 可以传参进模板 {"username":"yint"}
// 也可以传递结构体
// func _html(context *gin.Context) {
// 	context.HTML(http.StatusOK, "index.html", gin.H{})
// }

// 301 302区别
// func _redirect(context *gin.Context) {
// 	context.Redirect(301, "http://www.baidu.com")
// }

func _redirect(context *gin.Context) {
	//重定向自己的路由
	context.Redirect(http.StatusMovedPermanently, "/string")
}

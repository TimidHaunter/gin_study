//go:build ignore

package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func _query(content *gin.Context) {
	//接收query参数
	id := content.Query("id")
	user := content.Query("user")

	fmt.Println(id)
	fmt.Println(user)

	//address key没有 false
	//key有 value没有 是address=空 true
	//key有 value有 陕西省汉中市汉台区兴汉路 true
	fmt.Println(content.GetQuery("address"))

	//user=yintx&user=zxu [yint zxu]
	fmt.Println(content.QueryArray("user"))

	fmt.Println(content.DefaultQuery("age", "16"))
}

func _param(content *gin.Context) {
	fmt.Println(content.Param("user_id"))
	fmt.Println(content.Param("book_id"))
}

func _postform(content *gin.Context) {
	fmt.Println(content.PostForm("name"))
	//取不到key，给默认值
	fmt.Println(content.DefaultPostForm("age", "18"))

	//接收所有参数 &{map[name:[yintx]] map[file:[0xc000036840]]} <nil> 图片也可以
	forms, err := content.MultipartForm()
	fmt.Println(forms, err)
}

/**
 * 原始参数
 * 可以接收 post form-data x-www-form-unlencoded json
 */
// func _raw(content *gin.Context) {
// 	// fmt.Println(content.GetRawData())

// 	body, _ := content.GetRawData()
// 	contentType := content.GetHeader("Content-Type")

// 	switch contentType {
// 	case "application/json":
// 		//解析到结构体
// 		type User struct {
// 			Name string `json:"name"`
// 			Age  int    `json:"age"`
// 		}
// 		var user User
// 		//解析JSON
// 		err := json.Unmarshal(body, &user)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		fmt.Println(user)
// 	}

// 	// fmt.Println(string(body))

// 	//header zzz
// 	// header := content.GetHeader("YYY")
// 	// fmt.Println(header)
// }

func _raw(content *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	//指针
	err := bindJson(content, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

/*
 * obj any 泛型
 */
func bindJson(content *gin.Context, obj any) (err error) {
	body, _ := content.GetRawData()
	contentType := content.GetHeader("Content-Type")

	switch contentType {
	case "application/json":
		//解析JSON
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

func main() {
	router := gin.Default()

	//查询参数 Query
	router.GET("/query", _query)

	//动态参数 Param
	// router.GET("/param/:user_id", _param)
	router.GET("/param/:user_id/:book_id", _param)

	//表单 PostForm 文件上传
	router.POST("/form", _postform)

	//json GetRowData
	router.POST("/raw", _raw)

	router.Run(":8081")
}

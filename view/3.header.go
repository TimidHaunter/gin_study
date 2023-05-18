//go:build ignore

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/header", func(c *gin.Context) {
		//不区分首字母大小写
		fmt.Println(c.GetHeader("User-Agent"))

		fmt.Println(c.Request.Header)

		fmt.Println(c.Request.Header.Get("User-Agent"))

		c.JSON(http.StatusOK, gin.H{"msg": "sucessful"})
	})

	//响应头
	router.GET("/header/response", func(c *gin.Context) {
		c.Header("Response-Token", "yintx")
		c.JSON(http.StatusOK, gin.H{"msg": "response"})
	})

	router.Run(":8081")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//r := gin.Default()
	engine := gin.Default()



	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	/**
		demo1
	 */
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		//获取请求接口
		fmt.Println(context.FullPath())
		//获取字符串参数
		name := context.DefaultQuery("name", "")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("Hello ," + name))
	})

	/**
	demo2
	 */
	engine.Handle("GET", "/helloString", func(context *gin.Context) {
		//获取请求接口
		fmt.Println(context.FullPath())
		//获取字符串参数
		name := context.DefaultQuery("name", "")
		fmt.Println("---",name)

		//输出
		context.Writer.WriteString("HelloString ," + name)
	})


	/**
	static
	 */
	engine.LoadHTMLGlob("./html/*")
	engine.Static("/img","./img")



	/**
	demo3
	 */
	engine.GET("/helloHtml", func(context *gin.Context) {
		fullPath := "请求路径:" + context.FullPath()
		fmt.Println(fullPath)

		context.HTML(http.StatusOK,"index.html",gin.H{
			"fullPath": fullPath,
			"title": "gin测试",
		})
	})



	engine.Run(":8081") // 监听并在 0.0.0.0:8080 上启动服务

}

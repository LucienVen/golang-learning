package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	// 初始化引擎
	engine := gin.Default()

	// 注册一个路由器和处理函数
	engine.Any("/", WebRoot)

	// 绑定端口，然后启动应用
	engine.Run(":8083")
	
}

func WebRoot(content *gin.Context)  {
	content.String(http.StatusOK, "Hello, World!")
}
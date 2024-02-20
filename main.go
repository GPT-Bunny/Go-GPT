package main

import (
	"fmt"
	auth "go-gpt/api"   // 导入自定义的auth包，处理认证相关逻辑
	openai "go-gpt/api" // 导入自定义的openai包，处理与OpenAI相关的逻辑

	"github.com/gin-contrib/cors" // 导入CORS中间件
	"github.com/gin-gonic/gin"    // 导入Gin框架
)

func main() {
	// 初始化数据库连接
	err := auth.InitDB("go-gpt:go-gpt@tcp(127.0.0.1:3306)/go-gpt")
	if err != nil {
		panic(err) // 如果初始化失败，则抛出异常
	}
	defer auth.CloseDB() // 确保在main函数结束时关闭数据库连接

	r := gin.Default() // 创建一个Gin引擎实例

	// 使用 CORS 中间件允许跨域请求
	r.Use(cors.Default())

	// 注册登录和注册的路由处理函数
	r.POST("/login", auth.HandleLogin)
	r.POST("/register", auth.HandleRegister)

	// 注册聊天和设置模型的路由处理函数
	r.POST("/chat", func(c *gin.Context) {
		openai.ChatHandler(c.Writer, c.Request) // 聊天处理函数
	})
	r.POST("/set-model", func(c *gin.Context) {
		openai.SetModelHandler(c.Writer, c.Request) // 设置模型处理函数
	})

	// 启动服务器，并在8080端口监听
	fmt.Println("Server is running on http://localhost:8080")
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}

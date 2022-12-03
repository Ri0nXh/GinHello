package router

import (
	"GinHello/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/api/v1")
	{
		v1.GET("/index", handler.Index)
		v1.POST("/register", handler.Register)
		v1.POST("/login", handler.Login)
	}
	return router
}

package router

import "github.com/gin-gonic/gin"

func New() *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	r.GET("/ping")
	return r
}

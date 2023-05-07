package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vinoMamba/naruto-go/docs"
	"github.com/vinoMamba/naruto-go/internal/controller"
)

func New() *gin.Engine {
	r := gin.New()
	r.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	r.GET("/api/v1/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

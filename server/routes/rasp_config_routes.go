package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/middleware"
)

func InitRaspConfigRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	raspConfigController := controller.NewRaspConfigController()
	router := r.Group("/config")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.POST("/create", raspConfigController.CreateRaspConfig)
		router.GET("/list", raspConfigController.GetRaspConfigs)
		router.POST("/delete/batch", raspConfigController.BatchDeleteConfigByIds)
	}
	return r
}

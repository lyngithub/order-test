package initialize

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mxshop_api/middlewares"
	"mxshop_api/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "success",
		})
	})

	Router.Use(middlewares.Cors()) //跨域
	zap.S().Info("配置用户相关的url")
	ApiGroup := Router.Group("/o/v1")
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)
	return Router
}

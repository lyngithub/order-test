package router

import (
	"github.com/gin-gonic/gin"

	"mxshop_api/api/order"
	//"mxshop_api/order_web/api/pay"
	"mxshop_api/middlewares"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth())
	{
		OrderRouter.GET("", order.List)       // 订单列表
		OrderRouter.POST("", order.New)       // 新建订单
		OrderRouter.GET("/:id", order.Detail) // 订单详情
	}
	//PayRouter := Router.Group("pay")
	//{
	//	PayRouter.POST("alipay/notify", pay.Notify)
	//}
}

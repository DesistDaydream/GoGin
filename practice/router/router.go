package router

import (
	"github.com/DesistDaydream/GoGin/practice/handler"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由，设定路由信息
func InitRouter(r *gin.Engine) {
	r.GET("/", handler.IndexGet)
	r.GET("/login", handler.LoginGet)
	r.POST("/login", handler.LoginPost)
	r.GET("/order", handler.OrderGet)
	r.POST("/order", handler.OrderPost)
	r.GET("/stock-in", handler.StockInGet)
	r.POST("/stock-in", handler.StockInPost)
	r.GET("/stock-out", handler.StockOutGet)
	r.POST("/stock-out", handler.StockOutPost)
	r.GET("/query", handler.QueryGet)
	r.POST("/query", handler.QueryPost)
	r.GET("/inventory", handler.CommodityGet)
}

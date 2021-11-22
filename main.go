package main

import (
	"net/http"
	"os"

	"github.com/bitcou/bitcou-wrapper/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	App := GetApp()
	_ = App.Run(":" + port)
}

func GetApp() *gin.Engine {
	App := gin.Default()
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	corsConf.AllowHeaders = []string{"token", "service", "content-type"}
	App.Use(cors.New(corsConf))
	ApplyRoutes(App)
	return App
}

func ApplyRoutes(r *gin.Engine) {
	bc := controllers.NewBitcouController()
	apiBitcou := r.Group("/bitcou/")
	{
		apiBitcou.POST("order", bc.CreateOrder)
		apiBitcou.GET("vouchers", bc.GetVouchers)
		apiBitcou.GET("vouchers/compact", bc.GetCompactVouchers)
		apiBitcou.GET("account", bc.GetAccountInfo)
		apiBitcou.GET("account/balance", bc.GetAccountBalance)
		apiBitcou.GET("vouchers/:voucherId", bc.GetVoucher)
		apiBitcou.GET("order/:orderId", bc.GetOrder)
	}

	bl := controllers.NewBlockchainController()
	apiBlockchain := r.Group("/blockchain/")
	{
		apiBlockchain.POST("encrypt", bl.Encrypt)
		apiBlockchain.POST("decrypt", bl.Decrypt)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

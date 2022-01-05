package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/bitcou/bitcou-wrapper/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env ", err.Error())
		return
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
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
		apiBitcou.GET("vouchers/catalog", bc.GetCatalog)
		apiBitcou.GET("account", bc.GetAccountInfo)
		apiBitcou.GET("account/balance", bc.GetAccountBalance)
		apiBitcou.GET("vouchers/:voucherId", bc.GetVoucher)
		apiBitcou.GET("order/:orderId", bc.GetOrder)
		apiBitcou.GET("countries", bc.GetCountries)
		apiBitcou.GET("countries/:countryId", bc.GetCountries)
		apiBitcou.GET("categories", bc.GetCategories)
		apiBitcou.GET("categories/:categoryId", bc.GetCategories)
	}

	bl := controllers.NewBlockchainController()
	apiBlockchain := r.Group("/blockchain/")
	{
		apiBlockchain.POST("encrypt", bl.Encrypt)
		// apiBlockchain.POST("decrypt", bl.Decrypt)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

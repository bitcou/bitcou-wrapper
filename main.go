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
	firebase := controllers.NewFirebaseHandler()
	//err := firebase.RegisterPurchase("0xEa703E63BA6C9b5224969d6483327B8e65AF76CC", models.FirebaseAccount{
	//	Address: "0xEa703E63BA6C9b5224969d6483327B8e65AF76CC",
	//	Purchases: []models.FirebasePurchase{models.FirebasePurchase{
	//		ID:            "1",
	//		ProductId:     "9856",
	//		TotalValue:    6,
	//		TransactionId: "0x90132cfc9c656f59f21c3341718ad635b2a67ba30a993d1cd7ed61bbd2fcee7f",
	//		Status:        "REDEEMED",
	//		Timestamp:     1641841624,
	//	}},
	//})
	data, err := firebase.GetPurchasesByAddress("0xEa703E63BA6C9b5224969d6483327B8e65AF76CC")
	fmt.Println(data)
	if err != nil {
		fmt.Print("error purchasing: ", err)
	}
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
		apiBlockchain.POST("decrypt", bl.Decrypt)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}

package main

import (
	"github.com/bitcou/bitcou-wrapper/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	api := r.Group("/")
	{
		api.GET("balance", bc.GetAccountBalance)
	}

	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})
}
package controllers

import (
	"fmt"
	"github.com/bitcou/bitcou-wrapper/firebase"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type FirebaseController struct {
	handler *firebase.FireStoreHandler
}

func NewFirebaseController() *FirebaseController {
	fs := new(FirebaseController)
	fs.handler = firebase.NewFirebaseHandler()
	return fs
}

func (fs *FirebaseController) GetAccountPurchases(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}
	fmt.Println(value)
	// TODO Add logic to verify signature and remove hard coded signature
	purchases, err := fs.handler.GetPurchasesByAddress("0xEa703E63BA6C9b5224969d6483327B8e65AF76CC")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}

	c.IndentedJSON(http.StatusOK, purchases)
}

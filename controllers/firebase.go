package controllers

import (
	"encoding/json"
	"github.com/bitcou/bitcou-wrapper/bitcou"
	"github.com/bitcou/bitcou-wrapper/firebase"
	"github.com/bitcou/bitcou-wrapper/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
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
	var message bitcou.MessageVerification
	err = json.Unmarshal(value, &message)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	nonce, err := fs.handler.GetSecretByAddress(message.Address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	verified := utils.VerifySig(message.Address, message.Message, []byte(strconv.Itoa(nonce.Nonce)))
	if verified {
		purchases, err := fs.handler.GetPurchasesByAddress(message.Address)
		_, err = fs.handler.UpdateNonce(message.Address)
		if err != nil {
			return
		}
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, nil)
		}
		c.IndentedJSON(http.StatusOK, purchases)
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, nil)
	}

}

func (fs *FirebaseController) GetAccountSecret(c *gin.Context) {
	address := c.Param("address")
	secret, err := fs.handler.GetSecretByAddress(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, secret)
}

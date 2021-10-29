package controllers

import (
	"github.com/bitcou/bitcou-wrapper/bitcou"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BitcouController struct {
	client *bitcou.Bitcou
}

func NewBitcouController() *BitcouController {
	bc := new(BitcouController)
	bc.client = bitcou.NewBitcou("", true)
	return bc
}

func (b *BitcouController) GetAccountBalance(c *gin.Context) {
	accountBalance, err := b.client.GetBalance()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, accountBalance)
}

package controllers

import (
	"net/http"

	"github.com/bitcou/bitcou-wrapper/bitcou"
	"github.com/gin-gonic/gin"
)

type BitcouController struct {
	client *bitcou.Bitcou
}

func NewBitcouController() *BitcouController {
	bc := new(BitcouController)
	bc.client = bitcou.NewBitcou("", true)
	return bc
}

func (b *BitcouController) CreateOrder(c *gin.Context) {
	// orderInfo, err := b.client.CreateOrder()
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, nil)
	// 	return
	// }
	// c.IndentedJSON(http.StatusOK, orderInfo)
}

func (b *BitcouController) GetVouchers(c *gin.Context) {
	vouchers, err := b.client.Products(bitcou.FULL_PRODUCTS)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, vouchers)
}

func (b *BitcouController) GetCompactVouchers(c *gin.Context) {
	vouchers, err := b.client.Products(bitcou.COMPACT_PRODUCTS)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, vouchers)
}

func (b *BitcouController) GetAccountInfo(c *gin.Context) {
	accountInfo, err := b.client.AccountInfo(bitcou.ACCOUNT_INFO)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, accountInfo)
}

func (b *BitcouController) GetAccountBalance(c *gin.Context) {
	accountBalance, err := b.client.AccountInfo(bitcou.ACCOUNT_BALANCE)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, accountBalance)
}

func (b *BitcouController) GetVoucher(c *gin.Context) {
	voucherId := c.Param("voucherId")
	voucher, err := b.client.Products(bitcou.SINGULAR_PRODUCT, voucherId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, voucher)
}

func (b *BitcouController) GetOrder(c *gin.Context) {
	// orderId := c.Param("orderId")
	// order, err := b.client.GetOrder(orderId)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, nil)
	// 	return
	// }
	// c.IndentedJSON(http.StatusOK, order)
}

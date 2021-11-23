package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitcou/bitcou-wrapper/bitcou"
	"github.com/gin-gonic/gin"
)

type BitcouController struct {
	client *bitcou.Bitcou
}

func NewBitcouController() *BitcouController {
	bc := new(BitcouController)
	bc.client = bitcou.NewBitcou(os.Getenv("BITCOU_APIKEY"), true)
	return bc
}

func (b *BitcouController) CreateOrder(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	var purchaseInput bitcou.PurchaseInput
	err = json.Unmarshal(value, &purchaseInput)
	orderInfo, err := b.client.Purchases(bitcou.CREATE_ORDER, purchaseInput, "")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, orderInfo)
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

func (b *BitcouController) GetCatalog(c *gin.Context) {
	vouchers, err := b.client.Catalog()
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
	voucher, err := b.client.Products(bitcou.SINGLE_PRODUCT, voucherId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, voucher)
}

func (b *BitcouController) GetOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	order, err := b.client.Purchases(bitcou.GET_ORDER, *new(bitcou.PurchaseInput), orderId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, order)
}

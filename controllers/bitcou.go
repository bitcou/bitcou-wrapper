package controllers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

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
	orderInfo, err := b.client.Purchases(bitcou.CREATE_ORDER, value, "")
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
	variantProductID := c.Query("variantId")
	country := c.Query("country")
	category := c.Query("category")
	categoryNumeric := 0
	var err error

	if category != "" {
		categoryNumeric, err = strconv.Atoi(category)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, errors.New("invalid product id"))
			return
		}
	}

	vouchers, err := b.client.Catalog(variantProductID, country, categoryNumeric)
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
	order, err := b.client.Purchases(bitcou.GET_ORDER, []byte(""), orderId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, order)
}

func (b *BitcouController) GetCountries(c *gin.Context) {
	countryId := c.Param("countryId")
	countries, err := b.client.Countries(countryId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, countries)
}

func (b *BitcouController) GetCategories(c *gin.Context) {
	categoryId := c.Param("categoryId")
	categories, err := b.client.Categories(categoryId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, categories)
}

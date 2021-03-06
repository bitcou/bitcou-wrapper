package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bitcou/bitcou-wrapper/bitcou"
	wrap_err "github.com/bitcou/bitcou-wrapper/errors"
	"github.com/bitcou/bitcou-wrapper/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type BlockchainController struct {
	key []byte
}

func NewBlockchainController() *BlockchainController {
	bl := new(BlockchainController)
	godotenv.Load()
	bl.key = []byte(os.Getenv("BLOCKCHAIN_KEY"))
	return bl
}

type CipherString struct {
	CipherStr []byte `json:"cipherStr"`
}

func (b *BlockchainController) VerifyMessage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("blockchain::encrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	var message bitcou.MessageVerification
	err = json.Unmarshal(value, &message)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}

	isVerified := utils.VerifySig(message.Address, message.Message, []byte("hello"))
	if isVerified {
		fmt.Println("verified")
		c.IndentedJSON(http.StatusOK, "")
	} else {
		fmt.Println("not verified")
	}
}

func (b *BlockchainController) Encrypt(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("blockchain::encrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	var purchaseInput bitcou.PurchaseInput
	err = json.Unmarshal(value, &purchaseInput)
	if err != nil {
		log.Println("blockchain::encrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInvalidJson))
		return
	}
	inputJson, err := json.Marshal(purchaseInput)
	if err != nil {
		log.Println("blockchain::encrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	encryptedData, err := b.encryptData(inputJson)
	if err != nil {
		log.Println("blockchain::encrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	encryptedDataObj := CipherString{
		CipherStr: encryptedData,
	}
	c.IndentedJSON(http.StatusOK, encryptedDataObj)
}

func (b *BlockchainController) Decrypt(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("blockchain::decrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	var message CipherString
	err = json.Unmarshal(value, &message)
	if err != nil {
		log.Println("blockchain::decrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(err))
		return
	}
	plainText, err := b.decryptData(message.CipherStr)
	if err != nil {
		log.Println("blockchain::decrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(plainText, &jsonMap)
	if err != nil {
		log.Println("blockchain::decrypt::error ", err)
		c.IndentedJSON(http.StatusInternalServerError, wrap_err.New(wrap_err.ErrorInternalServer))
		return
	}
	c.IndentedJSON(http.StatusOK, jsonMap)
}

func (b *BlockchainController) encryptData(data []byte) ([]byte, error) {
	gcm, err := b.createGCM()
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func (b *BlockchainController) decryptData(data []byte) ([]byte, error) {
	gcm, err := b.createGCM()
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func (b *BlockchainController) createGCM() (cipher.AEAD, error) {
	block, err := aes.NewCipher(b.key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}

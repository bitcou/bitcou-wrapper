package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bitcou/bitcou-wrapper/bitcou"
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

func (b *BlockchainController) Encrypt(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	var purchaseInput bitcou.PurchaseInput
	err = json.Unmarshal(value, &purchaseInput)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	inputJson, err := json.Marshal(purchaseInput)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	encryptedData, err := b.encryptData(inputJson)
	encryptedDataObj := CipherString{
		CipherStr: encryptedData,
	}
	c.IndentedJSON(http.StatusOK, encryptedDataObj)
}

func (b *BlockchainController) Decrypt(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	var message CipherString
	err = json.Unmarshal(value, &message)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	plainText, err := b.decryptData(message.CipherStr)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(plainText, &jsonMap)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
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

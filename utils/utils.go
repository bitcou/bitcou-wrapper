package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"os"
)

type CipherString struct {
	CipherStr []byte `json:"cipherStr"`
}

func CreateGCM() (cipher.AEAD, error) {
	key := os.Getenv("BLOCKCHAIN_KEY")
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm, nil
}

func DecryptInit(value []byte) ([]byte, error) {
	var message CipherString
	err := json.Unmarshal(value, &message)
	if err != nil {
		return nil, err
	}
	plainText, err := decryptCipher(message.CipherStr)
	if err != nil {
		return nil, err
	}
	var jsonMap map[string]interface{}
	err = json.Unmarshal(plainText, &jsonMap)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}

func decryptCipher(data []byte) ([]byte, error) {
	gcm, err := CreateGCM()
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

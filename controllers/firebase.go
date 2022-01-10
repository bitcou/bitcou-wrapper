package controllers

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/bitcou/bitcou-wrapper/models"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
	"log"
	"os"
	"time"
)

type FireStoreHandler struct {
	firestore *firestore.Client
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env ", err.Error())
		return
	}
}

func NewFirebaseHandler() *FireStoreHandler {
	handler := new(FireStoreHandler)
	encodedCreds := os.Getenv("FIREBASE_CRED")
	fbCred, err := base64.StdEncoding.DecodeString(encodedCreds)
	if err != nil {
		log.Println("error loading firebase credentials")
	}
	opt := option.WithCredentialsJSON(fbCred)
	fbApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("unable to initialize firebase app")
	}

	// Init Database (Firestore)
	fs, err := fbApp.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	handler.firestore = fs
	return handler
}

func (fs *FireStoreHandler) GetPurchasesByAddress(address string) (account models.FirebaseAccount, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doc, err := fs.firestore.Collection("purchases").Doc(address).Get(ctx)
	if err != nil {
		return models.FirebaseAccount{}, err
	}
	err = doc.DataTo(&account)
	if err != nil {
		return models.FirebaseAccount{}, err
	}
	return account, nil
}

func (fs *FireStoreHandler) RegisterPurchase(address string, values models.FirebaseAccount) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doc, err := fs.firestore.Collection("purchases").Doc(address).Set(ctx, values)
	fmt.Println(doc)
	return err
}

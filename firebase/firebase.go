package firebase

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

func (fs *FireStoreHandler) GetSecretByAddress(address string) (nonce models.FirebaseNonce, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doc, err := fs.firestore.Collection("secrets").Doc(address).Get(ctx)
	if !doc.Exists() {
		_, err = fs.firestore.Collection("secrets").Doc(address).Set(ctx, nonce)
		if err != nil {
			return models.FirebaseNonce{Nonce: -1}, err
		}
		return nonce, nil
	} else {
		err = doc.DataTo(&nonce)
		return nonce, err
	}
}

// UpdateNonce
// Updates nonce value with a +1, meant to be called on every message verification
func (fs *FireStoreHandler) UpdateNonce(address string) (nonce models.FirebaseNonce, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doc, err := fs.firestore.Collection("secrets").Doc(address).Get(ctx)
	if doc.Exists() {
		err = doc.DataTo(&nonce)
		nonce.Nonce += 1
		_, err = fs.firestore.Collection("secrets").Doc(address).Set(ctx, nonce)
		if err != nil {
			return models.FirebaseNonce{Nonce: -1}, err
		}
		return nonce, nil
	}
	return models.FirebaseNonce{Nonce: -1}, nil
}

func (fs *FireStoreHandler) RegisterPurchase(address string, newPurchase models.FirebasePurchase) error {
	currentInfo, err := fs.GetPurchasesByAddress(address)
	if err != nil {
		currentInfo = models.FirebaseAccount{}
	}
	currentInfo.Purchases = append(currentInfo.Purchases, newPurchase)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	doc, err := fs.firestore.Collection("purchases").Doc(address).Set(ctx, currentInfo)
	fmt.Println(doc)
	return err
}

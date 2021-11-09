package bitcou

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/hasura/go-graphql-client"
	"github.com/joho/godotenv"
)

type Bitcou struct {
	apiKey string
	URL    string
	dev    bool
	client graphql.Client
}

func NewBitcou(apiKey string, dev bool) *Bitcou {
	b := new(Bitcou)
	b.apiKey = apiKey
	if dev {
		b.URL = "https://sandbox-bitcou.kindynos.com/query"
	} else {
		b.URL = "https://api-bitcou.kindynos.com/query"
	}
	b.client = b.newBitcouClient()
	return b
}

func (b *Bitcou) Products(prodInfo ...string) (interface{}, error) {
	if prodInfo[0] == FULL_PRODUCTS {
		err := b.client.Query(context.Background(), &productsQuery, nil)
		if err != nil {
			log.Println("gql::products::error ", err)
			return nil, err
		}
		return productsQuery.Products, nil
	} else if prodInfo[0] == COMPACT_PRODUCTS {
		err := b.client.Query(context.Background(), &compactProductsQuery, nil)
		if err != nil {
			log.Println("gql::products::error ", err)
			return nil, err
		}
		return compactProductsQuery.Products, nil
	} else if prodInfo[0] == SINGULAR_PRODUCT {
		log.Println("product id to retrieve: ", prodInfo[1])
		variables := map[string]interface{}{
			"prodId": graphql.ID(prodInfo[1]),
		}
		err := b.client.Query(context.Background(), &singularProductQuery, variables)
		if err != nil {
			log.Println("gql::products::error ", err)
			return nil, err
		}
		return singularProductQuery.Products[0], nil
	} else {
		return nil, nil
	}
}

func (b *Bitcou) AccountInfo(info string) (interface{}, error) {
	if info == ACCOUNT_INFO {
		err := b.client.Query(context.Background(), &accountInfoQuery, nil)
		if err != nil {
			log.Println("gql::accountInfo::error ", err)
			return nil, err
		}
		return accountInfoQuery.AccountInfo, nil
	} else if info == ACCOUNT_BALANCE {
		err := b.client.Query(context.Background(), &accountBalanceQuery, nil)
		if err != nil {
			log.Println("gql::accountInfo::error ", err)
			return nil, err
		}
		return accountBalanceQuery.AccountInfo, nil
	} else {
		return nil, nil
	}
}

func (b *Bitcou) Purchases(info string, purchaseInfo PurchaseInput, id string) (interface{}, error) {
	if info == CREATE_ORDER {
		variables := map[string]interface{}{
			"purchaseInput": purchaseInfo,
		}
		err := b.client.Mutate(context.Background(), &createPurchaseQuery, variables)
		if err != nil {
			log.Println("gql::purchases::error ", err)
			return nil, err
		}
		return createPurchaseQuery.CreatePurchase, nil
	} else if info == GET_ORDER {
		variables := map[string]interface{}{
			"orderId": graphql.ID(id),
		}
		err := b.client.Query(context.Background(), &getPurchaseQuery, variables)
		if err != nil {
			log.Println("gql::purchases::error ", err)
			return nil, err
		}
		return getPurchaseQuery.Purchases[0], nil
	} else {
		return nil, nil
	}
}

type MyRoundTripper struct {
	r http.RoundTripper
}

func (mrt MyRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	godotenv.Load()
	r.Header.Add("X-API-Key", os.Getenv("BITCOU_APIKEY"))
	return mrt.r.RoundTrip(r)
}

func (b *Bitcou) newBitcouClient() graphql.Client {
	httpClient := &http.Client{
		Transport: MyRoundTripper{r: http.DefaultTransport},
	}
	return *graphql.NewClient(b.URL, httpClient)
}

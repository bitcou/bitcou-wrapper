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

func (b *Bitcou) Products(getFullProducts bool) (interface{}, error) {
	if getFullProducts {
		err := b.client.Query(context.Background(), &productsQuery, nil)
		if err != nil {
			log.Println("gql::products::error ", err)
			return nil, err
		}
		return productsQuery.Products, nil
	} else {
		err := b.client.Query(context.Background(), &compactProductsQuery, nil)
		if err != nil {
			log.Println("gql::products::error ", err)
			return nil, err
		}
		return compactProductsQuery.Products, nil
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

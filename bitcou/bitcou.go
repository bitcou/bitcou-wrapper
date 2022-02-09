package bitcou

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/bitcou/bitcou-wrapper/models"

	wrap_err "github.com/bitcou/bitcou-wrapper/errors"
	"github.com/bitcou/bitcou-wrapper/utils"
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
	b.dev = dev
	if dev {
		b.URL = "https://sandbox-bitcou.kindynos.mx/query"
	} else {
		b.URL = "https://api-bitcou.kindynos.mx/query"
	}
	b.client = b.newBitcouClient()
	return b
}

func (b *Bitcou) Catalog(variantProductID string, country string, category int) (interface{}, error) {
	variables := make(map[string]interface{})
	filter := CatalogFilter{
		VariantProductID: graphql.String(""),
		Country:          graphql.String(""),
		Category:         graphql.ID(nil),
	}
	//mapFilters := make(map[string]interface{})
	if variantProductID != "" {
		filter.VariantProductID = graphql.String(variantProductID)
	}
	if country != "" {
		filter.Country = graphql.String(country)
	}
	if category != 0 {
		filter.Category = graphql.ID(category)
	}

	variables["filter"] = filter

	err := b.client.Query(context.Background(), &catalogQuery, variables)
	if err != nil {
		log.Println("gql::catalog::error ", err)
		return nil, err
	}
	return catalogQuery.Brand, nil
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
	} else if prodInfo[0] == SINGLE_PRODUCT {
		variables := map[string]interface{}{
			"prodId": graphql.ID(prodInfo[1]),
		}
		err := b.client.Query(context.Background(), &singleProductQuery, variables)
		if err != nil {
			log.Println("gql::products::error ", err)
			return nil, err
		}
		if len(singleProductQuery.Products) > 0 {
			return singleProductQuery.Products[0], nil
		} else {
			return nil, wrap_err.ErrorProductNotFound
		}
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

func (b *Bitcou) Purchases(option OrderOperations, purchaseInfo []byte, id string) (interface{}, error) {
	if option == CREATE_ORDER {
		var data models.CreateOrderEncryptedInput
		err := json.Unmarshal(purchaseInfo, &data)
		if err != nil {
			return nil, wrap_err.ErrorInternalServer
		}
		plainText, err := utils.DecryptInit(data.UserInfo)
		if err != nil {
			log.Println("gql::purchases::error ", err)
			return nil, wrap_err.ErrorInternalServer
		}

		var input PurchaseInput
		err = json.Unmarshal(plainText, &input)
		input.TransactionID = data.TransactionId
		if err != nil {
			log.Println("gql::purchases::error ", err)
			return nil, wrap_err.ErrorInternalServer
		}

		variables := map[string]interface{}{
			"purchaseInput": input,
		}
		err = b.client.Mutate(context.Background(), &createPurchaseQuery, variables)
		if err != nil {
			log.Println("gql::purchases::error ", err)
			return nil, err
		}
		return createPurchaseQuery.CreatePurchase, nil
	} else if option == GET_ORDER {
		variables := map[string]interface{}{
			"orderId": graphql.ID(id),
		}
		err := b.client.Query(context.Background(), &getPurchaseQuery, variables)
		if err != nil {
			log.Println("gql::purchases::error ", err)
			return nil, err
		}

		if len(getPurchaseQuery.Purchases) > 0 {
			return getPurchaseQuery.Purchases[0], nil
		} else {
			return nil, wrap_err.ErrorOrderNotFound
		}
	} else {
		return nil, nil
	}
}

func (b *Bitcou) Countries(id string) (interface{}, error) {
	variables := map[string]interface{}{
		"countryId": graphql.String(id),
	}
	err := b.client.Query(context.Background(), &getCountries, variables)
	if err != nil {
		log.Println("gql::countries::error ", err)
		return nil, err
	}
	if id != "" {
		if len(getCountries.Countries) > 0 {
			return getCountries.Countries[0], nil
		} else {
			return nil, wrap_err.ErrorInvalidCountry
		}
	} else {
		return getCountries.Countries, nil
	}
}

func (b *Bitcou) Categories(id string) (interface{}, error) {
	var err error
	if id == "" {
		err = b.client.Query(context.Background(), &getCategories, nil)
		if err != nil {
			log.Println("gql::categories::error ", err)
			return nil, err
		}
		return getCategories.Categories, nil
	} else {
		variables := map[string]interface{}{
			"categoryId": graphql.ID(id),
		}
		err = b.client.Query(context.Background(), &getCategoriesFilter, variables)
		if err != nil {
			log.Println("gql::categories::error ", err)
			return nil, err
		}
		return getCategoriesFilter.Categories[0], nil
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

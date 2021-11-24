package bitcou

import "github.com/hasura/go-graphql-client"

// Queries for products
const (
	FULL_PRODUCTS    string = "FULL"
	COMPACT_PRODUCTS        = "COMPACT"
	SINGLE_PRODUCT          = "SINGLE_PRODUCT"
)

var catalogQuery struct {
	Brand []struct {
		VariantID graphql.String `graphql:"variantID"`
		BrandImageURL graphql.String `graphql:"brandImageURL"`
		BrandName graphql.String
		Products []struct {
			ID                       graphql.String
			Available                graphql.Boolean
			Description              graphql.String
			CustomDescription        graphql.String
			DiscountAbsolute         graphql.Float
			DiscountPercentage       graphql.Float
			CustomDiscount           graphql.Float
			Currency                 graphql.String
			InternalCurrency         graphql.String
			IsVariablePrice          graphql.Boolean
			FixedPrice               graphql.Float
			FaceValue                graphql.Float
			EurFixedPrice            graphql.Float
			InternalFixedPrice       graphql.Float
			MinVariablePrice         graphql.Float
			MaxVariablePrice         graphql.Float
			EurMinVariablePrice      graphql.Float
			EurMaxVariablePrice      graphql.Float
			InternalMinVariablePrice graphql.Float
			InternalMaxVariablePrice graphql.Float
			FullName                 graphql.String
			CustomFullName           graphql.String
			HasDiscount              graphql.Boolean
			IsPremium                graphql.Boolean
			Locale                   graphql.String
			OnlineTc                 graphql.String
			RedeemInstructions   graphql.String
			CustomInstructions   graphql.String
			RedeemSite           graphql.String
			CustomRedeemSite     graphql.String
			RequireMail          graphql.Boolean
			RequirePhone         graphql.Boolean
			RequireOther         graphql.Boolean
			RequiresUserIdentity graphql.Boolean
			Tc                   graphql.String
			CustomOnlineTc       graphql.String
			URLImage             graphql.String
			CustomURLImage       graphql.String
			Countries            []struct {
				ID   graphql.String
				Name graphql.String
			}
			Categories []struct {
				ID   graphql.String
				Name graphql.String
			}
			ProductType        graphql.String
			EmailDelivery      graphql.Boolean
			PrintDelivery      graphql.Boolean
			MailDelivery       graphql.Boolean
			SmsDelivery        graphql.Boolean
			OnlineRedeem       graphql.Boolean
			StoreRedeem        graphql.Boolean
			IsVariant          graphql.Boolean
			VariantProductID   graphql.String
			UatNumber          graphql.String
			ValidationNumber   graphql.String
			RecommendedAmounts graphql.String
			VariantName        graphql.String
			Validity           graphql.String
		}
	} `graphql:"catalog(filter: $filter)"`
}

var productsQuery struct {
	Products []struct {
		ID                       graphql.String
		Available                graphql.Boolean
		Description              graphql.String
		CustomDescription        graphql.String
		DiscountAbsolute         graphql.Float
		DiscountPercentage       graphql.Float
		CustomDiscount           graphql.Float
		Currency                 graphql.String
		InternalCurrency         graphql.String
		IsVariablePrice          graphql.Boolean
		FixedPrice               graphql.Float
		FaceValue                graphql.Float
		EurFixedPrice            graphql.Float
		InternalFixedPrice       graphql.Float
		MinVariablePrice         graphql.Float
		MaxVariablePrice         graphql.Float
		EurMinVariablePrice      graphql.Float
		EurMaxVariablePrice      graphql.Float
		InternalMinVariablePrice graphql.Float
		InternalMaxVariablePrice graphql.Float
		FullName                 graphql.String
		CustomFullName           graphql.String
		HasDiscount              graphql.Boolean
		IsPremium                graphql.Boolean
		Locale                   graphql.String
		OnlineTc                 graphql.String
		// OriginalID               graphql.String
		MetaProvider struct {
			ID   graphql.String
			Name graphql.String
		}
		Provider struct {
			ID               graphql.String
			Image            graphql.String
			Name             graphql.String
			OriginalID       graphql.String
			CustomName       graphql.String
			CustomImage      graphql.String
			ValidationNumber graphql.String
		}
		RedeemInstructions   graphql.String
		CustomInstructions   graphql.String
		RedeemSite           graphql.String
		CustomRedeemSite     graphql.String
		RequireMail          graphql.Boolean
		RequirePhone         graphql.Boolean
		RequireOther         graphql.Boolean
		RequiresUserIdentity graphql.Boolean
		Tc                   graphql.String
		CustomOnlineTc       graphql.String
		URLImage             graphql.String
		CustomURLImage       graphql.String
		Countries            []struct {
			ID   graphql.String
			Name graphql.String
		}
		Categories []struct {
			ID   graphql.String
			Name graphql.String
		}
		ProductType        graphql.String
		EmailDelivery      graphql.Boolean
		PrintDelivery      graphql.Boolean
		MailDelivery       graphql.Boolean
		SmsDelivery        graphql.Boolean
		OnlineRedeem       graphql.Boolean
		StoreRedeem        graphql.Boolean
		IsVariant          graphql.Boolean
		VariantProductID   graphql.String
		UatNumber          graphql.String
		ValidationNumber   graphql.String
		RecommendedAmounts graphql.String
		VariantName        graphql.String
		Validity           graphql.String
		// } `graphql:"products(limit: 3)"`
	} `graphql:"products"`
}

var compactProductsQuery struct {
	Products []struct {
		ID             graphql.String
		Available      graphql.Boolean
		Currency       graphql.String
		FixedPrice     graphql.Float
		FullName       graphql.String
		CustomFullName graphql.String
		MetaProvider   struct {
			Name graphql.String
		}
		Provider struct {
			Name graphql.String
		}
		RedeemSite       graphql.String
		URLImage         graphql.String
		CustomURLImage   graphql.String
		IsVariant        graphql.Boolean
		VariantProductID graphql.String
		// } `graphql:"products(limit: 3)"`
	} `graphql:"products"`
}

var singleProductQuery struct {
	Products []struct {
		ID                       graphql.String
		Available                graphql.Boolean
		Description              graphql.String
		CustomDescription        graphql.String
		DiscountAbsolute         graphql.Float
		DiscountPercentage       graphql.Float
		CustomDiscount           graphql.Float
		Currency                 graphql.String
		InternalCurrency         graphql.String
		IsVariablePrice          graphql.Boolean
		FixedPrice               graphql.Float
		FaceValue                graphql.Float
		EurFixedPrice            graphql.Float
		InternalFixedPrice       graphql.Float
		MinVariablePrice         graphql.Float
		MaxVariablePrice         graphql.Float
		EurMinVariablePrice      graphql.Float
		EurMaxVariablePrice      graphql.Float
		InternalMinVariablePrice graphql.Float
		InternalMaxVariablePrice graphql.Float
		FullName                 graphql.String
		CustomFullName           graphql.String
		HasDiscount              graphql.Boolean
		IsPremium                graphql.Boolean
		Locale                   graphql.String
		OnlineTc                 graphql.String
		// OriginalID               graphql.String
		MetaProvider struct {
			ID   graphql.String
			Name graphql.String
		}
		Provider struct {
			ID               graphql.String
			Image            graphql.String
			Name             graphql.String
			OriginalID       graphql.String
			CustomName       graphql.String
			CustomImage      graphql.String
			ValidationNumber graphql.String
		}
		RedeemInstructions   graphql.String
		CustomInstructions   graphql.String
		RedeemSite           graphql.String
		CustomRedeemSite     graphql.String
		RequireMail          graphql.Boolean
		RequirePhone         graphql.Boolean
		RequireOther         graphql.Boolean
		RequiresUserIdentity graphql.Boolean
		Tc                   graphql.String
		CustomOnlineTc       graphql.String
		URLImage             graphql.String
		CustomURLImage       graphql.String
		Countries            []struct {
			ID   graphql.String
			Name graphql.String
		}
		Categories []struct {
			ID   graphql.String
			Name graphql.String
		}
		ProductType        graphql.String
		EmailDelivery      graphql.Boolean
		PrintDelivery      graphql.Boolean
		MailDelivery       graphql.Boolean
		SmsDelivery        graphql.Boolean
		OnlineRedeem       graphql.Boolean
		StoreRedeem        graphql.Boolean
		IsVariant          graphql.Boolean
		VariantProductID   graphql.String
		UatNumber          graphql.String
		ValidationNumber   graphql.String
		RecommendedAmounts graphql.String
		VariantName        graphql.String
		Validity           graphql.String
	} `graphql:"products(filter: {id: $prodId})"`
}

// Queries for account info
const (
	ACCOUNT_BALANCE string = "ACCOUNT_BALANCE"
	ACCOUNT_INFO           = "ACCOUNT_INFO"
)

var accountInfoQuery struct {
	AccountInfo struct {
		ID            graphql.String
		Name          graphql.String
		AddressStreet graphql.String
		// AddressPC       graphql.Int
		AddressCity    graphql.String
		AddressState   graphql.String
		AddressCountry graphql.String
		MonthlyFee     graphql.Int
		Tc             graphql.String
		// BusinessTaxID   graphql.String
		ContactName     graphql.String
		ContactLastName graphql.String
		ContactTitle    graphql.String
		ContactEmail    graphql.String
		IsPremium       graphql.Boolean
		UserName        graphql.String
		Balance         graphql.Float
		Movements       []struct {
			Amount graphql.Float
			// ClientID     graphql.String
			ID           graphql.String
			MovementType graphql.String
			Note         graphql.String
			// PurchaseID   graphql.String
			Reference graphql.String
			Timestamp graphql.Int
		}
		Provision graphql.Int
		APIKeys   []struct {
			ID      graphql.String
			Key     graphql.String
			IsDev   graphql.Boolean
			IsAdmin graphql.Boolean
			// AllowedIP graphql.String
		}
	}
}

var accountBalanceQuery struct {
	AccountInfo struct {
		Balance graphql.Float
	}
}

type ProductFilter struct {
	Id graphql.ID `json:"id"`
	Country graphql.String `json:"country"`
	Category graphql.ID `json:"category"`
}

// Queries for purchases
const (
	CREATE_ORDER string = "CREATE_ORDER"
	GET_ORDER           = "GET_ORDER"
)

type PurchaseInput struct {
	TransactionID string  `json:"transactionID"`
	ProductID     int     `json:"productID"`
	TotalValue    float64 `json:"totalValue"`
	UserInfo      struct {
		Email            string `json:"email"`
		Name             string `json:"name"`
		Country          string `json:"country"`
		PhoneCountryCode string `json:"phoneCountryCode"`
		PhoneNumber      string `json:"phoneNumber"`
		ServiceNumber    string `json:"serviceNumber"`
	} `json:"userInfo"`
}

var createPurchaseQuery struct {
	CreatePurchase struct {
		ID graphql.String
		// TransactionID graphql.String
		Client struct {
			ID   graphql.String
			Name graphql.String
		}
		TotalValue    graphql.Float
		OriginalValue graphql.Float
		// EndUserName             graphql.String
		// EndUserEmail            graphql.String
		// EndUserCountry          graphql.String
		// EndUserPhoneCountryCode graphql.String
		// EndUserPhoneNumber      graphql.String
		// EndUserSecondNumber     graphql.String
		TimestampRequest   graphql.Int
		TimestampFulfilled graphql.Int
		RedeemCode         graphql.String
		Receipt            graphql.String
		ErrorMessage       graphql.String
		Status             graphql.String
	} `graphql:"createPurchase(purchase: $purchaseInput)"`
}

var getPurchaseQuery struct {
	Purchases []struct {
		ID graphql.String
		// TransactionID graphql.String
		Client struct {
			ID   graphql.String
			Name graphql.String
		}
		TotalValue    graphql.Int
		OriginalValue graphql.Int
		// EndUserName             graphql.String
		// EndUserEmail            graphql.String
		// EndUserCountry          graphql.String
		// EndUserPhoneCountryCode graphql.String
		// EndUserPhoneNumber      graphql.String
		// EndUserSecondNumber     graphql.String
		TimestampRequest   graphql.Int
		TimestampFulfilled graphql.Int
		RedeemCode         graphql.String
		Receipt            graphql.String
		ErrorMessage       graphql.String
		Status             graphql.String
	} `graphql:"purchases(filter: {id: $orderId})"`
}

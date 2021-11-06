package bitcou

import "github.com/hasura/go-graphql-client"

// Queries for products
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

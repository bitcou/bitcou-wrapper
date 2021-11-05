package bitcou_queries

import "github.com/hasura/go-graphql-client"

var productsQuery struct {
	Products []struct {
		Id       graphql.String
		FullName graphql.String
	} `graphql:"products(limit: 3)"`
}

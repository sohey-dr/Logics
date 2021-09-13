package main

import (
	"context"
	"fmt"
	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
	"os"
)

var query struct {
	Me struct {
		Name graphql.String
	}
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GRAPHQL_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := graphql.NewClient("https://example.com/graphql", httpClient)
	// Use client...
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		// Handle error.
	}
	fmt.Println(query.Me.Name)

	// Output: Luke Skywalker
}

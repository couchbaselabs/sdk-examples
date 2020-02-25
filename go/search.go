package main

import (
	"fmt"

	"github.com/couchbase/gocb/v2"
	"github.com/couchbase/gocb/v2/search"
)

func main() {
	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			"Administrator",
			"password",
		},
	}
	cluster, err := gocb.Connect("localhost", opts)
	if err != nil {
		panic(err)
	}

	collection := cluster.Bucket("my-bucket").DefaultCollection()

	matchResult, err := cluster.SearchQuery(
		"travel-sample-index-hotel-description",
		search.NewMatchQuery("swanky"),
		&gocb.SearchOptions{
			Limit: 10,
		},
	)
	fmt.Println(matchResult)

	dateRangeResult, err := cluster.SearchQuery(
		"travel-sample-index-hotel-description",
		search.NewDateRangeQuery().Start("2019-01-01", true).End("2019-02-01", false),
		&gocb.SearchOptions{
			Limit: 10,
		},
	)
	fmt.Println(dateRangeResult)

	conjunctionResult, err := cluster.SearchQuery(
		"travel-sample-index-hotel-description",
		search.NewConjunctionQuery(
			search.NewMatchQuery("swanky"),
			search.NewDateRangeQuery().Start("2019-01-01", true).End("2019-02-01", false),
		),
		&gocb.SearchOptions{
			Limit: 10,
		},
	)
	fmt.Println(conjunctionResult)

	for matchResult.Next() {
		row := matchResult.Row()
		docID := row.ID
		score := row.Score

		var fields interface{}
		err := row.Fields(&fields)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Document ID: %s, search score: %f, fields included in result: %v\n", docID, score, fields)
	}

	// always check for errors after iterating
	err = matchResult.Err()
	if err != nil {
		panic(err)
	}

	facets, err := matchResult.Facets()
	if err != nil {
		panic(err)
	}
	for _, facet := range facets {
		field := facet.Field
		total := facet.Total

		fmt.Printf("Facet field: %s, total: %d\n", field, total)
	}

	hotel := struct {
		Description string `json:"description"`
	}{Description: "super swanky"}
	myWriteResult, err := collection.Upsert("a-new-hotel", hotel, nil)
	if err != nil {
		panic(err)
	}

	consistentWithResult, err := cluster.SearchQuery(
		"travel-sample-index-hotel-description",
		search.NewMatchQuery("swanky"),
		&gocb.SearchOptions{
			Limit:          10,
			ConsistentWith: gocb.NewMutationState(*myWriteResult.MutationToken()),
		},
	)
	fmt.Println(consistentWithResult)
}

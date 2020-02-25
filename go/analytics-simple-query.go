package main

import (
	"fmt"

	"github.com/couchbase/gocb/v2"
)

func main() {
	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			"Administrator",
			"password",
		},
	}
	cluster, err := gocb.Connect("10.112.194.101", opts)
	if err != nil {
		panic(err)
	}

	results, err := cluster.AnalyticsQuery("select airportname, country from airports where country = 'France';", nil)
	if err != nil {
		panic(err)
	}

	var rows []interface{}
	for results.Next() {
		var row interface{}
		if err := results.Row(&row); err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}

	if err := results.Err(); err != nil {
		panic(err)
	}

	fmt.Println(rows)

	// make sure that results has been iterated (and therefore closed) before calling this.
	metadata, err := results.MetaData()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Client context id: %s\n", metadata.ClientContextID)
	fmt.Printf("Elapsed time: %d\n", metadata.Metrics.ElapsedTime)
	fmt.Printf("Execution time: %d\n", metadata.Metrics.ExecutionTime)
	fmt.Printf("Result count: %d\n", metadata.Metrics.ResultCount)
	fmt.Printf("Error count: %d\n", metadata.Metrics.ErrorCount)

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

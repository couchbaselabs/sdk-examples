package main

import (
	"fmt"
	"time"

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

	collection := cluster.Bucket("travel-sample").DefaultCollection()

	// Upsert
	mops := []gocb.MutateInSpec{
		gocb.UpsertSpec("fax", "311-555-0151", &gocb.UpsertSpecOptions{}),
	}
	upsertResult, err := collection.MutateIn("customer123", mops, &gocb.MutateInOptions{
		Timeout: 50 * time.Millisecond,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(upsertResult.Cas())

	// Insert
	mops = []gocb.MutateInSpec{
		gocb.InsertSpec("purchases.complete", []interface{}{32, true, "None"}, &gocb.InsertSpecOptions{}),
	}
	insertResult, err := collection.MutateIn("customer123", mops, &gocb.MutateInOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.Cas())

	// Multiple specs
	mops = []gocb.MutateInSpec{
		gocb.RemoveSpec("addresses.billing[2]", nil),
		gocb.ReplaceSpec("email", "dougr96@hotmail.com", nil),
	}
	multiMutateResult, err := collection.MutateIn("customer123", mops, &gocb.MutateInOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(multiMutateResult.Cas())

	// Create path
	mops = []gocb.MutateInSpec{
		gocb.UpsertSpec("level_0.level_1.foo.bar.phone", map[string]interface{}{
			"num": "311-555-0101",
			"ext": 16,
		}, &gocb.UpsertSpecOptions{
			CreatePath: true,
		}),
	}
	createPathUpsertResult, err := collection.MutateIn("customer123", mops, &gocb.MutateInOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(createPathUpsertResult.Cas())

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

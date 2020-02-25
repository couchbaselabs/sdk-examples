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

	collection := cluster.Bucket("travel-sample").DefaultCollection()

	// Observe based
	mops := []gocb.MutateInSpec{
		gocb.InsertSpec("name", "mike", nil),
	}
	observeResult, err := collection.MutateIn("key", mops, &gocb.MutateInOptions{
		PersistTo:   1,
		ReplicateTo: 1,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(observeResult.Cas())

	// Enhanced
	mops = []gocb.MutateInSpec{
		gocb.InsertSpec("name", "mike", nil),
	}
	durableResult, err := collection.MutateIn("key", mops, &gocb.MutateInOptions{
		DurabilityLevel: gocb.DurabilityLevelMajority,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(durableResult.Cas())

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

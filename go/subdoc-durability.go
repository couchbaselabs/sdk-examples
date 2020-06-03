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
	cluster, err := gocb.Connect("localhost", opts)
	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket("default")
	collection := bucket.DefaultCollection()

	// We wait until the bucket is definitely connected and setup.
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	// Observe based
	mops := []gocb.MutateInSpec{
		gocb.InsertSpec("name", "mike", nil),
	}
	observeResult, err := collection.MutateIn("key", mops, &gocb.MutateInOptions{
		PersistTo:     1,
		ReplicateTo:   1,
		StoreSemantic: gocb.StoreSemanticsUpsert,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(observeResult.Cas())

	// Enhanced
	mops = []gocb.MutateInSpec{
		gocb.InsertSpec("name", "mike", nil),
	}
	durableResult, err := collection.MutateIn("key2", mops, &gocb.MutateInOptions{
		DurabilityLevel: gocb.DurabilityLevelMajority,
		StoreSemantic:   gocb.StoreSemanticsUpsert,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(durableResult.Cas())

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

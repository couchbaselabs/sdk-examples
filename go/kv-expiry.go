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

	type myDoc struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	document := myDoc{Foo: "bar", Bar: "foo"}

	key := "document-key"
	// Upsert with Expiry
	expiryResult, err := collection.Upsert(key, &document, &gocb.UpsertOptions{
		Timeout: 25 * time.Millisecond,
		Expiry:  60 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(expiryResult)

	getRes, err := collection.Get(key, nil)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Expiry value: %d\n", getRes.Expiry())

	// Touch
	touchResult, err := collection.Touch(key, 60*time.Second, &gocb.TouchOptions{
		Timeout: 100 * time.Millisecond,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(touchResult)

	// GetAndTouch
	getAndTouchResult, err := collection.GetAndTouch(key, 60, &gocb.GetAndTouchOptions{
		Timeout: 100 * time.Millisecond,
	})
	if err != nil {
		panic(err)
	}

	var getAndTouchDoc myDoc
	err = getAndTouchResult.Content(&getAndTouchDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println(getAndTouchDoc)
	fmt.Printf("Expiry value: %d\n", getRes.Expiry())

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

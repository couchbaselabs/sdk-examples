package main

import (
	"github.com/couchbase/gocb/v2"
	"github.com/pquerna/ffjson/ffjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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

	bucket := cluster.Bucket("bucket-name")

	collection := bucket.DefaultCollection()

	// Create a new raw transcoder and use it to Upsert the document.
	user := User{Name: "John Smith", Age: 27}
	transcoder := gocb.NewRawJSONTranscoder()

	b, err := ffjson.Marshal(user)
	if err != nil {
		panic(err)
	}

	_, err = collection.Upsert("john-smith", b, &gocb.UpsertOptions{
		Transcoder: transcoder,
	})
	if err != nil {
		panic(err)
	}

	// Get the document and unmarshal it using the same transcoder.
	getRes, err := collection.Get("john-smith", &gocb.GetOptions{
		Transcoder: transcoder,
	})
	if err != nil {
		panic(err)
	}

	var returned []byte
	err = getRes.Content(&returned)
	if err != nil {
		panic(err)
	}

	err = ffjson.Unmarshal(returned, &user)
	if err != nil {
		panic(err)
	}

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"time"

	gocb "github.com/couchbase/gocb/v2"
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

	// get a bucket reference
	bucket := cluster.Bucket("travel-sample")

	// We wait until the bucket is definitely connected and setup.
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	landmarksResult, err := bucket.ViewQuery("landmarks", "by_name", &gocb.ViewOptions{
		Key:       "<landmark-name>",
		Namespace: gocb.DesignDocumentNamespaceDevelopment,
	})
	if err != nil {
		panic(err)
	}

	for landmarksResult.Next() {
		landmarkRow := landmarksResult.Row()
		fmt.Printf("Document ID: %s\n", landmarkRow.ID)
		var key string
		err = landmarkRow.Key(&key)
		if err != nil {
			panic(err)
		}

		var landmark interface{}
		err = landmarkRow.Value(&landmark)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Landmark named %s has value %v\n", key, landmark)
	}

	// always check for errors after iterating
	err = landmarksResult.Err()
	if err != nil {
		panic(err)
	}

	if err := cluster.Close(nil); err != nil {
		panic(err)
	}
}

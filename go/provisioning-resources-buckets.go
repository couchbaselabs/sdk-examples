package main

import (
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

	// For Server versions 6.5 or later you do not need to open a bucket here
	b := cluster.Bucket("travel-sample")

	// We wait until the bucket is definitely connected and setup.
	// For Server versions 6.5 or later if we hadn't opened a bucket then we could use cluster.WaitUntilReady here.
	err = b.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	bucketMgr := cluster.Buckets()


	createBucket(bucketMgr)
	updateBucket(bucketMgr)
	flushBucket(bucketMgr)
	removeBucket(bucketMgr)
}

func createBucket(bucketMgr *gocb.BucketManager) {
	err := bucketMgr.CreateBucket(gocb.CreateBucketSettings{
		BucketSettings: gocb.BucketSettings{
			Name:                 "hello",
			FlushEnabled:         false,
			ReplicaIndexDisabled: true,
			RAMQuotaMB:           200,
			NumReplicas:          1,
			BucketType:           gocb.CouchbaseBucketType,
		},
		ConflictResolutionType: gocb.ConflictResolutionTypeSequenceNumber,
	}, nil)
	if err != nil {
		panic(err)
	}
}

func updateBucket(bucketMgr *gocb.BucketManager) {
	settings, err := bucketMgr.GetBucket("hello", nil)
	if err != nil {
		panic(err)
	}

	settings.FlushEnabled = true
	err = bucketMgr.UpdateBucket(*settings, nil)
	if err != nil {
		panic(err)
	}
}

func removeBucket(bucketMgr *gocb.BucketManager) {
	err := bucketMgr.DropBucket("hello", nil)
	if err != nil {
		panic(err)
	}
}

func flushBucket(bucketMgr *gocb.BucketManager) {
	err := bucketMgr.FlushBucket("hello", nil)
	if err != nil {
		panic(err)
	}
}

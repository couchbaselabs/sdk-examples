package main

import (
	"github.com/couchbase/gocb/v2"
)

func main() {
	opts := gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			"Administrator",
			"password",
		},
	}

	cluster, err := gocb.Connect("10.112.193.101", opts)
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
			RAMQuotaMB:           1024,
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
	settings, err := bucketMgr.GetBucket("test", nil)
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
	err := bucketMgr.DropBucket("test", nil)
	if err != nil {
		panic(err)
	}
}

func flushBucket(bucketMgr *gocb.BucketManager) {
	err := bucketMgr.FlushBucket("test", nil)
	if err != nil {
		panic(err)
	}
}

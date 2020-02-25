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

	cluster, err := gocb.Connect("10.112.193.101", opts)
	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket("default")
	viewMgr := bucket.ViewIndexes()

	createView(viewMgr)
	getView(viewMgr)
	publishView(viewMgr)
	removeView(viewMgr)
}

func createView(viewMgr *gocb.ViewIndexManager) {
	designDoc := gocb.DesignDocument{
		Name: "landmarks",
		Views: map[string]gocb.View{
			"by_country": {
				Map:    "function (doc, meta) { if (doc.type == 'landmark') { emit([doc.country, doc.city], null); } }",
				Reduce: nil,
			},
			"by_activity": {
				Map:    "function (doc, meta) { if (doc.type == 'landmark') { emit(doc.activity, null); } }",
				Reduce: "_count",
			},
		},
	}

	err := viewMgr.UpsertDesignDocument(designDoc, gocb.DesignDocumentNamespaceDevelopment, nil)
	if err != nil {
		panic(err)
	}
}

func getView(viewMgr *gocb.ViewIndexManager) {
	ddoc, err := viewMgr.GetDesignDocument("landmarks", gocb.DesignDocumentNamespaceDevelopment, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(ddoc)
}

func publishView(viewMgr *gocb.ViewIndexManager) {
	err := viewMgr.PublishDesignDocument("landmarks", nil)
	if err != nil {
		panic(err)
	}
}

func removeView(viewMgr *gocb.ViewIndexManager) {
	err := viewMgr.DropDesignDocument("landmarks", gocb.DesignDocumentNamespaceProduction, nil)
	if err != nil {
		panic(err)
	}
}

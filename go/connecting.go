package main

import (
	"crypto/x509"
	"time"

	gocb "github.com/couchbase/gocb/v2"
)

func simpleconnect() {
	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	}
	cluster, err := gocb.Connect("10.112.193.101", opts)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func multinodeconnect() {
	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	}
	cluster, err := gocb.Connect("192.168.56.101,192.168.56.102", opts)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func customports() {
	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	}
	cluster, err := gocb.Connect("couchbase://192.168.56.101:1234,192.168.56.102:5678", opts)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func tlsconnect() {
	// We use the system certificate pool here and assume the Couchbase root certificate(s) have
	// been installed there, but it is also possible to load these from a file.
	rootCAs, err := x509.SystemCertPool()
	if err != nil {
		panic(err)
	}

	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
		SecurityConfig: gocb.SecurityConfig{
			TLSRootCAs: rootCAs,
		},
	}
	cluster, err := gocb.Connect("couchbases://10.112.193.101", opts)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func dnssrv() {
	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	}
	cluster, err := gocb.Connect("couchbase://mysrvrecord.hostname.com", opts)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func waitUntilReady() {
	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	}
	cluster, err := gocb.Connect("couchbase://10.112.193.101", opts)
	if err != nil {
		panic(err)
	}

	err = cluster.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func waitUntilReadyBucket() {
	opts := gocb.ClusterOptions{
		Username: "Administrator",
		Password: "password",
	}
	cluster, err := gocb.Connect("couchbase://10.112.193.101", opts)
	if err != nil {
		panic(err)
	}

	bucket := cluster.Bucket("default")

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		panic(err)
	}

	cluster.Close(nil)
}

func main() {
	simpleconnect()
	multinodeconnect()
	customports()
	tlsconnect()
	dnssrv()
	waitUntilReady()
}

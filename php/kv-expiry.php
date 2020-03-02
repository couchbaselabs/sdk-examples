<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\UpsertOptions;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$collection = $cluster->bucket("travel-sample")->defaultCollection();

$document = [
    "foo" => "bar",
    "bar" => "foo",
];

$key = "document-key";

// Upsert with Expiry
$opts = new UpsertOptions();
$opts->expiry(60 /* seconds */);
$res = $collection->upsert($key, $document, $opts);

// Retrieve the document immediately, must be exist
$res = $collection->get($key);
printf("[get] document content: %s\n", var_export($res->content(), true));

// Touch the document to adjust expiration time
$collection->touch($key, 60 /* seconds */);

// Get and touch retrieves the document and adjusting expiration time
$res = $collection->getAndTouch($key, 1 /* seconds */);
printf("[getAndTouch] document content: %s\n", var_export($res->content(), true));

sleep(2); // wait until the document will expire

try {
    $collection->get($key);
} catch (Couchbase\DocumentNotFoundException $ex) {
    printf("The document does not exist\n");
}

// Output:
//
//     [get] document content: array (
//       'foo' => 'bar',
//       'bar' => 'foo',
//     )
//     [getAndTouch] document content: array (
//       'foo' => 'bar',
//       'bar' => 'foo',
//     )
//     The document does not exist

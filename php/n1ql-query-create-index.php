<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\QueryOptions;
use \Couchbase\CreateQueryPrimaryIndexOptions;
use \Couchbase\CreateQueryIndexOptions;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$idxMgr = $cluster->queryIndexes();

// Create a primary index
$opts = new CreateQueryPrimaryIndexOptions();
$opts->ignoreIfExists(true);
$idxMgr->createPrimaryIndex("travel-sample", $opts);

// Create a deferred named index
$opts = new CreateQueryIndexOptions();
$opts->ignoreIfExists(true);
$opts->deferred(true);
$idxMgr->createIndex("travel-sample", "my_index", ["name"], $opts);

// Build deferred indexes
$res = $idxMgr->builddeferredIndexes("travel-sample");
var_dump($res);

// Drop the primary index
$idxMgr->dropPrimaryIndex("travel-sample");

// Drop the named index
$idxMgr->dropIndex("travel-sample", "my_index");

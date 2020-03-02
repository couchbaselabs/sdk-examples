<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\UpsertOptions;
use \Couchbase\DurabilityLevel;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$collection = $cluster->bucket("travel-sample")->defaultCollection();

$document = ["foo" => "bar", "bar" => "foo"];

// Upsert with Durability level Majority
$opts = new UpsertOptions();
$opts->durabilityLevel(DurabilityLevel::MAJORITY);
$res = $collection->upsert("durable-key", $document, $opts);
printf("Successfully created document \"durable-key\". CAS=\"%s\"\n", $res->cas());

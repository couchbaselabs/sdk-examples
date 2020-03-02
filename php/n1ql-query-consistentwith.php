<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\QueryOptions;
use \Couchbase\MutationState;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$collection = $cluster->bucket('travel-sample')->defaultCollection();

// create/update document (mutation)
$res = $collection->upsert("id", ["name" => "somehotel", "type" => "hotel"]);

// create mutation state from mutation results
$state = new MutationState();
$state->add($res);

// use mutation state with query optionss
$opts = new QueryOptions();
$opts->consistentWith($state);
$res = $cluster->query('SELECT x.* FROM `travel-sample` x WHERE x.`type`="hotel" AND x.name LIKE "%hotel%" LIMIT 10', $opts);

$idx = 1;
foreach ($res->rows() as $row) {
    printf("%d. %s\n", $idx++, $row['name']);
}

printf("Execution Time: %d\n", $res->metaData()->metrics()['executionTime']);

// Output:
//
//     1. The Falcondale hotel and restaurant
//     2. myhotel Chelsea
//     3. pentahotel Birmingham
//     4. somehotel
//     Execution Time: 286


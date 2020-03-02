<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\QueryOptions;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

// NOTE: we are using single-quotes here to avoid PHP $-expansion
$query = 'SELECT x.* FROM `travel-sample` x WHERE x.`type`=$type LIMIT 10';
$opts = new QueryOptions();
$opts->namedParameters(["type" => "hotel"]);
$res = $cluster->query($query, $opts);

$idx = 1;
foreach ($res->rows() as $row) {
    printf("%d. %s, \"%s\"\n", $idx++, $row['country'], $row['name']);
}

printf("Execution Time: %d\n", $res->metaData()->metrics()['executionTime']);

// Output:
//
//     1. United Kingdom, "Medway Youth Hostel"
//     2. United Kingdom, "The Balmoral Guesthouse"
//     3. France, "The Robins"
//     4. France, "Le Clos Fleuri"
//     5. United Kingdom, "Glasgow Grand Central"
//     6. United Kingdom, "Glencoe Youth Hostel"
//     7. United Kingdom, "The George Hotel"
//     8. United Kingdom, "Windy Harbour Farm Hotel"
//     9. United Kingdom, "Avondale Guest House"
//     10. United Kingdom, "The Bulls Head"
//     Execution Time: 7


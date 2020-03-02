<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\ViewOptions;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$bucket = $cluster->bucket("beer-sample");

$opts = new ViewOptions();
$opts->range("b", null);
$opts->limit(10);
$res = $bucket->viewQuery("beer", "brewery_beers", $opts);

foreach ($res->rows() as $row) {
    printf("id: %s, key: %s, value: %s\n", $row->id(), json_encode($row->key()), json_encode($row->value()));
}

// Output:
//
//
//     id: landmark_16359, key: Albert Memorial, value: array (
//       'accuracy' => 'APPROXIMATE',
//       'lat' => 51.50238,
//       'lon' => -0.17771,
//     )


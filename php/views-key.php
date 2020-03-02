<?php
use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\ViewOptions;

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$bucket = $cluster->bucket("travel-sample");

/*
 * The view represented as single view function like one below:
 *
 *     function (doc, meta) {
 *       if (doc.type == "landmark") {
 *       	emit(doc.name, doc.geo);
 *       }
 *     }
 */
$opts = new ViewOptions();
$opts->key("Albert Memorial");
$res = $bucket->viewQuery("landmarks", "by_name", $opts);

foreach ($res->rows() as $row) {
    printf("id: %s, key: %s, value: %s\n", $row->id(), $row->key(), var_export($row->value(), true));
}

// Output:
//
//
//     id: landmark_16359, key: Albert Memorial, value: array (
//       'accuracy' => 'APPROXIMATE',
//       'lat' => 51.50238,
//       'lon' => -0.17771,
//     )


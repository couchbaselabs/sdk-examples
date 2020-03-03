<?php
$connectionString = "couchbase://10.112.193.101";
$options = new \Couchbase\ClusterOptions();
$options->credentials("Administrator", "password");
$cluster = new \Couchbase\Cluster($connectionString, $options);

$bucket = $cluster->bucket("travel-sample");
$collection = $bucket->defaultCollection();

$options = new \Couchbase\AnalyticsQueryOptions();
$result = $cluster->analyticsQuery('SELECT "hello" as greeting;', $options);

foreach($result->rows() as $row) {
    printf("result: %s\n", $row["greeting"]);
}

$options = new \Couchbase\AnalyticsQueryOptions();
$result = $cluster->analyticsQuery('SELECT airportname, country FROM airports WHERE country = "France";', $options);

$options = new \Couchbase\AnalyticsQueryOptions();
$options->positionalParameters(["France"]);
$result = $cluster->analyticsQuery('SELECT airportname, country FROM airports WHERE country = $1;', $options);

$options = new \Couchbase\AnalyticsQueryOptions();
$options->namedParameters(['$country' => "France"]);
$result = $cluster->analyticsQuery('SELECT airportname, country FROM airports WHERE country = $country;', $options);

$options = new \Couchbase\AnalyticsQueryOptions();
$options->timeout(100);
$result = $cluster->analyticsQuery('SELECT airportname, country FROM airports WHERE country = "France";', $options);

$options = new \Couchbase\AnalyticsQueryOptions();
$result = $cluster->analyticsQuery('SELECT airportname, country FROM airports WHERE country = "France";', $options);

foreach($result->rows() as $row) {
printf("Name: %s, Country: %s\n", $row["name"], $row["country"]);
}

$options = new \Couchbase\AnalyticsQueryOptions();
$result = $cluster->analyticsQuery('SELECT airportname, country FROM airports WHERE country = "France";', $options);

$metadata = $result->metadata();
$metrics = $metadata->metrics();
printf("Elapsed time: %d\n", $metrics["elapsedTime"]);
printf("Execution time: %d\n", $metrics["executionTime"]);
printf("Result count: %d\n", $metrics["resultCount"]);

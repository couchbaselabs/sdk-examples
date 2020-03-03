<?php

use \Couchbase\ClusterOptions;
use \Couchbase\Cluster;
use \Couchbase\Collection;
use \Couchbase\ReplaceOptions;
use \Couchbase\CasMismatchError;

function incrementVisitCount(Collection $collection, string $userId) {
    $maxRetries = 10;
    for ($i = 0; $i < $maxRetries; $i++) {
        // Get the current document contents
        $res = $collection->get($userId);

        // Increment the visit count
        $user = $res->content();
        $user["visit_count"]++;

        try {
            // Attempt to replace the document using CAS
            $opts = new ReplaceOptions();
            $opts->cas($res->cas());
            $collection->replace($userId, $user, $opts);
        } catch (CasMismatchError $ex) {
            continue;
        }

        // If no errors occured during the replace, we can exit our retry loop
        return;
    }
    printf("Replace failed after %d attempts\n", $maxRetries);
}


function lockingAndCas(Collection $collection, string $userId) {
    $res = $collection->getAndLock($userId, 2 /* seconds */);
    $lockedCas = $res->cas();

    /* // an example of simply unlocking the document:
     * $collection->unlock($userId, $lockedCas);
     */

    // Increment the visit count
    $user = $res->content();
    $user["visit_count"]++;

    $opts = new ReplaceOptions();
    $opts->cas($lockedCas);
    $collection->replace($userId, $user, $opts);
}

$opts = new ClusterOptions();
$opts->credentials("Administrator", "password");
$cluster = new Cluster("couchbase://192.168.1.101", $opts);

$bucket = $cluster->bucket("default");
$collection = $bucket->defaultCollection();

$collection->upsert("userId", ["visit_count" => 0]);

replaceWithCas($collection, "userId");
lockingAndCas($collection, "userId");

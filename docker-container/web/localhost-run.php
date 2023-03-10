<?php
/*
 * This PHP-file proxies between the localhost-run docker container and your website.
 * It will try to read the address of the localhost-run host
 * 
 * Author: Mats Otten
 * Date: 31 july 2022
 */

header('Content-Type: application/json; charset=utf-8');

// Tries to fetch the contents of the log file from the server, if something goes wrong we going to ignore this.
$jsonDataContents = @file_get_contents("http://localhost-run:8000/localhost-run");

// The log-file can contain multiple JSON-objects, we match all of these.
preg_match_all("/({.*?})/", $jsonDataContents, $forwardMatches);

// Check if there is atleast one JSON-object, if available we need the data from the latest one.
if(count($forwardMatches[0]) > 0) {
    if($latest = json_decode($forwardMatches[0][count($forwardMatches[0]) - 1], true)) {
        if(isset($latest['address'])) {
            die(json_encode(['success' => true, 'address' => $latest['address']]));
        } else {
            die(json_encode(['success' => false, 'message' => 'Cannot find address in the data: ' . $forwardMatches[0][count($forwardMatches[0]) - 1]]));
        }
    }
}

die(json_encode(['success' => false, 'message' => 'Cannot find localhost-run log-file.']));
?>
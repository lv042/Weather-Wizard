/*
* Author: Luca von Kannen
* Date: March 7, 2023
* Description: This file contains the PHP code to ping the wemos
* License: MIT License
*/

<?php

class WemosPing extends Thread
{
    public $wemos_connected;
    private $url;
    private $timeout;
    private $interval;

    public function __construct($url, $timeout = 5, $interval = 60)
    {
        $this->url = $url;
        $this->timeout = $timeout;
        $this->interval = $interval;
        $this->wemos_connected = false;
    }

    public function run()
    {
        while (true) {
            $this->ping();
            sleep($this->interval);
        }
    }

    private function ping()
    {
        // Initialize a new cURL session
        $curl = curl_init();

        // Set the cURL options
        curl_setopt($curl, CURLOPT_URL, $this->url);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($curl, CURLOPT_CONNECTTIMEOUT, $this->timeout);
        curl_setopt($curl, CURLOPT_TIMEOUT, $this->timeout);

        // Execute the cURL request and get the response
        $response = curl_exec($curl);

        // Check for errors or timeouts
        if (curl_errno($curl)) {
            $this->wemos_connected = false;
        } else {
            $http_code = curl_getinfo($curl, CURLINFO_HTTP_CODE);
            if ($http_code == 200) {
                $this->wemos_connected = true;
                echo "Wemos is connected\n";
            } else {
                $this->wemos_connected = false;
                echo "Wemos is not connected\n";
            }
        }

        // Close the cURL session
        curl_close($curl);
    }
}

?>

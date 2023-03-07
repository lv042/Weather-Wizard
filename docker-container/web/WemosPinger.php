<?php

//This class is used to ping the Wemos in a seperate thread to check if it is connected or not
class WemosPinger {
    private $url;
    private $timeout;
    private $interval;
    public $wemos_connected;

    public function __construct($url, $timeout=5, $interval=60) {
        $this->url = $url;
        $this->timeout = $timeout;
        $this->interval = $interval;
        $this->wemos_connected = false;
    }

    public function ping() {
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
        if(curl_errno($curl)) {
            $this->wemos_connected = false;
        } else {
            $http_code = curl_getinfo($curl, CURLINFO_HTTP_CODE);
            if ($http_code == 200) {
                $this->wemos_connected = true;
                echo "Wemos is connected";
            } else {
                $this->wemos_connected = false;
                echo "Wemos is not connected";
            }
        }

        // Close the cURL session
        curl_close($curl);
    }

    public function start() {
        // Create a new thread for the pinging
        $thread = new Thread(function() {
            while (true) {
                $this->ping();
                sleep($this->interval);
            }
        });

        // Start the thread
        $thread->start();
    }
}


?>
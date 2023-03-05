<?php

// Check if the script is being executed as part of an HTTP request
if (isset($_SERVER['REQUEST_METHOD'])) {

    // Define the API route
    if ($_SERVER['REQUEST_METHOD'] == 'GET' && strpos($_SERVER['REQUEST_URI'], 'weather_data') !== false && $_SERVER['REMOTE_ADDR'] == '127.0.0.1') {

        // Create an array to store the data
        $data = array(
            array(
                "id" => 1,
                "temperature" => 25.6,
                "humidity" => 65.3,
                "pressure" => 1013.2,
                "obstacle_detected" => false,
                "light_intensity" => 124.5
            ),
            array(
                "id" => 2,
                "temperature" => 24.3,
                "humidity" => 68.9,
                "pressure" => 1011.8,
                "obstacle_detected" => true,
                "light_intensity" => 78.9
            ),
            array(
                "id" => 3,
                "temperature" => 23.1,
                "humidity" => 71.2,
                "pressure" => 1009.4,
                "obstacle_detected" => false,
                "light_intensity" => 42.1
            )
        );

        // Set the content type header to JSON
        header('Content-Type: application/json');

        // Encode the data as JSON and output it
        echo json_encode($data);
    }
}
else{
    echo "Error: This script is not being executed as part of an HTTP request.";
}
?>

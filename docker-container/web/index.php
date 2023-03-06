<?php
include 'db.php';

//set up the database connection
$db = new Database("mariadb", "root", "7YKyE8R2AhKzswfN", "WS");

if (isset($_SERVER['REQUEST_METHOD'])) {
    if ($_SERVER['REQUEST_METHOD'] == 'PUT') {
        // Get the request body and decode it as JSON
        $json = file_get_contents('php://input');
        $data = json_decode($json, true);

        // Check if the JSON decoding failed
        if ($data === null) {
            $response = array('error' => 'Invalid JSON format.');
            http_response_code(400);
            echo json_encode($response);
            die();
        }

        // Extract the values from the JSON data
        $timestamp = $data['timestamp'];
        $temperature = $data['temperature'];
        $humidity = $data['humidity'];
        $pressure = $data['pressure'];
        $obstacle_detected = $data['obstacle_detected'];
        $light_intensity = $data['light_intensity'];

        // Add the new datapoint to the weather_data table
        $db->addWeatherData($timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity);

        // Return a success message
        $response = array('message' => 'Weather data added successfully.');
        echo json_encode($response);
        die();
    }

    if ($_SERVER['REQUEST_METHOD'] == 'GET' && $_SERVER['REQUEST_URI'] == '/') {
        // Return the index.html file
        header('Content-type: text/html');
        readfile('index.html');
        die();
    }

    else if ($_GET['action'] == 'weather_data') {
        $result = $db->getWeatherData();
        echo $result;
        die();
    }
}
else {
    echo "Error: This script is not being executed as part of an HTTP request.";
    die();
}



$db->close();
?>

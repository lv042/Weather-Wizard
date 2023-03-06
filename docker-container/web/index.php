<?php
include 'db.php';

//set up the database connection
$db = new Database("mariadb", "root", "7YKyE8R2AhKzswfN", "WS");







if (isset($_SERVER['REQUEST_METHOD'])) {

    //PUT request

    if ($_SERVER['REQUEST_METHOD'] == 'PUT') {
        // Get the request body and decode it as JSON
        $json = file_get_contents('php://input');
        $data = json_decode($json, true);

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


    //GET request

    if ($_GET['action'] == 'weather_data') {
        $result = $db->getWeatherData();
        echo $result;
        die();
    }
}
else {
    echo "Error: This script is not being executed as part of an HTTP request.";
}



$db->close();

?>

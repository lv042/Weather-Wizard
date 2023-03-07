<?php
include 'db.php';

//set up the database connection
$db = new Database("mariadb", "root", "7YKyE8R2AhKzswfN", "WS");

//check if the request method is set
if (isset($_SERVER['REQUEST_METHOD'])) {

    //Post operation to add data from the wemos
    if ($_SERVER['REQUEST_METHOD'] == 'POST') {
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
        $success = $db->addWeatherData($timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity);

        if (!$success) {
            $response = array('error' => 'Error adding weather data.');
            http_response_code(400);
            echo json_encode($response);
            die();
        }

        // Return a success message
        $response = array('message' => 'Weather data added successfully.');
        echo json_encode($response);
        die();
    }

    //Put operation for updating weather data
    else if ($_SERVER['REQUEST_METHOD'] == 'PUT') {
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
        $new_timestamp = $data['new_timestamp'];

        $temperature = $data['temperature'];
        $humidity = $data['humidity'];
        $pressure = $data['pressure'];
        $obstacle_detected = $data['obstacle_detected'];
        $light_intensity = $data['light_intensity'];


        // Update the existing datapoint in the weather_data table
        $success = $db->updateWeatherData($timestamp,$new_timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity);

        if (!$success) {
            $response = array('error' => 'Error updating weather data.');
            http_response_code(400);
            echo json_encode($response);
            die();
        }

        // Return a success message
        $response = array('message' => 'Weather data updated successfully.');
        echo json_encode($response);
        die();
    }

    //Delete operation
    else if ($_SERVER['REQUEST_METHOD'] == 'DELETE') {
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

    // Delete the datapoint from
        // Delete the datapoint from the weather_data table
        $success = $db->deleteWeatherData($timestamp);

        if(!$success) {
            $response = array('error' => 'Error deleting weather data.');
            http_response_code(400);
            echo json_encode($response);
            die();
        }

        // Return a success message
        $response = array('message' => 'Weather data deleted successfully.');
        echo json_encode($response);
        die();
    }
    // Serve the main page
    else if ($_SERVER['REQUEST_METHOD'] == 'GET' && $_SERVER['REQUEST_URI'] == '/') {
        //Serves the main page
        header('Content-type: text/html');
        readfile('index.html');
        die();
    }
    //Get operation for the weather data for the frontend
    else if ($_GET['action'] == 'weather_data') {
        // Return the weather data
        $result = $db->getWeatherData();
        echo $result;
        die();
    }
}
//Error handling
else {
    //When the backend is not on a server
    echo "Error: This script is not being executed as part of an HTTP request.";
    die();
}
//Close the database connection
$db->close();
?>


<?php
class Database {
    private $servername;
    private $username;
    private $password;
    private $dbname;
    private $conn;

    public function __construct($servername, $username, $password, $dbname) {
        $this->servername = $servername;
        $this->username = $username;
        $this->password = $password;
        $this->dbname = $dbname;
        $this->connect();
    }

    private function connect() {
        $this->conn = new mysqli($this->servername, $this->username, $this->password, $this->dbname);
        if ($this->conn->connect_error) {
            die("Connection failed: " . $this->conn->connect_error);
        }
    }

    public function query($sql) {
        $result = $this->conn->query($sql);
        $data = array();
        while ($row = $result->fetch_assoc()) {
            $data[] = $row;
        }
        return json_encode($data);
    }

    public function getWeatherData() {
        $sql = "SELECT * FROM weather_data";
        $result = $this->conn->query($sql);
        $data = array();
        while ($row = $result->fetch_assoc()) {
            $data[] = $row;
        }


        return json_encode($data);
    }

    public function deleteWeatherData($id) {
        $sql = "SELECT * FROM weather_data WHERE timestamp = '$id'";
        $result = $this->conn->query($sql);

        if (!$result) {
            die("Error checking weather data: " . $this->conn->error);
        }

        if ($result->num_rows == 0) {
            return false;
        }

        $sql = "DELETE FROM weather_data WHERE timestamp ='$id'";
        $result = $this->conn->query($sql);

        if (!$result) {
            die("Error deleting weather data: " . $this->conn->error);
        }

        return true;
    }

    public function addWeatherData($timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity) {
        // Check if the record already exists
        $check_sql = "SELECT COUNT(*) AS count FROM weather_data WHERE timestamp = '$timestamp'";
        $check_result = $this->conn->query($check_sql);
        if (!$check_result) {
            die("Error checking weather data: " . $this->conn->error);
        }
        $count = $check_result->fetch_assoc()['count'];
        if ($count > 0) {
            // Record already exists, do not add
            return false;
        }

        // Record does not exist, add it
        $sql = "INSERT INTO weather_data (timestamp, temperature, humidity, pressure, obstacle_detected, light_intensity) VALUES ('$timestamp', '$temperature', '$humidity', '$pressure', '$obstacle_detected', '$light_intensity')";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error adding weather data: " . $this->conn->error);
        }
        return true;
    }

    public function updateWeatherData($timestamp, $new_timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity) {
        // Check if the record exists
        $check_sql = "SELECT COUNT(*) AS count FROM weather_data WHERE timestamp = '$timestamp'";
        $check_result = $this->conn->query($check_sql);
        if (!$check_result) {
            die("Error checking weather data: " . $this->conn->error);
        }
        $count = $check_result->fetch_assoc()['count'];
        if ($count == 0) {
            // Record does not exist, return false
            return false;
        }

        // Record exists, update it
        $sql = "UPDATE weather_data SET timestamp='$new_timestamp', temperature='$temperature', humidity='$humidity', pressure='$pressure', obstacle_detected='$obstacle_detected', light_intensity='$light_intensity' WHERE timestamp='$timestamp'";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error updating weather data: " . $this->conn->error);
        }
        return true;
    }

    public function deleteAllWeatherData() {
        $sql = "DELETE FROM weather_data";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error deleting weather data: " . $this->conn->error);
        }
    }





    public function close() {
        $this->conn->close();
    }
}



?>
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
        $sql = "DELETE FROM weather_data WHERE id='$id'";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error deleting weather data: " . $this->conn->error);
        }
    }

    public function addWeatherData($timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity) {
        $sql = "INSERT INTO weather_data (timestamp, temperature, humidity, pressure, obstacle_detected, light_intensity) VALUES ('$timestamp', '$temperature', '$humidity', '$pressure', '$obstacle_detected', '$light_intensity')";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error adding weather data: " . $this->conn->error);
        }
    }

    public function updateWeatherData($id, $timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity) {
        $sql = "UPDATE weather_data SET timestamp='$timestamp', temperature='$temperature', humidity='$humidity', pressure='$pressure', obstacle_detected='$obstacle_detected', light_intensity='$light_intensity' WHERE id='$id'";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error updating weather data: " . $this->conn->error);
        }
    }



    public function close() {
        $this->conn->close();
    }
}



?>
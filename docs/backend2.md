# Backend Structure

This page of the documentation describes the structure of the backend and how the required assignments
were implemented.

## PSR-12 Coding Style

PSR12 stands for "PHP Standard Recommendation 12", which is a coding style guide for writing clean and consistent PHP code. It is a continuation of PSR1 and PSR2, and it aims to further improve the interoperability and readability of PHP code across different projects and developers.

Some of the key rules of the coding style are:

- Use 4 spaces for indentation, not tabs
- Code should be organized into namespaces and classes
- Function and method signatures should be properly formatted
- Control structures should have braces on the same line
- Comments should be used to explain the code

Since I have to show my whole code for this part, you can see the full code of the backend at the bottom of the page.

## Code comments 

All the important parts of my program code are commented to explain what they do. This is a good practice to make your code more readable and understandable.

Here is an example of my comments:

```php
//Set the credentials
$host = "mariadb";
$user = "root";
$pass = "7YKyE8R2AhKzswfN";
$dbname = "WS";


//set up the database connection
$db = new Database($host, $user, $pass, $dbname);
```

## Preventing Code duplication

The code above is also really important to prevent code duplication. Code duplication is when you have the same code in multiple places. This is bad because it makes your code harder to maintain and understand. It also makes it more likely that you will make a mistake when you change the code.
For bigger projects it makes it harder to find the right parts of the code and writing tests for the code is also harder, since you have to write multiple tests for the same code. If you separate code in multiple functions and classes, you can reuse the code in multiple places. This is called DRY (Don't Repeat Yourself).

In the code snippet above I am creating an object of the Database class, this allows me to interact with Mariadb. It handles all the important functionality like connecting to the database, executing queries and fetching results. This is a good example of code reuse, because I can use the same code in multiple places.

Here are some examples for the functions the Database class provides:



```php
public function __construct($servername, $username, $password, $dbname)
    {   
    // Constructor for the Database class, it sets the credentials and connects to the database
        $this->servername = $servername;
        $this->username = $username;
        $this->password = $password;
        $this->dbname = $dbname;
        $this->connect();
    }
```
Constructor for the Database class, it sets the credentials and connects to the database


```php
    public function close()
    {
        //closes the connection to the database
        $this->conn->close();
    }
```
Closes the connection to the database


```php
    public function query($sql)
    {
    //Executes a query on the database and returns the result as a JSON string

        $result = $this->conn->query($sql);
        $data = array();
        while ($row = $result->fetch_assoc()) {
            $data[] = $row;
        }
        return json_encode($data);
    }
```
Executes a query on the database and returns the result as a JSON string

## Crud operations

The backend provides CRUD (Create, Read, Update, Delete) operations for managing weather data. This means that the php backend
can create new weather data points, delete them, update them and retrieve them. This is done by making HTTP requests. You can do Post, Put, Delete and Get requests to the backend. The backend will then execute the corresponding function. 

Here are the corresponding functions for the CRUD operations:

Read Data:

```php
    public function getWeatherData()
    {
        $sql = "SELECT * FROM weather_data";
        $result = $this->conn->query($sql);
        $data = array();
        while ($row = $result->fetch_assoc()) {
            $data[] = $row;
        }


        return json_encode($data);
    }

    public function query($sql)
    {
        //Executes a query on the database and returns the result as a JSON string
        $result = $this->conn->query($sql);
        $data = array();
        while ($row = $result->fetch_assoc()) {
            $data[] = $row;
        }
        return json_encode($data);
    }
```

Delete Data:

```php
    public function deleteWeatherData($id)
    {
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

    public function deleteAllWeatherData()
    {
        $sql = "DELETE FROM weather_data";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error deleting weather data: " . $this->conn->error);
        }
    }
```

Create Data:
```php
    public function addWeatherData($timestamp, $temperature, $humidity, $pressure, $obstacle_detected, $light_intensity)
    {
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
```

Update Data:
```php
    public function updateWeatherData(
        $timestamp,
        $new_timestamp,
        $temperature,
        $humidity,
        $pressure,
        $obstacle_detected,
        $light_intensity
    ) {
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

        // Record exists, update it then
        $sql = "UPDATE weather_data SET timestamp='$new_timestamp', temperature='$temperature', humidity='$humidity', pressure='$pressure', obstacle_detected='$obstacle_detected', light_intensity='$light_intensity' WHERE timestamp='$timestamp'";
        $result = $this->conn->query($sql);
        if (!$result) {
            die("Error updating weather data: " . $this->conn->error);
        }
        return true;
    }
```

Not all of these functions are necessary for the the basic CRUD operations, but they are useful for the development process. For example the deleteAllWeatherData function is not necessary for the basic CRUD operations, but it is useful for testing purposes.


## The REST API entrance of the Embedded Device 

Only via different request types (GET, POST, PUT, DELETE) the backend can be accessed. The backend is accessed with the following URL:




## Full code

To see the details of my backend the whole code is also displayed here:

    ```php


    ```
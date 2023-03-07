# Backend documentation

This API provides CRUD (Create, Read, Update, Delete) operations for managing weather data.


## Methods:

- POST /?action=weather_data - Creates a new weather data point
- PUT /?action=weather_data - Updates an existing weather data point
- DELETE /?action=weather_data - Deletes an existing weather data point
- GET /?action=weather_data - Retrieves all the weather data points

## Request Parameters:

timestamp: The timestamp of the weather data point (required)
new_timestamp: The new timestamp for an updated weather data point
temperature: The temperature of the weather data point (required)
humidity: The humidity of the weather data point (required)
pressure: The pressure of the weather data point (required)
obstacle_detected: Whether an obstacle was detected (required)
light_intensity: The light intensity of the weather data point (required)

### __Every request returns a JSON file__
- Success: JSON object with a message
- Failure: JSON object with an error message


# Details




### POST request

To create a new weather data point:


```json
{
    "timestamp": "2022-03-07 12:00:00",
    "temperature": 25,
    "humidity": 70,
    "pressure": 1013,
    "obstacle_detected": false,
    "light_intensity": 100
}
```

Response:
```json
{
    "message": "Weather data added successfully."
}
```





### GET request

To retrieve all the weather data points:


Response:

```json
[
  {
    "timestamp": "2022-12-09",
    "temperature": "69420.3",
    "humidity": "45.1",
    "pressure": "1013.2",
    "obstacle_detected": "1",
    "light_intensity": "800.2"
  },
  {
    "timestamp": "2022-12-10",
    "temperature": "69420.3",
    "humidity": "45.1",
    "pressure": "1013.2",
    "obstacle_detected": "1",
    "light_intensity": "800.2"
  }
]
```

### PUT request
To update an existing weather data point:

Request:

```json
{
"timestamp": "2022-03-07 12:00:00",
"new_timestamp": "2022-03-07 13:00:00",
"temperature": 25,
"humidity": 70,
"pressure": 1013,
"obstacle_detected": false,
"light_intensity": 100
}
```

Response:

```json
{
"message": "Weather data updated successfully."
}
```

### DELETE request

To delete an existing weather data point:

Request:

```json
{
"timestamp": "2022-03-07 12:00:00"
}
```

Response:

```json
{
"message": "Weather data deleted successfully."
}
```


# System Requirements
(Reviewed by Rafail)

## EMBRQ#01
After measuring the data the sensors provide, the data gets send to the server with the "send_data(temperature, humidity, 0, 0, light);"
function. The server then stores the data in a database. The data is then displayed on the website.

POST request:

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

Not all the data is actually collected by the wemos, but it allows for future expansion.

Code:
```cpp
void send_data(float temperature, float humidity, float pressure, int obstacle_detected, float light_intensity) {
  // create the JSON payload string
  String payload = "{\"timestamp\":\"current_time,\"temperature\":" + String(temperature) + ",\"humidity\":" + String(humidity) + ",\"pressure\":" + String(pressure) + 
  ",\"obstacle_detected\":" + String(obstacle_detected) + ",\"light_intensity\":" + String(light_intensity) + "}";
  
  // specify the target URL
  httpClient.begin(client, "https://abc3.loca.lt/index.php?action=weather_data");
  
  // set the content type header to JSON
  httpClient.addHeader("Content-Type", "application/json");
  
  // send the POST request with the JSON payload
  int httpCode = httpClient.POST(payload);
  
  // check if the request was successful
  if (httpCode == HTTP_CODE_OK) {
    Serial.println("Data sent successfully!");
  } else {
    Serial.println("Error sending data!");
  }
  
  // free resources
  httpClient.end();
}
```


## EMBRQ#02
The wemos is able to do a get request to do backend server to ask for a config file. The config file looks currently like this:
```json
{
  "timeout": 5000,
  "retry_count": 3,
  "debug_mode": true
}
```

Code:
```cpp
void get_config(){
  httpClient.begin(client, "https://abc3.loca.lt/index.php?action=config"); // Set up HTTP GET request to server
  int httpCode = httpClient.GET(); // Send HTTP GET request and store response code

  if(httpCode == HTTP_CODE_OK) { // Check if response is successful
    String payload = httpClient.getString(); // Get response payload
    Serial.println(payload); // Print payload to serial monitor
  } else {
    Serial.println("Unable to connect :("); // Print error message if response is unsuccessful
  }

}
```

## EMBRQ#03
The wemos uses the humidity and temperature sensor to measure the temperature and humidity. The light sensor is used to measure the light intensity.
A button is used to turn the Weather Wizard on and off.

Humidity/Temperature and light intensity sensor:
```cpp
void read_sensors(){
  // Humidity is measured
  humidity = dht.readHumidity();
  // temperature is measured
  temperature = dht.readTemperature();
  // light
  light = analogRead(light_intensity_sensor); 
   
  // Checking if the measurements have passed without errors
  // if an error is detected, a error message is displayed here
  if (isnan(humidity) || isnan(temperature)) {
    Serial.println("Error reading the sensor");
    return; /quits the function
  }
}
```

Button:
```cpp
void button_input(){
  int buttonState = digitalRead(buttonPin);

  // If the button is pressed, set running_ww to true
  if (buttonState == HIGH) {
    running_ww = true;
  } else {
    running_ww = false;
  }
}
```



## EMBRQ#04
The Weather Wizard uses the LCD screen to display the temperature, humidity. Furthermore, it has two LEDs, one which is green and one which is red. The green LED is used to indicate that the Weather Wizard is on.
The red LED is used to indicate that the Weather Wizard is off.

Function to display the temperature and humidity on the LCD screen:
```cpp
void run_lcd(){
  lcd.setCursor(0, 0);
  lcd.setBacklight(HIGH);
  // print message
  lcd.print("Temp: ");
  
  char temp_c[5];
  lcd.print(dtostrf(temperature, 3, 2, temp_c));
  
  // set cursor to first column, second row
  lcd.setCursor(0,1);
  lcd.print("Humidity: ");
  
  char hum_c[5];
  lcd.print(dtostrf(humidity, 3, 2, hum_c));
  lcd.print("%");
}
```

Main loop of the Weather Wizard to turn the LEDs on and off:

If Weather Wizard is running the green LED is turned on and the red LED is turned off.
If the Weather Wizard is not running the green LED is turned off and the red LED is turned on.

```cpp
void loop(){
  //check if the button was pressed to turn on Weather Wizard
  button_input();

  if(running_ww){
    //turns on the green LED if the Weather Wizard is running
    digitalWrite(led_red, LOW);
    digitalWrite(led_green, HIGH);

  //after doing a get request the backend sends config file 
  get_config(); // Function to get configuration file from server
  //send data with a post request
  send_data(temperature, humidity, 0, 0, light); // Function to send sensor data to server
    
  run_lcd(); // Function to update LCD screen with sensor data
  read_sensors(); // Function to read sensor values

  //waits 2 sec till it makes new measurements
  delay(2000); // Delay before taking new measurements
  lcd.clear(); // Clear LCD screen
  //lcd clears the screen again
  }
  else{
    //turn on the red light and turn off the green light
    digitalWrite(led_red, HIGH);
    digitalWrite(led_green, LOW);
  }
}
```

## EMBRQ#05
This function sets up the Wi-Fi manager, which allows the user to connect to the Weather Wizard via a web interface. The user can then enter the SSID and password of the Wi-Fi network they want to connect to.
Code:
```cpp
void setup_wifi_manager() {
  WiFiManager wifiManager; // Create a Wi-Fi manager object
  bool con = wifiManager.autoConnect("WeatherWizard AutoConnect", "password"); // Attempt to connect to Wi-Fi using saved credentials or create a new AP with the specified SSID and password
  pinMode(LED_BUILTIN, OUTPUT); // Set the built-in LED pin to output mode

  if(!con) {
      Serial.println("Failed to connect."); // Print error message if unable to connect to Wi-Fi
      digitalWrite(LED_BUILTIN, HIGH); // Turn on LED to indicate error
  } 
  else {
      Serial.println("Connected..."); // Print success message if connected to Wi-Fi
      Serial.println(WiFi.localIP()); // Print the IP address of the device
      digitalWrite(LED_BUILTIN, LOW); // Turn off LED to indicate successful connection
  }
}
```

## Full code 

Go to:

```
./emb/src/main/main.ino
```



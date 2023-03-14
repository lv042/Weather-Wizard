/* Author: Luca von Kannen
 * License: Apache License, version 2.0.
 * Goal: This file contains code to operate the WeatherWizard and to connect it to the backend.
 */

#include <LiquidCrystal_I2C.h> // Library for interfacing with the LCD screen
#include <ESP8266WiFi.h> // Library for connecting to WiFi
#include <ESP8266HTTPClient.h> // Library for making HTTP requests
#include <WiFiClient.h>
#include <WiFiManager.h>
// Adafruit_DHT library for interfacing with the DHT11 temperature and humidity sensor
#include "DHT.h"


// DHT input pin
#define DHTPIN 0   
// The DHT-11 is initialized
#define DHTTYPE DHT11 // DHT 11
DHT dht(DHTPIN, DHTTYPE);
 
//parameter for my home wifi
const char* ssid     = "Hi"; // Wi-Fi network name
const char* password = "prinsengracht225d"; // Wi-Fi network password
 
// set the LCD number of columns and rows
int lcdColumns = 16;
int lcdRows = 2;

int buttonPin = D3; 
float humidity = 0; //variable to store the humidity
float temperature = 0; //variable to store the temperature
int light = 0; // variable to store light intensity value

int led_green = D0;
int led_red = D5; // GPIO pin for the LED
int light_intensity_sensor = A0; // GPIO pin for the light intensity sensor

bool running_ww = false; //if the Weather Wizard is currently running
String current_time = "";


// create the HTTP client object
WiFiManager wm; // Object for managing Wi-Fi connection
LiquidCrystal_I2C lcd(0x27, lcdColumns, lcdRows); // Object for interfacing with the LCD screen 
WiFiClient client; // Object for managing Wi-Fi client connection
HTTPClient httpClient; // Object for making HTTP requests

void setup() {
  //sets baud rate to 115200
  Serial.begin(115200);
  
  //wifi
  setup_wifi_manager(); // Function to set up Wi-Fi manager
  //config for lcd
  setup_lcd(); // Function to set up LCD
  setup_dht(); // Function to set up DHT11 sensor
  
  pinMode(led_red, OUTPUT); // Set LED pin to output mode
  pinMode(buttonPin, INPUT);

  //some sensors ned some time before they start working
  delay(2000); // Delay to allow sensors to stabilize
}

void loop(){


  //check if the button was pressed to turn on Weather Wizard
  button_input();

  if(running_ww){
    //get current time
    get_current_time();

    //turns on the green LED if the Weather Wizard is running
    digitalWrite(led_red, LOW);
    digitalWrite(led_green, HIGH);

  //after doing a get request the backend sends config file 
  get_config(); // Function to get configuration file from server
  //send data with a post request
  send_data(temperature, humidity, 0, 0, light); // Function to send sensor data to server
    
  run_lcd(); // Function to update LCD screen with sensor data
  read_sensors(); // Function to read sensor values


  print_to_console();
  //waits 2 sec till it makes new measurements
  delay(2000); // Delay before taking new measurements
  lcd.clear(); // Clear LCD screen
  
  }
  else{
    //turn on the red light and turn off the green light
    digitalWrite(led_red, HIGH);
    digitalWrite(led_green, LOW);
  }
}

void get_current_time() {
  
  httpClient.begin(client, "http://worldtimeapi.org/api/ip");
  int httpCode = httpClient.GET();
  
  //does a get request to receive the current time from this api
  int httpResponseCode = httpClient.GET();
  if (httpResponseCode == HTTP_CODE_OK) {
    String payload = httpClient.getString();
    int startIndex = payload.indexOf("datetime") + 12;
    int endIndex = startIndex + 19;
    current_time = payload.substring(startIndex, endIndex);
    //assign the global variable current_time to the response
  } else {
    Serial.println("HTTP request failed");
  }
}

void button_input(){
  int buttonState = digitalRead(buttonPin);

  // If the button is pressed, set running_ww to true
  if (buttonState == HIGH) {
    running_ww = true;
  } else {
    running_ww = false;
  }
}

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



void read_sensors(){
  // Humidity is measured
  humidity = dht.readHumidity();
  // temperature is measured
  temperature = dht.readTemperature();
  // light
  light = analogRead(light_intensity_sensor); 
   
  // Checking if the measurements have passed without errors
  // if an error is detected, a error message is displayed here
  if (isnan(humidity) || isnan(temperature) || isnan(light)) {
    Serial.println("Error reading the sensor");
    return; //quits the function
  }
}

void print_to_console(){
    // Output to serial console
  Serial.println("-----------------------------------------------------------");
  Serial.print("Humidity: ");
  Serial.print(humidity);
  Serial.print(" %\t");
  Serial.print("Temperature: ");
  Serial.print(temperature);
  Serial.print(char(186)); //Output <°> symbol
  Serial.println("C ");
  Serial.println("-----------------------------------------------------------");
  Serial.print("Light intensity: ");
	Serial.println( light, DEC);
  Serial.println("-----------------------------------------------------------");
}

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

void setup_lcd(){
  lcd.begin(5, 4);// initialize LCD
  lcd.init();
  // turn on LCD backlight                      
  lcd.backlight();
}

void setup_dht(){
  //humidity and temperture 
  Serial.println("KY-015 test - temperature and humidity test:");
  // Measurement is started
  dht.begin();
}
void connectToWiFi(const char* ssid, const char* password) {
  //this code is only needed to connect to a specific wifi
  //this method won't be called because currently the wifi manager is used
  delay(10);
  Serial.println('\n');

  WiFi.begin(ssid, password);
  Serial.print("Connecting to ");
  Serial.print(ssid);
  Serial.println("...");
  
  int i = 0;
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(++i);
    Serial.print(' ');
  }

  Serial.println('\n');
  Serial.println("Connection established!");
  Serial.print("IP address:\t");
  Serial.println(WiFi.localIP());
}

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
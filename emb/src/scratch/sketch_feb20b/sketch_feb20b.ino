/* Author: Luca von Kannen
 * License: MIT
 * Goal: This file contains code to operate the WeatherWizard and to connect it to the backend.
 */

#include <LiquidCrystal_I2C.h>
#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>
#include <WiFiManager.h>


// Adafruit_DHT library 
#include "DHT.h"
// Here the respective input pin can be declared
#define DHTPIN 0   
// The sensor is initialized
#define DHTTYPE DHT11 // DHT 11
DHT dht(DHTPIN, DHTTYPE);
 

 



// set the LCD number of columns and rows
int lcdColumns = 16;
int lcdRows = 2;
float humidity = 0;
float t = 0;
int led = 15;
int sensorPin = A0; 
int value = 0;

//Initilisation
WiFiManager wm;
LiquidCrystal_I2C lcd(0x27, lcdColumns, lcdRows);  



void get_config(){
  WiFiClient client;
  HTTPClient httpClient;
  
  httpClient.begin(client, "https://abc3.loca.lt/index.php?action=config");
  int httpCode = httpClient.GET();

  if(httpCode == HTTP_CODE_OK) { // HTTP_CODE_OK == 200
    String payload = httpClient.getString();
    Serial.println(payload);
  } else {
    Serial.println("Unable to connect :(");
  }

  delay(5000);
}


void loop(){
    get_config();

    digitalWrite(led, HIGH);
    

  // set cursor to first column, first row
  run_lcd();

  lcd.setCursor(0, 0);
  lcd.setBacklight(HIGH);
  // print message
  lcd.print("Temp: ");
  


  char temp_c[5];
  lcd.print(dtostrf(t, 3, 2, temp_c));
  //delay(1000);
  // clears the display to print new message
  //lcd.clear();
  // set cursor to first column, second row
  lcd.setCursor(0,1);
  lcd.print("Humidity: ");

  char hum_c[5];
  lcd.print(dtostrf(humidity, 3, 2, hum_c));
  lcd.print("%");
  delay(1000);

   
  // Two seconds pause between measurements
  delay(2000);
 
  // Humidity is measured
  humidity = dht.readHumidity();
  // temperature is measured
  t = dht.readTemperature();
  
   
  // Checking if the measurements have passed without errors
  // if an error is detected, a error message is displayed here
  if (isnan(humidity) || isnan(t)) {
    Serial.println("Error reading the sensor");
    return;
  }
 
  // Output to serial console
  Serial.println("-----------------------------------------------------------");
  Serial.print("Humidity: ");
  Serial.print(humidity);
  Serial.print(" %\t");
  Serial.print("Temperature: ");
  Serial.print(t);
  Serial.print(char(186)); //Output <°> symbol
  Serial.println("C ");
  Serial.println("-----------------------------------------------------------");
  value = analogRead(sensorPin); 
  Serial.print("Light intensity: ");
	Serial.println( value, DEC);
  Serial.println("-----------------------------------------------------------");

  lcd.clear(); 
}

run_lcd(){




}


//wifi normal
#include <ESP8266WiFi.h>        // Include the Wi-Fi library

const char* ssid     = "Hi";         // The SSID (name) of the Wi-Fi network you want to connect to
const char* password = "prinsengracht225d";     // The password of the Wi-Fi network

void setup() {
  //sets baud rate to 115200
  Serial.begin(115200);
  
  //wifi
  setup_wifi_manager();
  //config for lcd
  setup_lcd();
  
  
  pinMode(led, OUTPUT);

  //some sensors ned some time before they start working
  delay(2000);
}

void setup_lcd(){
  lcd.begin(5, 4);// initialize LCD
  lcd.init();
  // turn on LCD backlight                      
  lcd.backlight();
}

void setup_dht(){
  setup_dht();
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
  WiFiManager wifiManager;
  bool con = wifiManager.autoConnect("WeatherWizard AutoConnect", "password");
  pinMode(LED_BUILTIN, OUTPUT);

  if(!con) {
      Serial.println("Failed to connect.");
      digitalWrite(LED_BUILTIN, HIGH);
  } 
  else {
      Serial.println("Connected...");
      Serial.println(WiFi.localIP());
      digitalWrite(LED_BUILTIN, LOW);
  }

  while (WiFi.status() != WL_CONNECTED) {
    delay(1000); // Waiting on connection...
  }
}
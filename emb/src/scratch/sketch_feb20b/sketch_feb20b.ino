#include <LiquidCrystal_I2C.h>

// Adafruit_DHT library is inserted
#include "DHT.h"
 
// Here the respective input pin can be declared
#define DHTPIN 0   
 
// The sensor is initialized
#define DHTTYPE DHT11 // DHT 11
DHT dht(DHTPIN, DHTTYPE);
 

 



// set the LCD number of columns and rows
int lcdColumns = 16;
int lcdRows = 2;
float h = 0;
float t = 0;
int led = 15;
int sensorPin = A0; 
int value = 0; 



// set LCD address, number of columns and rows
// if you don't know your display address, run an I2C scanner sketch
LiquidCrystal_I2C lcd(0x27, lcdColumns, lcdRows);  

void setup(){
  lcd.begin(5, 4);// initialize LCD
  lcd.init();
  // turn on LCD backlight                      
  lcd.backlight();

  Serial.begin(9600);

  //humidity and temperture -> dht
  Serial.println("KY-015 test - temperature and humidity test:");
  // Measurement is started
  dht.begin();
  pinMode(led, OUTPUT);

}

void loop(){
    digitalWrite(led, HIGH);
    

  // set cursor to first column, first row
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
  lcd.print(dtostrf(h, 3, 2, hum_c));
  lcd.print("%");
  delay(1000);

   
  // Two seconds pause between measurements
  delay(2000);
 
  // Humidity is measured
  h = dht.readHumidity();
  // temperature is measured
  t = dht.readTemperature();
  
   
  // Checking if the measurements have passed without errors
  // if an error is detected, a error message is displayed here
  if (isnan(h) || isnan(t)) {
    Serial.println("Error reading the sensor");
    return;
  }
 
  // Output to serial console
  Serial.println("-----------------------------------------------------------");
  Serial.print("Humidity: ");
  Serial.print(h);
  Serial.print(" %\t");
  Serial.print("Temperature: ");
  Serial.print(t);
  Serial.print(char(186)); //Output <°> symbol
  Serial.println("C ");
  Serial.println("-----------------------------------------------------------");
  value = analogRead(sensorPin); 
	Serial.println(value, DEC);
  Serial.println("-----------------------------------------------------------");

  lcd.clear(); 

}
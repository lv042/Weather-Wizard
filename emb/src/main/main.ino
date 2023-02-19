/*
  Blink

  Turns an led_red on for one second, then off for one second, repeatedly.

  Most Arduinos have an on-board led_red you can control. On the UNO, MEGA and ZERO
  it is attached to digital pin 13, on MKR1000 on pin 6. led_red_BUILTIN is set to
  the correct led_red pin independent of which board is used.
  If you want to know what pin the on-board led_red is connected to on your Arduino
  model, check the Technical Specs of your board at:
  https://www.arduino.cc/en/Main/Products

  modified 8 May 2014
  by Scott Fitzgerald
  modified 2 Sep 2016
  by Arturo Guadalupi
  modified 8 Sep 2016
  by Colby Newman

  This example code is in the public domain.

  https://www.arduino.cc/en/Tutorial/BuiltInExamples/Blink
*/

// the setup function runs once when you press reset or power the board

int speaker = 13;
int led_red = 12;
int led_green = 11;

void setup() {
  pinMode(led_red, OUTPUT); // set pin D4 as output
  pinMode(led_green, OUTPUT);
  pinMode(speaker, OUTPUT); // set pin D4 as output
}

void loop() {
  displayChargedBattery();
  easteregg();

  delay(1000000); //so it only runs once
}



void displayChargedBattery(){
  digitalWrite(led_green, HIGH); // turn on the led_red  
  
  delay(1000); // wait for 1 second
  digitalWrite(led_green, LOW); // turn off the led_red
}


void displayLowBattery(){
  digitalWrite(led_red, HIGH); // turn on the led_red
  digitalWrite(speaker, HIGH); // turn on the speaker
  
  
  delay(1000); // wait for 1 second

  digitalWrite(led_red, LOW); // turn off the led_red
  digitalWrite(speaker, LOW); // turn off the speaker
}




void easteregg(){
   tone(speaker, 440, 500); // A4
  delay(500);
  tone(speaker, 523, 500); // C5
  delay(500);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 494, 1000); // B4
  delay(1000);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 440, 500); // A4
  delay(500);
  tone(speaker, 494, 1000); // B4
  delay(1000);
  tone(speaker, 494, 500); // B4
  delay(500);
  tone(speaker, 587, 500); // D5
  delay(500);
  tone(speaker, 740, 500); // F#5
  delay(500);
  tone(speaker, 659, 1000); // E5
  delay(1000);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 740, 500); // F#5
  delay(500);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 523, 500); // C5
  delay(500);
  tone(speaker, 659, 1000); // E5
  delay(1000);
  // Repeat chorus
  tone(speaker, 440, 500); // A4
  delay(500);
  tone(speaker, 523, 500); // C5
  delay(500);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 494, 1000); // B4
  delay(1000);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 440, 500); // A4
  delay(500);
  tone(speaker, 494, 1000); // B4
  delay(1000);
  tone(speaker, 494, 500); // B4
  delay(500);
  tone(speaker, 587, 500); // D5
  delay(500);
  tone(speaker, 740, 500); // F#5
  delay(500);
  tone(speaker, 659, 1000); // E5
  delay(1000);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 740, 500); // F#5
  delay(500);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 523, 500); // C5
  delay(500);
  tone(speaker, 659, 1000); // E5
  delay(1000);
  tone(speaker, 659, 500); // E5
  delay(500);
  tone(speaker, 523, 500); // C5
}

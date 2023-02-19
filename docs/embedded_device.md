# Embedded device

## Hardware overview
### Sensors
- Temperature sensor (KY-001 or KY-013): To measure the ambient temperature.

- Humidity sensor (KY-015): To measure the relative humidity of the air.

- Light sensor (KY-018): To measure the amount of light, which can help determine cloud cover and daylight conditions.

- Obstacle avoidance sensor (KY-032): To detect the presence of physical obstructions, such as rain or snow, which could affect the readings of other sensors.

- Barometric pressure sensor (not included in the list of available sensors): To measure atmospheric pressure, which can help predict weather changes.

### Output devices
- Display (KY-?): To show current temperature.
- LED (KY-?): To indicate the battery level or if the solar panel gets enough energy.

## Connectivity options
- WiFi (ESP???): To connect to the internet and send data to the cloud.
- LoRa (not included in the list of available sensors): To send data to the cloud.


## Energy management
- Solar panel (KY-?): To charge the battery.
- Battery (KY-?): To store energy.

- A charge controller is required to control the charging and discharging of the batteries
  
- The charge controller regulates the flow of current from the solar panel to the battery to ensure that the battery is not overcharged or discharged.

- During the day, the solar module charges the battery via the charge controller.

- At night or when there is little sunlight, the battery supplies power to the weather station via the charge controller.


## Information for myself

Der Wemos Lolin D1 ist ein kleines, kostengünstiges Entwicklungsboard, das mit der Arduino IDE kompatibel ist und für die Entwicklung von Internet of Things (IoT)-Projekten konzipiert wurde. Das Board basiert auf dem ESP8266-Mikrocontroller und bietet eine Reihe von Konnektivitätsoptionen, einschließlich Wi-Fi und Bluetooth.

Hier ist ein kurzer Überblick über einige wichtige Pins auf dem Wemos Lolin D1:

TX/RX: Dies sind die Pins, die für die serielle Kommunikation verwendet werden. TX steht für Senden und RX steht für Empfangen.
SCL/SDA: Dies sind die Pins für die I2C-Kommunikation, ein Zweidraht-Kommunikationsprotokoll, das zum Anschluss mehrerer Geräte an einen einzigen Mikrocontroller verwendet wird.
GND: Dieser Pin ist die Masseverbindung für die Platine und wird zum Abschließen elektrischer Schaltungen verwendet.
VBUS: Dieser Pin wird zur Stromversorgung externer Geräte verwendet, die an die Platine angeschlossen sind.
RST: Dieser Pin wird zum Zurücksetzen der Platine verwendet.
Ruhezustand: Dieser Pin wird verwendet, um die Karte in den Sleep-Modus zu versetzen, einen Energiesparmodus, der zur Verlängerung der Batterielebensdauer nützlich sein kann.
SCK/MISO/MOSI/SS: Diese Pins werden für die SPI-Kommunikation verwendet, ein Vier-Draht-Kommunikationsprotokoll, mit dem mehrere Geräte an einen einzigen Mikrocontroller angeschlossen werden können. SCK steht für den seriellen Takt, MISO steht für Master-Input-Slave-Output, MOSI steht für Master-Output-Slave-Input und SS steht für Slave Select.
3v3: Dieser Pin ist die 3,3-V-Stromversorgung für die Platine und wird zur Versorgung externer Geräte verwendet, die an die Platine angeschlossen sind.

Übersetzt mit www.DeepL.com/Translator (kostenlose Version)




## Wiring diagram

## Bill of Materials

## Links

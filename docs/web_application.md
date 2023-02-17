# Database
## Structure

### Table: weather_data

- __id__ (integer, primary key): unique identifier for each record</br>
- __temperature__ (float): measured temperature in Celsius</br>
- __humidity__ (float): measured relative humidity in percentage</br>
- __light_intensity__ (float): measured light intensity in Lux</br>
- __pressure__ (float): measured barometric pressure</br>
- __obstacle_detected__ (boolean): whether an obstacle was detected by the obstacle avoidance sensor</br>
- __timestamp__ (timestamp): the time when the data was recorded</br>
  
---

### Table: energy_data

- __id__ (integer, primary key): unique identifier for each record
- __battery_level__ (float): current battery level in percentage
- __solar_panel_voltage__ (float): measured solar panel voltage in volts
- __timestamp__ (timestamp): the time when the data was recorded 


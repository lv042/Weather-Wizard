
# __Database__
## __Structure__

### __Table: data__
- __timestamp__ (integer, primary key): unique identifier for each record</br>
- __weather_data_id__ (integer, foreign key): foreign key to weather_data table</br>
- __energy_data_id__ (integer, foreign key): foreign key to energy_data table</br>

### __Table: weather_data__

- __id__ (integer, primary key): unique identifier for each record</br>
- __temperature__ (float): measured temperature in Celsius</br>
- __humidity__ (float): measured relative humidity in percentage</br>
- __light_intensity__ (float): measured light intensity in Lux</br>
- __pressure__ (float): measured barometric pressure</br>
- __obstacle_detected__ (boolean): Can be used to see if the weather station is covered</br>

  
---

### __Table: energy_data__

- __id__ (integer, primary key): unique identifier for each record
- __battery_level__ (float): current battery level in percentage
- __solar_panel_voltage__ (float): measured solar panel voltage in volts

---

### __ER-diagram__

![Screenshot](docs/../images/er_diagram.png)
__This is a 1:1 relationship between data table and weather/energy table. Could also be summarized in only one table, but in my opinion it is better structured this way__</br>

//Because appearntly MySQL Workbench is very buggy and old, I can't change some of the column's datatypes


## __ReBuild script__

```sql

DROP TABLE IF EXISTS `WS`.`data`;
DROP TABLE IF EXISTS `WS`.`energy_data`;
DROP TABLE IF EXISTS `WS`.`weather_data`;

CREATE TABLE IF NOT EXISTS `ws`.`energy_data`(
    `id` INT NOT NULL AUTO_INCREMENT,
    `battery_level` FLOAT NOT NULL,
    `solar_panel_voltage` FLOAT NOT NULL,
    PRIMARY KEY(`id`)
); 

CREATE TABLE IF NOT EXISTS `ws`.`weather_data`(
    `id` INT NOT NULL,
    `temperature` FLOAT NOT NULL,
    `humidity` FLOAT NOT NULL,
    `pressure` FLOAT NOT NULL,
    `obstacle_detected` BOOLEAN NOT NULL,
    `light_intensity` FLOAT NOT NULL,
    PRIMARY KEY(`id`)
);

CREATE TABLE IF NOT EXISTS `ws`.`data`(
    `timestamp` DATE NOT NULL,
    `energy_data_id` INT NOT NULL,
    `weather_data_id` INT NOT NULL,
    PRIMARY KEY(`timestamp`),
    FOREIGN KEY (energy_data_id) REFERENCES `ws`.`energy_data`(id),
    FOREIGN KEY (weather_data_id) REFERENCES `ws`.`weather_data`(id)
);


```

1. First connect to your machine and select the database you want to use. 
Execute the SQL script by typing the following command:

``` bash
source /path/to/your/file/create_tables.sql;
```


2. After executing the script, you can check that the tables were created successfully by typing the following command:

```sql
SHOW TABLES;
```

3. You should see three tables: data, energy_data, and weather_data.

## __Insert script__

To add some data to the tables, you can use the following script:

```sql

-- Insert sample data for energy_data
INSERT INTO ws.energy_data (id, battery_level, solar_panel_voltage)
VALUES (1, 12.5, 3.2), (2, 11.1, 4.0), (3, 13.2, 2.7);

-- Insert sample data for weather_data
INSERT INTO ws.weather_data (id, temperature, humidity, pressure, obstacle_detected, light_intensity)
VALUES (1, 25.5, 65.0, 1013.25, true, 200.0),
       (2, 23.0, 70.0, 1012.50, false, 150.0),
       (3, 26.5, 60.0, 1014.00, true, 220.0);

-- Insert sample data for data
INSERT INTO ws.data (timestamp, energy_data_id, weather_data_id)
VALUES ('2022-02-22 10:00:00', 1, 1),
       ('2022-02-23 12:00:00', 2, 2),
       ('2022-02-24 14:00:00', 3, 3);

```



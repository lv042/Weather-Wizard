
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

DROP TABLE IF EXISTS `data`;
DROP TABLE IF EXISTS `energy_data`;
DROP TABLE IF EXISTS `weather_data`;

CREATE TABLE IF NOT EXISTS `mydb`.`data` (
  `timestamp` DATE NOT NULL,
  `energy_data_id` INT NOT NULL,
  `weather_data_id` INT NOT NULL,
  PRIMARY KEY (`timestamp`, `energy_data_id`, `weather_data_id`),
  INDEX `fk_data_energy_data_idx` (`energy_data_id` ASC) VISIBLE,
  INDEX `fk_data_weather_data1_idx` (`weather_data_id` ASC) VISIBLE,
  CONSTRAINT `fk_data_energy_data`
    FOREIGN KEY (`energy_data_id`)
    REFERENCES `mydb`.`energy_data` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_data_weather_data1`
    FOREIGN KEY (`weather_data_id`)
    REFERENCES `mydb`.`weather_data` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)

CREATE TABLE IF NOT EXISTS `mydb`.`energy_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `battery_level` FLOAT NOT NULL,
  `solar_panel_voltage` FLOAT NOT NULL,
  PRIMARY KEY (`id`))

CREATE TABLE IF NOT EXISTS `mydb`.`weather_data` (
  `id` INT NOT NULL,
  `temperature` FLOAT NOT NULL,
  `humidity` FLOAT NOT NULL,
  `pressure` FLOAT NOT NULL,
  `obstacle_detected` BOOLEAN NOT NULL,
  `light_intensity` FLOAT NOT NULL,
  PRIMARY KEY (`id`))
  
ENGINE = InnoDB


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

```sql

INSERT INTO `mydb`.`energy_data` (`id`, `battery_level`, `solar_panel_voltage`)
VALUES (1, 50.2, 12.5),
       (2, 45.8, 11.9),
       (3, 48.6, 13.2),
       (4, 51.1, 14.5);

INSERT INTO `mydb`.`weather_data` (`id`, `temperature`, `humidity`, `pressure`, `obstacle_detected`, `light_intensity`)
VALUES (1, 24.5, 56.8, 1013.2, 0, 120),
       (2, 23.1, 55.2, 1012.8, 1, 95),
       (3, 25.8, 57.5, 1014.1, 0, 130),
       (4, 22.6, 53.7, 1012.4, 0, 115);

INSERT INTO `mydb`.`data` (`timestamp`, `energy_data_id`, `weather_data_id`)
VALUES ('2023-02-16', 1, 2),
       ('2023-02-16', 3, 1),
       ('2023-02-17', 2, 4),
       ('2023-02-17', 4, 3);

```



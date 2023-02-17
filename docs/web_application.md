//maybe rename the page to backend.md

# Database
## Structure

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
- __obstacle_detected__ (boolean): whether an obstacle was detected by the obstacle avoidance sensor</br>

  
---

### __Table: energy_data__

- __id__ (integer, primary key): unique identifier for each record
- __battery_level__ (float): current battery level in percentage
- __solar_panel_voltage__ (float): measured solar panel voltage in volts

---

### __ER diagram__

![Screenshot](docs/../images/er_diagram.png)
//Because appearntly MySQL Workbench is very buggy and old, I can't change some of the column's datatype
__
## __Rebuild script__

```sql

DROP DATABASE IF EXISTS `mydb`;

CREATE DATABASE IF NOT EXISTS `mydb`;

USE `mydb`;

DROP TABLE IF EXISTS `data`;
DROP TABLE IF EXISTS `energy_data`;
DROP TABLE IF EXISTS `weather_data`;

CREATE TABLE IF NOT EXISTS `mydb`.`data` (
  `timestamp` DATE NOT NULL,
  `energy_data_id` INT NOT NULL,
  `weather_data_id` INT NOT NULL,
  PRIMARY KEY (`timestamp`),
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
    ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS `mydb`.`energy_data` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `battery_level` FLOAT NOT NULL,
  `solar_panel_voltage` FLOAT NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `mydb`.`weather_data` (
  `id` INT NOT NULL,
  `temperature` FLOAT NOT NULL,
  `humidity` FLOAT NOT NULL,
  `pressure` FLOAT NOT NULL,
  `obstacle_detected` TINYINT NOT NULL,
  `light_intensity` FLOAT NULL,
  PRIMARY KEY (`id`, `obstacle_detected`)
);

ENGINE = InnoDB;

```


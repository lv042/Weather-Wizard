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



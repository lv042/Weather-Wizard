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
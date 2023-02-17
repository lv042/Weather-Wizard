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

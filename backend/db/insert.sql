INSERT INTO `ws`.`energy_data` (`timestamp`, `battery_level`, `solar_panel_voltage`)
VALUES ('2022-02-01', 0.85, 14.2),
       ('2022-02-02', 0.87, 12.8),
       ('2022-02-03', 0.89, 11.5),
       ('2022-02-04', 0.90, 10.2);

INSERT INTO `ws`.`weather_data` (`timestamp`, `temperature`, `humidity`, `pressure`, `obstacle_detected`, `light_intensity`)
VALUES ('2022-02-01', 21.5, 65.2, 1015.7, 0, 1024.8),
       ('2022-02-02', 22.1, 61.8, 1016.2, 0, 1025.1),
       ('2022-02-03', 23.0, 57.4, 1016.4, 1, 1025.5),
       ('2022-02-04', 23.5, 54.6, 1016.7, 1, 1025.8);
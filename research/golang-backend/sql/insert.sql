INSERT INTO ws.energy_data (timestamp, battery_level, solar_panel_voltage)
VALUES ('2022-02-01'::date, 0.85, 14.2),
       ('2022-02-02'::date, 0.87, 12.8),
       ('2022-02-03'::date, 0.89, 11.5),
       ('2022-02-04'::date, 0.90, 10.2);

INSERT INTO ws.weather_data (timestamp, temperature, humidity, pressure, obstacle_detected, light_intensity)
VALUES ('2022-02-01'::date, 21.5, 65.2, 1015.7, false, 1024.8),
       ('2022-02-02'::date, 22.1, 61.8, 1016.2, false, 1025.1),
       ('2022-02-03'::date, 23.0, 57.4, 1016.4, true, 1025.5),
       ('2022-02-04'::date, 23.5, 54.6, 1016.7, true, 1025.8);

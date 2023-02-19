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
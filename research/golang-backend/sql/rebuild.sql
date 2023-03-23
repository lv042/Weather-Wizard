CREATE SCHEMA IF NOT EXISTS ws;

DROP TABLE IF EXISTS "WS"."weather_data";
DROP TABLE IF EXISTS "WS"."energy_data";

CREATE TABLE IF NOT EXISTS "ws"."energy_data"(
                                                 "timestamp" DATE NOT NULL,
                                                 "battery_level" FLOAT NOT NULL,
                                                 "solar_panel_voltage" FLOAT NOT NULL,
                                                 PRIMARY KEY("timestamp")
    );

CREATE TABLE IF NOT EXISTS "ws"."weather_data"(
                                                  "timestamp" DATE NOT NULL,
                                                  "temperature" FLOAT NOT NULL,
                                                  "humidity" FLOAT NOT NULL,
                                                  "pressure" FLOAT NOT NULL,
                                                  "obstacle_detected" BOOLEAN NOT NULL,
                                                  "light_intensity" FLOAT NOT NULL,
                                                  PRIMARY KEY("timestamp")
    );
CREATE TABLE IF NOT EXISTS temp_data (
    id SERIAL PRIMARY KEY,
    station_id INTEGER NOT NULL REFERENCES stations(id),
    temperature NUMERIC(10, 2) NOT NULL,
    recorded_time TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE temp_data ADD CONSTRAINT temp_data_station_id_recorded_time_unique UNIQUE (station_id, recorded_time);
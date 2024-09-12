package data

import (
	"context"
	"database/sql"
	"time"
)

type TempDataModel struct {
	DB *sql.DB
}

type TempDataModelInterface interface {
	Insert(data *SensorData) error
	ListByStation(station *Station, filters *SensorFilters) ([]*TempData, error)
}

type TempData struct {
	ID          int64     `json:"id"`
	StationID   int64     `json:"station_id"`
	Temperature float64   `json:"temperature"`
	RecordedAt  time.Time `json:"recorded_at"`
}

func (m TempDataModel) Insert(data *SensorData) error {
	query := `
		INSERT INTO temp_data (station_id, temperature, recorded_time)
		VALUES ($1, $2, $3)`

	args := []interface{}{data.StationID, data.Data, time.Now()}
	_, err := m.DB.Exec(query, args...)
	return err
}

func (m TempDataModel) ListByStation(station *Station, filters *SensorFilters) ([]*TempData, error) {
	query := `
		SELECT id, temperature, recorded_time
		FROM temp_data
		WHERE station_id = $1 AND recorded_time > $2 AND recorded_time < $3
		ORDER BY recorded_time DESC
		LIMIT $4`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{station.ID, filters.TimestampGTE, filters.TimestampLT, filters.Limit}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []*TempData{}

	for rows.Next() {
		var temp TempData
		err := rows.Scan(&temp.ID, &temp.Temperature, &temp.RecordedAt)
		if err != nil {
			return nil, err
		}

		temp.StationID = station.ID
		data = append(data, &temp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

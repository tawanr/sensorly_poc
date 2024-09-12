package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDuplicateStation = errors.New("duplicate station")
)

type Station struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"-"`
	Name      string    `json:"name"`
	Latitude  *float64  `json:"latitude"`
	Longitude *float64  `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StationModel struct {
	DB *sql.DB
}

type StationModelInterface interface {
	Insert(station *Station) error
	GetById(id int64, user *User) (*Station, error)
	ListByUser(user *User) ([]*Station, error)
}

func (m StationModel) Insert(station *Station) error {
	query := `
		INSERT INTO stations (name, latitude, longitude, user_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`

	args := []interface{}{station.Name, station.Latitude, station.Longitude, station.UserID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&station.ID, &station.CreatedAt, &station.UpdatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "stations_name_user_id_unique"`:
			return ErrDuplicateStation
		default:
			return err
		}
	}
	return nil
}

func (m StationModel) GetById(id int64, user *User) (*Station, error) {
	query := `
		SELECT id, name, latitude, longitude, user_id, created_at, updated_at
		FROM stations
		WHERE id = $1 AND user_id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var station Station
	err := m.DB.QueryRowContext(ctx, query, id, user.ID).Scan(
		&station.ID,
		&station.Name,
		&station.Latitude,
		&station.Longitude,
		&station.UserID,
		&station.CreatedAt,
		&station.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &station, nil
}

func (m StationModel) ListByUser(user *User) ([]*Station, error) {
	query := `
		SELECT id, name, latitude, longitude, user_id, created_at, updated_at
		FROM stations
		WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var stations []*Station
	rows, err := m.DB.QueryContext(ctx, query, user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var station Station
		err := rows.Scan(
			&station.ID,
			&station.Name,
			&station.Latitude,
			&station.Longitude,
			&station.UserID,
			&station.CreatedAt,
			&station.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		stations = append(stations, &station)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stations, nil
}

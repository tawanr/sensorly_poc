package data

import (
	"time"

	"github.org/tawanr/sensorly/internal/validator"
)

type SensorFilters struct {
	TimestampGTE time.Time `json:"timestamp_gte"`
	TimestampLT  time.Time `json:"timestamp_lt"`
	Period       string    `json:"period"`
	Limit        int       `json:"limit"`
}

type SensorData struct {
	StationID int64   `json:"station_id"`
	Data      float64 `json:"data"`
}

type SensorDataModelInterface interface {
	Insert(data *SensorData) error
}

func ValidateTimePeriod(v *validator.Validator, period string) {
	v.Check(period != "", "period", "must be specified")
	v.Check(validator.In(period, "hour", "day", "week", "month", "year"), "period", "must be one of: hour, day, week, month, year")
}

func ValidateSensorFilters(v *validator.Validator, filters *SensorFilters) {
	v.Check(filters.TimestampGTE.After(time.Date(1970, 1, 1, 0, 0, 0, 0, time.Local)), "timestamp_gte", "must be greater than or equal to 1970-01-01 00:00:00")
	v.Check(filters.TimestampLT.After(filters.TimestampGTE), "timestamp_lt", "must be greater than timestamp_gte")
	ValidateTimePeriod(v, filters.Period)
}

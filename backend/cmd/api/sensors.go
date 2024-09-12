package main

import (
	"net/http"
	"time"

	"github.org/tawanr/sensorly/internal/data"
	"github.org/tawanr/sensorly/internal/validator"
)

type tempInput struct {
	Temperature float64 `json:"temperature"`
	RecordedAt  int64   `json:"recorded_at"`
}

func (app *application) addSensorDataHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		StationID int64            `json:"station_id"`
		TempData  *data.SensorData `json:"temp_data"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.TempData != nil {
		input.TempData.StationID = input.StationID
		err = app.models.TempData.Insert(input.TempData)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	app.writeJSON(w, http.StatusCreated, map[string]interface{}{"station_id": input.StationID}, nil)
}

func (app *application) listSensorDataHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		StationID int `json:"station_id"`
		data.SensorFilters
	}

	v := validator.New()
	qs := r.URL.Query()
	user := app.contextGetUser(r)

	input.StationID = app.readInt(qs, "station_id", -1, v)
	input.Period = app.readString(qs, "period", "hour")
	input.TimestampGTE = app.readTime(qs, "timestamp_gte", time.Date(1970, 1, 1, 0, 0, 0, 1, time.Local), v)
	input.TimestampLT = app.readTime(qs, "timestamp_lt", time.Now(), v)
	input.Limit = app.readInt(qs, "limit", 100, v)

	if input.StationID < 0 {
		v.AddError("station_id", "must have a value")
	}
	data.ValidateSensorFilters(v, &input.SensorFilters)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	station, err := app.models.Stations.GetById(int64(input.StationID), user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	responses, err := app.models.TempData.ListByStation(station, &input.SensorFilters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": responses}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

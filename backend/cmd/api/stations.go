package main

import (
	"net/http"
	"strconv"

	"github.org/tawanr/sensorly/internal/data"
	"github.org/tawanr/sensorly/internal/validator"
)

func (app *application) getStationHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("stationID"))
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}
	user := app.contextGetUser(r)

	station, err := app.models.Stations.GetById(int64(id), user)
	if err != nil {
		switch err {
		case data.ErrRecordNotFound:
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, station, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listStationsHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	stations, err := app.models.Stations.ListByUser(user)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"stations": stations}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) addStationHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name      string   `json:"name"`
		Latitude  *float64 `json:"latitude"`
		Longitude *float64 `json:"longitude"`
	}

	user := app.contextGetUser(r)
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	v.Check(input.Name != "", "name", "must be specified")

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	station := &data.Station{
		Name:      input.Name,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		UserID:    user.ID,
	}

	err = app.models.Stations.Insert(station)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, station, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

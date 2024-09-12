package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.notFoundHandler)
	mux.HandleFunc("GET /api/v1/health", app.healthcheckHandler)

	mux.HandleFunc("POST /api/v1/users", app.registerUserHandler)
	mux.HandleFunc("POST /api/v1/tokens/authentication", app.createTokenHandler)

	mux.HandleFunc("GET /api/v1/stations", app.requireAuthentication(app.listStationsHandler))
	mux.HandleFunc("GET /api/v1/stations/{stationID}", app.requireAuthentication(app.getStationHandler))
	mux.HandleFunc("POST /api/v1/stations", app.requireAuthentication(app.addStationHandler))
	mux.HandleFunc("GET /api/v1/sensors", app.requireAuthentication(app.listSensorDataHandler))
	mux.HandleFunc("POST /api/v1/sensors", app.requireAuthentication(app.addSensorDataHandler))

	return app.authenticate(mux)
}

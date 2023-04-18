package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	if err := app.writeJSON(w, http.StatusOK, data, nil); err != nil {
		app.logger.Println(err)
		http.Error(w, "Could not write to json", http.StatusInternalServerError)
	}
}

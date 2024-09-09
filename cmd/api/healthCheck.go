package main

import (
	"fmt"
	"net/http"
)

// Harcoded JSON
func (a *applicationDependencies) healthCheckHandler(w http.ResponseWriter,
	r *http.Request) {
	jsResponse := `{"status": "available", "evironment": %q, "version": %q}`
	jsResponse = fmt.Sprintf(jsResponse, a.config.environment, appVersion)
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsResponse))
}

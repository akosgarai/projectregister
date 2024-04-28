package controller

import (
	"encoding/json"
	"net/http"
)

// HealthController is the health controller.
// It is responsible for handling health checks.
// TODO: return only a status code.
func HealthController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

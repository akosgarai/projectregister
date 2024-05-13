package controller

import (
	"net/http"
)

// HealthController is the health controller.
// It is responsible for handling health checks.
func (c *Controller) HealthController(w http.ResponseWriter, r *http.Request) {
	c.renderer.Status(w, http.StatusOK)
}

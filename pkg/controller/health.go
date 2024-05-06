package controller

import (
	"net/http"
)

// HealthController is the health controller.
// It is responsible for handling health checks.
// TODO: return only a status code.
func (c *Controller) HealthController(w http.ResponseWriter, r *http.Request) {
	c.renderer.JSON(w, http.StatusOK, nil)
}

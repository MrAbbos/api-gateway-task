package handlers

import (
	"github.com/gin-gonic/gin"

	"api_gateway_task/api/http"
	"api_gateway_task/config"
)

// GetConfig godoc
// @ID get_config
// @Router /config [GET]
// @Summary get config data on the debug mode
// @Description show service config data when the service environment set to debug mode
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=config.Config} "Response data"
// @Failure 400 {object} http.Response{}
func (h *Handler) GetConfig(c *gin.Context) {

	switch h.cfg.Environment {
	case config.DebugMode:
		h.handleResponse(c, http.OK, h.cfg)
		return
	case config.TestMode:
		h.handleResponse(c, http.OK, h.cfg.Environment)
		return
	case config.ReleaseMode:
		h.handleResponse(c, http.OK, "private data")
		return
	}

	h.handleResponse(c, http.BadEnvironment, "wrong environment value passed")
}

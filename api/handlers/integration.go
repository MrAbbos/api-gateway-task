package handlers

import (
	"api_gateway_task/api/http"
	"api_gateway_task/genproto/integrations_service"

	"github.com/gin-gonic/gin"
)

type empty struct{} 

// Content pulling godoc
// @ID content-pulling
// @Router /contect/pulling [POST]
// @Summary Content pulling
// @Description Content pulling
// @Tags contect
// @Accept json
// @Produce json
// @Param content body integrations_service.ContentPullingRequest true "ContentPullingRequest"
// @Success 201 {object} http.Response{data=empty} "Created"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ContentPulling(c *gin.Context) {

	var requestBody integrations_service.ContentPullingRequest

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		h.handleErrorResponse(c, 400, "parse error", err.Error())
		return
	}

	resp, err := h.services.IntegrationsService().ContentPulling(
		c.Request.Context(),
		&requestBody,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

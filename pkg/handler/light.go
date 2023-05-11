package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLight(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input Smarthouse_server.DeviceLight
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lightId, err := h.services.DeviceLight.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lightId": lightId,
	})
}

type getAllLightsResponse struct {
	Data []Smarthouse_server.DeviceLight `json:"lights"`
}

func (h *Handler) getAllLights(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lights, err := h.services.DeviceLight.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllLightsResponse{
		Data: lights,
	})
}

func (h *Handler) getLightById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	light, err := h.services.DeviceLight.GetById(userId, lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, light)
}

func (h *Handler) updateLight(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	var input Smarthouse_server.UpdateLightInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeviceLight.Update(userId, lightId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteLight(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	err = h.services.DeviceLight.Delete(userId, lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

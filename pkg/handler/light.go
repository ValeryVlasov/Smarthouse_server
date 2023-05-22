package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLight(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	var input Smarthouse_server.DeviceLight
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lightId, err := h.services.DeviceLight.Create(user.Id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lightId": lightId,
	})
}

func (h *Handler) getAllLights(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	lights, err := h.services.DeviceLight.GetAll(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lights)
}

func (h *Handler) getLightById(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	light, err := h.services.DeviceLight.GetById(user.Id, lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, light)
}

func (h *Handler) updateLight(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
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

	if err := h.services.DeviceLight.Update(user.Id, lightId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteLight(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	err = h.services.DeviceLight.Delete(user.Id, lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) toggleLight(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	err = h.services.DeviceLight.Toggle(user.Id, lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

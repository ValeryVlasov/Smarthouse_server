package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

func (h *Handler) createLight(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	var input Smarthouse_server.DeviceLight
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lightId, err := h.services.DeviceLight.Create(cast.ToInt(userId), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"lightId": lightId,
	})
}

func (h *Handler) getAllLights(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	lights, err := h.services.DeviceLight.GetAll(cast.ToInt(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lights)
}

func (h *Handler) getLightById(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	light, err := h.services.DeviceLight.GetById(cast.ToInt(userId), lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, light)
}

func (h *Handler) updateLight(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
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

	if err := h.services.DeviceLight.Update(cast.ToInt(userId), lightId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteLight(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	err = h.services.DeviceLight.Delete(cast.ToInt(userId), lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) toggleLight(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	lightId, err := strconv.Atoi(c.Param("light_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lightId param")
		return
	}

	err = h.services.DeviceLight.Toggle(cast.ToInt(userId), lightId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

func (h *Handler) createDetector(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	var input Smarthouse_server.DeviceDetector
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	detectorId, err := h.services.DeviceDetector.Create(cast.ToInt(userId), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"detectorId": detectorId,
	})
}

func (h *Handler) getAllDetectors(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	detectors, err := h.services.DeviceDetector.GetAll(cast.ToInt(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, detectors)
}

func (h *Handler) getDetectorById(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	detector, err := h.services.DeviceDetector.GetById(cast.ToInt(userId), detectorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, detector)
}

func (h *Handler) updateDetector(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	var input Smarthouse_server.UpdateDetectorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeviceDetector.Update(cast.ToInt(userId), detectorId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteDetector(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	err = h.services.DeviceDetector.Delete(cast.ToInt(userId), detectorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

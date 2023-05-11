package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createDetector(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input Smarthouse_server.DeviceDetector
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	detectorId, err := h.services.DeviceDetector.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"detectorId": detectorId,
	})
}

type getAllDetectorsResponse struct {
	Data []Smarthouse_server.DeviceDetector `json:"detectors"`
}

func (h *Handler) getAllDetectors(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	detectors, err := h.services.DeviceDetector.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDetectorsResponse{
		Data: detectors,
	})
}

func (h *Handler) getDetectorById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	detector, err := h.services.DeviceDetector.GetById(userId, detectorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, detector)
}

func (h *Handler) updateDetector(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
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

	if err := h.services.DeviceDetector.Update(userId, detectorId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteDetector(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	err = h.services.DeviceDetector.Delete(userId, detectorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

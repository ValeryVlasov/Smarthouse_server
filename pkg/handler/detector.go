package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createDetector(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	var input Smarthouse_server.DeviceDetector
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	detectorId, err := h.services.DeviceDetector.Create(user.Id, input)
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
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	detectors, err := h.services.DeviceDetector.GetAll(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDetectorsResponse{
		Data: detectors,
	})
}

func (h *Handler) getDetectorById(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	detector, err := h.services.DeviceDetector.GetById(user.Id, detectorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, detector)
}

func (h *Handler) updateDetector(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
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

	if err := h.services.DeviceDetector.Update(user.Id, detectorId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteDetector(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	detectorId, err := strconv.Atoi(c.Param("detector_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid detectorId param")
		return
	}

	err = h.services.DeviceDetector.Delete(user.Id, detectorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

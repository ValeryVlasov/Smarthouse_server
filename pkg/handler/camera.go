package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCamera(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input Smarthouse_server.DeviceCamera
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cameraId, err := h.services.DeviceCamera.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"cameraId": cameraId,
	})
}

type getAllCamerasResponse struct {
	Data []Smarthouse_server.DeviceCamera `json:"cameras"`
}

func (h *Handler) getAllCameras(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cameras, err := h.services.DeviceCamera.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCamerasResponse{
		Data: cameras,
	})
}

func (h *Handler) getCameraById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	camera, err := h.services.DeviceCamera.GetById(userId, cameraId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, camera)
}

func (h *Handler) updateCamera(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	var input Smarthouse_server.UpdateCameraInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DeviceCamera.Update(userId, cameraId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteCamera(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	err = h.services.DeviceCamera.Delete(userId, cameraId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

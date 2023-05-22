package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCamera(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	var input Smarthouse_server.DeviceCamera
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cameraId, err := h.services.DeviceCamera.Create(user.Id, input)
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
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	cameras, err := h.services.DeviceCamera.GetAll(user.Id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cameras)
}

func (h *Handler) getCameraById(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	camera, err := h.services.DeviceCamera.GetById(user.Id, cameraId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, camera)
}

func (h *Handler) updateCamera(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
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

	if err := h.services.DeviceCamera.Update(user.Id, cameraId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteCamera(c *gin.Context) {
	user, ok := h.GetUser(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect login or password")
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	err = h.services.DeviceCamera.Delete(user.Id, cameraId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

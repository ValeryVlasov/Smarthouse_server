package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

func (h *Handler) createCamera(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	var input Smarthouse_server.DeviceCamera
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	cameraId, err := h.services.DeviceCamera.Create(cast.ToInt(userId), input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"cameraId": cameraId,
	})
}

func (h *Handler) getAllCameras(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}
	cameras, err := h.services.DeviceCamera.GetAll(cast.ToInt(userId))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cameras)
}

func (h *Handler) getCameraById(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	camera, err := h.services.DeviceCamera.GetById(cast.ToInt(userId), cameraId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, camera)
}

func (h *Handler) updateCamera(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
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

	if err := h.services.DeviceCamera.Update(cast.ToInt(userId), cameraId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteCamera(c *gin.Context) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return
	}

	cameraId, err := strconv.Atoi(c.Param("camera_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid cameraId param")
		return
	}

	err = h.services.DeviceCamera.Delete(cast.ToInt(userId), cameraId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

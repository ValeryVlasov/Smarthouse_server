package handler

import (
	"github.com/ValeryVlasov/Smarthouse_server/pkg/service"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	home := router.Group("/home")
	{
		home.GET("", h.home)
	}

	api := router.Group("/api", h.userIdentity)
	{
		devices := api.Group("/devices")
		{
			lights := devices.Group("/lights")
			{
				lights.POST("/", h.createLight)
				lights.GET("/", h.getAllLights)
				lights.GET("/:light_id", h.getLightById)
				lights.PUT("/:light_id", h.updateLight)
				lights.DELETE("/:light_id", h.deleteLight)
				lights.POST("/:light_id/toggle/", h.toggleLight)
			}

			cameras := devices.Group("/cameras")
			{
				cameras.POST("/", h.createCamera)
				cameras.GET("/", h.getAllCameras)
				cameras.GET("/:camera_id", h.getCameraById)
				cameras.PUT("/:camera_id", h.updateCamera)
				cameras.DELETE("/:camera_id", h.deleteCamera)
			}

			detectors := devices.Group("/detectors")
			{
				detectors.POST("/", h.createDetector)
				detectors.GET("/", h.getAllDetectors)
				detectors.GET("/:detector_id", h.getDetectorById)
				detectors.PUT("/:detector_id", h.updateDetector)
				detectors.DELETE("/:detector_id", h.deleteDetector)
			}
		}
	}

	return router
}

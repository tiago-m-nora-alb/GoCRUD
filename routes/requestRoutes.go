package routes

import (
	"github.com/TiagoNora/GoCRUDV2/controller"
	"github.com/gin-gonic/gin"
)

func RequestRoutes(engine *gin.Engine) {
	requestController := controller.NewRequestController()
	engine.POST("/request/test", requestController.MakeRequest)
}
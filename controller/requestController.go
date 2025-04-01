package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RequestController interface {
	MakeRequest(c *gin.Context)
}

type requestController struct {
}

type message struct {
	Info string `json:"message"`
}


func (r *requestController) MakeRequest(c *gin.Context) {
	log.Info().Msg("Request to a endpoint")
	result, err := http.Get("http://localhost:8000/health")
	if err != nil {
		log.Error().Msgf("Error making request: %v", err)
		sendError(c, http.StatusInternalServerError, "Error making request")
	}

	decoder := json.NewDecoder(result.Body)
    var m message
	err = decoder.Decode(&m)
    if err != nil {
        panic(err)
    }
	
	sendSuccess(c, "Request to a endpoint", m)
}

func NewRequestController() RequestController {
	return &requestController{}
}
package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)



func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", h.homepageHandler)
	return router
}

func (h *Handler) homepageHandler(c *gin.Context) {
	log.Print("homepage kitchen was GETTED")
}

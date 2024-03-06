package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/", homepageHandler)
	return router
}
func homepageHandler(c *gin.Context) {
	log.Print("homepage kitchen was GETTED")
}

package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// UUID
	server.GET("/uuid", getUUID)
}
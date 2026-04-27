package routes

import (
	"api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use(middlewares.Authenticate)
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events/:eventId", updateEvent)
	authenticatedGroup.DELETE("/events/:eventId", deleteEvent)
	authenticatedGroup.POST("/events/:eventId/register", registerForEvent)
	authenticatedGroup.DELETE("/events/:eventId/register", cancelRegistration)

	server.GET("/events", getEvents)
	server.GET("/events/:eventId", getEvent)

}

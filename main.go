package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/project-sa-g03/config"
	"example.com/project-sa-g03/controller"
)

const PORT = "8000"

func main() {

	// open connection database

	config.ConnectionDB()

	// Generate databases

	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Auth Route
	r.PATCH("/locks", controller.ClearStatus)

	r.POST("/locks", controller.CreateLock)

	r.GET("/locks/:id", controller.GetLockById)

	r.PUT("/locks/:id", controller.UpdateLock)

	r.PUT("/locks/:id/update", controller.UpdateLockStatus)

	r.POST("/clear-status", controller.ClearStatus)

	r.GET("/api/count-shops", controller.CountShops)

	r.GET("/api/count-users", controller.CountUsers)

	r.GET("/api/sum-reservations", controller.CountReservations)

	r.GET("/api/dashboard", controller.GetDashboardData)

	r.DELETE("/locks/:id", controller.DeleteLock)

	r.GET("/locks", controller.GetLocks)

	r.GET("/", func(c *gin.Context) {

		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)

	})

	// Run the server

	r.Run("localhost:" + PORT)

}

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}

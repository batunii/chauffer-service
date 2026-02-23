package routes

import (
    "chauffer-service/handlers"
    "chauffer-service/middleware"
    "github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
    r.Use(middleware.Logger())

    // API v1 group
    v1 := r.Group("/api/v1")
    {
        v1.GET("/checkHealth", handlers.CheckHealth)
    }
}

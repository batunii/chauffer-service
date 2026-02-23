package main

import (
    "chauffer-service/config"
    "chauffer-service/routes"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    cfg := config.Load()

    r := gin.Default()
    routes.RegisterRoutes(r)

    log.Printf("Service starting on port %s", cfg.Port)
    if err := r.Run(":" + cfg.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}

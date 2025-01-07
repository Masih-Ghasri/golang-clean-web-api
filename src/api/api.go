package main

import (
  "fmt"
  "api/routers"
  "api/config"
  "github.com/gin-gonic/gin"
  
)

func InitServer() {
  cfg := config.GetConfig()  

  r := gin.Default()
  r.Use(gin.Logger(), gin.Recovery())

  api := r.Group("/api")

  v1 := api.Group("/v1")
  {
    health := v1.Group("/health")
    routers.Health(health)
  }

  v2 := api.Group("/v2")
  {
    test_router := v2.Group("/test")
    routers.TestRouter(test_router)
  }

  r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
}

func main(){
  InitServer()
}
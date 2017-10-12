package main

import (
  "github.com/gin-gonic/gin"

  "routes"
)

func main() {
  engine := gin.Default()
  routes.Init(engine)
  engine.Run(":8080")
}

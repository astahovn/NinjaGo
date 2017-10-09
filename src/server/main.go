package main

import (
  "github.com/gin-gonic/gin"

  "routes"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("tpl/**/*")
  router.GET("/", routes.Index)
  router.GET("/about", routes.About)
  router.Run(":8080")
}

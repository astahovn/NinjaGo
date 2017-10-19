package routes

import (
  "github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
  // init templates
  engine.LoadHTMLGlob("tpl/**/*")

  // init static
  engine.Static("/images", "./assets/images")
  engine.Static("/css", "./assets/css")

  // init dynamic
  engine.GET("/", Index)
  engine.GET("/login", Login)
  engine.POST("/login", LoginPost)
}

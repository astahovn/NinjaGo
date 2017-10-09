package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Index(c *gin.Context) {
  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "title": "Index",
  })
}

func About(c *gin.Context) {
  c.HTML(http.StatusOK, "index/about.tmpl", gin.H{
    "title": "About",
  })
}

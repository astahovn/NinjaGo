package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "models/post"
)

func Index(c *gin.Context) {
  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "title": "Index",
    "posts": post.FetchLast(),
  })
}

func Login(c *gin.Context) {
  c.HTML(http.StatusOK, "index/login.tmpl", gin.H{
  })
}

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

func LoginPost(c *gin.Context) {
  login := c.PostForm("login")
  password := c.PostForm("password")
  if login == "habr" && password == "habr" {
    c.Redirect(http.StatusFound, "/")
    return
  }
  c.HTML(http.StatusOK, "index/login.tmpl", gin.H{
    "login": login,
    "error": true,
  })
}

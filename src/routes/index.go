package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "models/post"
  "lib/session"
)

func Index(c *gin.Context) {
  session.CheckSession(c)
  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "title": "Index",
    "posts": post.FetchLast(),
    "auth": session.GetAuth(),
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
    session.Init(c, 1, c.Request.UserAgent())
    c.Redirect(http.StatusFound, "/")
    return
  }
  c.HTML(http.StatusOK, "index/login.tmpl", gin.H{
    "login": login,
    "error": true,
  })
}

func Logout(c *gin.Context) {
  session.CheckSession(c)
  session.Close(c)
  c.Redirect(http.StatusFound, "/")
}

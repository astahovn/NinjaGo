package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/astahovn/ninja/models/user"
  "github.com/astahovn/ninja/lib/session"
)

// Index page
func Index(c *gin.Context) {
  if session.GetAuth().UserId > 0 {
    c.Redirect(http.StatusFound, "/profile")
    return
  }

  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "title": "Index",
    "auth": session.GetAuth().Nick,
  })
}

// Registration form
func Register(c *gin.Context) {
  c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
  })
}

// Registration request
func RegisterPost(c *gin.Context) {
  login := c.PostForm("login")
  //password := c.PostForm("password")

  newUser := user.User{Username: login}
  if Error := user.Register(newUser); Error != "" {
    c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
      "login": login,
      "errors": Error,
    })
    return
  }

  c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
    "login": login,
    "success": true,
    "errors": false,
  })
}

// Login request
func LoginPost(c *gin.Context) {
  login := c.PostForm("login")
  //password := c.PostForm("password")
  tmpUser := user.LoadByUsername(login)
  if tmpUser.ID > 0 {
    session.Init(c, tmpUser.ID, c.Request.UserAgent())
    c.Redirect(http.StatusFound, "/profile")
    return
  }
  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "login": login,
    "error": true,
  })
}

// Logout request
func Logout(c *gin.Context) {
  session.Close(c)
  c.Redirect(http.StatusFound, "/")
}

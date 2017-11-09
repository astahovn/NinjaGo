package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/astahovn/ninja/models/post"
  "github.com/astahovn/ninja/models/user"
  "github.com/astahovn/ninja/lib/session"
  "github.com/astahovn/ninja/lib/db"
)

func Index(c *gin.Context) {
  if session.GetAuth().UserId > 0 {
    c.Redirect(http.StatusFound, "/profile")
    return
  }

  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "title": "Index",
    "posts": post.FetchLast(),
    "auth": session.GetAuth().Nick,
  })
}

func Register(c *gin.Context) {
  c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
  })
}

func RegisterPost(c *gin.Context) {
  login := c.PostForm("login")
  //password := c.PostForm("password")
  var tmpUser user.User
  db.GetInstance().Where("username = ?", login).First(&tmpUser)
  if tmpUser.ID > 0 {
    c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
      "login": login,
      "errors": "Login is busy",
    })
    return
  }

  newUser := user.User{Username: login}
  db.GetInstance().Create(&newUser)

  c.HTML(http.StatusOK, "index/register.tmpl", gin.H{
    "login": login,
    "success": true,
    "errors": false,
  })
}

func LoginPost(c *gin.Context) {
  login := c.PostForm("login")
  //password := c.PostForm("password")
  var tmpUser user.User
  db.GetInstance().Where("username = ?", login).First(&tmpUser)
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

func Logout(c *gin.Context) {
  session.Close(c)
  c.Redirect(http.StatusFound, "/")
}

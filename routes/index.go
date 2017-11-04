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
  c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
    "title": "Index",
    "posts": post.FetchLast(),
    "auth": session.GetAuth().Realname,
  })
}

func Login(c *gin.Context) {
  c.HTML(http.StatusOK, "index/login.tmpl", gin.H{
  })
}

func LoginPost(c *gin.Context) {
  login := c.PostForm("login")
  //password := c.PostForm("password")
  var tmpUser user.User
  db.GetInstance().Where("username = ?", login).First(&tmpUser)
  if tmpUser.ID > 0 {
    session.Init(c, tmpUser.ID, c.Request.UserAgent())
    c.Redirect(http.StatusFound, "/")
    return
  }
  c.HTML(http.StatusOK, "index/login.tmpl", gin.H{
    "login": login,
    "error": true,
  })
}

func Logout(c *gin.Context) {
  session.Close(c)
  c.Redirect(http.StatusFound, "/")
}

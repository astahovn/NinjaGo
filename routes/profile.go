package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/astahovn/ninja/lib/session"
)

func ProfileIndex(c *gin.Context) {
  if session.GetAuth().UserId == 0 {
    c.Redirect(http.StatusFound, "/")
    return
  }
  c.HTML(http.StatusOK, "profile/index.tmpl", gin.H{
    "title": "Profile",
    "auth": session.GetAuth().Realname,
  })
}

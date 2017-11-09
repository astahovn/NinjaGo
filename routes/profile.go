package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/astahovn/ninja/lib/session"
  "github.com/astahovn/ninja/lib/db"
  "github.com/astahovn/ninja/models/user"
)

// Profile index page
func ProfileIndex(c *gin.Context) {
  if session.GetAuth().UserId == 0 {
    c.Redirect(http.StatusFound, "/")
    return
  }

  var authUser user.User
  db.GetInstance().Where("id = ?", session.GetAuth().UserId).First(&authUser)

  c.HTML(http.StatusOK, "profile/index.tmpl", gin.H{
    "title": "Profile",
    "auth": authUser,
  })
}

package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/astahovn/ninja/lib/session"
  "github.com/astahovn/ninja/models/user"
)

// Profile index page
func ProfileIndex(c *gin.Context) {
  if session.IsGuest(c) {
    c.Redirect(http.StatusFound, "/")
    return
  }

  authUser, _ := user.LoadById(session.GetAuth(c).UserId)

  c.HTML(http.StatusOK, "profile/index.tmpl", gin.H{
    "title": "Profile",
    "auth": authUser,
  })
}

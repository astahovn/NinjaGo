package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/astahovn/ninja/lib/session"
  "net/http"
)

func Init(engine *gin.Engine) {
  // init templates
  engine.LoadHTMLGlob("tpl/**/*")

  // init static
  engine.Static("/images", "./assets/images")
  engine.Static("/css", "./assets/css")

  // init dynamic
  engine.GET("/", Index)
  engine.GET("/register", Register)
  engine.POST("/register_user", RegisterPost)
  engine.POST("/login", LoginPost)
  engine.GET("/logout", Logout)

  engine.GET("/profile", ProfileIndex)
  engine.GET("/profile/edit", ProfileEdit)
  engine.POST("/profile/edit_save", ProfileEditSave)
}

var guestRoutes = map[string]bool{
    "/": true,
    "/register": true,
    "/register_user": true,
    "/login": true,
  }

// Access middleware
func Access() gin.HandlerFunc {
  return func(c *gin.Context) {
    if session.IsGuest(c) && !guestRoutes[c.Request.URL.Path] {
      c.Redirect(http.StatusFound, "/")
      c.Abort()
      return
    }
    if !session.IsGuest(c) && guestRoutes[c.Request.URL.Path] {
      c.Redirect(http.StatusFound, "/profile")
      c.Abort()
      return
    }

    c.Next()
  }
}

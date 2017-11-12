package routes

import (
  "github.com/gin-gonic/gin"
  "github.com/astahovn/ninja/lib/session"
  "strings"
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

// Access middleware
func Access() gin.HandlerFunc {
  return func(c *gin.Context) {
    if strings.Contains(c.Request.URL.Path, "/profile") {
      if session.IsGuest(c) {
        c.Redirect(http.StatusFound, "/")
        c.Abort()
        return
      }
    }

    c.Next()
  }
}

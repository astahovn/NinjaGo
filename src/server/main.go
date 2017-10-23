package main

import (
  "github.com/gin-gonic/gin"

  "routes"
  "lib/db"
  "lib/session"
)

func main() {
  // init db instance and defer close
  dbObj := db.GetInstance()
  defer dbObj.Close()

  // init engine, routes and run app
  engine := gin.Default()
  engine.Use(session.Auth())
  routes.Init(engine)
  engine.Run(":8080")
}

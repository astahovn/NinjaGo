package main

import (
  "github.com/gin-gonic/gin"

  "github.com/astahovn/ninja/routes"
  "github.com/astahovn/ninja/lib/db"
  "github.com/astahovn/ninja/lib/session"
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

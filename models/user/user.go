package user

import (
  "github.com/astahovn/ninja/lib/db"
)

type User struct {
  ID int
  Username string
  Password string
  Nick string
}

func LoadById(id int) User {
  var tmpUser User
  db.GetInstance().Where("id = ?", id).First(&tmpUser)
  return tmpUser
}

func LoadByUsername(login string) User {
  var tmpUser User
  db.GetInstance().Where("username = ?", login).First(&tmpUser)
  return tmpUser
}

func Register(user User) (Error string) {
  tmpUser := LoadByUsername(user.Username)
  if tmpUser.ID > 0 {
    Error = "Login is busy"
    return
  }

  db.GetInstance().Create(&user)

  Error = ""
  return
}
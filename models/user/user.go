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

func LoadById(id int) (tmpUser User, found bool) {
  notFound := db.GetInstance().Where("id = ?", id).First(&tmpUser).RecordNotFound()
  found = !notFound
  return
}

func LoadByUsername(login string) (tmpUser User, found bool) {
  notFound := db.GetInstance().Where("username = ?", login).First(&tmpUser).RecordNotFound()
  found = !notFound
  return
}

func Register(user User) (Error string) {
  _, found := LoadByUsername(user.Username)
  if found {
    Error = "Login is busy"
    return
  }

  db.GetInstance().Create(&user)

  Error = ""
  return
}
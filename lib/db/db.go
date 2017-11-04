package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init() *gorm.DB {
  db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=my_habr sslmode=disable password=postgres")
  if err != nil {
    panic("failed to connect database")
  }
  return db
}

func GetInstance() *gorm.DB {
  if db == nil {
    db = Init()
  }
  return db
}

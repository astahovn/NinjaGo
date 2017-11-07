package db

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "os"
  "encoding/json"
  "fmt"
)

const ConfigFile = "config/db.json"

type Configuration struct {
  Database string
  Host string
  Username string
  Password string
}

var db *gorm.DB

func Init() *gorm.DB {
  config := ReadConfiguration()
  dbConfigString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", config.Host, config.Username, config.Database, config.Password)

  db, err := gorm.Open("postgres", dbConfigString)
  if err != nil {
    panic("failed to connect database")
  }
  return db
}

func ReadConfiguration() Configuration {
  file, _ := os.Open(ConfigFile)
  decoder := json.NewDecoder(file)
  configuration := Configuration{}
  err := decoder.Decode(&configuration)
  if err != nil {
    panic("failed to load database configuration file db.json")
  }
  return configuration
}

func GetInstance() *gorm.DB {
  if db == nil {
    db = Init()
  }
  return db
}

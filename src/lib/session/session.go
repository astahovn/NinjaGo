package session

import (
  "time"
  "lib/db"
  "github.com/gin-gonic/gin"
)

type Session struct {
  ID int
  UserId int
  AuthToken string
  DateLogin time.Time
  UserAgent string
}

var currentSession Session

func Init(c *gin.Context, userId int, userAgent string) {
  RemoveOldSession(userId)
  sessionItem := Session{UserId: userId, AuthToken: "123", DateLogin: time.Now(), UserAgent: userAgent}
  db.GetInstance().NewRecord(sessionItem)
  db.GetInstance().Create(&sessionItem)

  c.SetCookie("token", "123", 60 * 60 * 24, "/", "", false, false)
}

func Close(c *gin.Context) {
  if currentSession.AuthToken != "" {
    db.GetInstance().Exec("delete from sessions where auth_token = ?", currentSession.AuthToken)
  }
  c.SetCookie("token", "", 0, "/", "", false, false)
  currentSession = Session{}
}

func CheckSession(c *gin.Context) {
  token, err := c.Cookie("token")
  if err != nil {
    currentSession = Session{}
    return
  }
  db.GetInstance().Where("auth_token = ?", token).First(&currentSession)
}

func RemoveOldSession(userId int) {
  db.GetInstance().Exec("delete from sessions where user_id = ?", userId)
}

func GetAuth() string {
  if currentSession.UserId > 0 {

    type Result struct {
      Username string
    }
    var result Result
    db.GetInstance().Raw("SELECT username FROM users WHERE id = ?", currentSession.UserId).Scan(&result)

    return result.Username
  }
  return ""
}
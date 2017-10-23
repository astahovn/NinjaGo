package session

import (
  "time"
  "lib/db"
  "github.com/gin-gonic/gin"
  "models/user"
)

type Session struct {
  ID int
  UserId int
  AuthToken string
  DateLogin time.Time
  UserAgent string
}

type AuthData struct {
  UserId int
  Realname string
}

var currentSession Session

func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {

    token, err := c.Cookie("token")
    if err != nil {
      currentSession = Session{}
      return
    }
    db.GetInstance().Where("auth_token = ?", token).First(&currentSession)

    c.Next()
  }
}

func Init(c *gin.Context, userId int, userAgent string) {
  RemoveOldSession(userId)
  sessionItem := Session{UserId: userId, AuthToken: "123", DateLogin: time.Now(), UserAgent: userAgent}
  db.GetInstance().NewRecord(sessionItem)
  db.GetInstance().Create(&sessionItem)

  c.SetCookie("token", "123", 60 * 60 * 24, "/", "", false, false)
}

func Close(c *gin.Context) {
  if currentSession.AuthToken != "" {
    db.GetInstance().Delete(Session{}, "auth_token = ?", currentSession.AuthToken)
  }
  c.SetCookie("token", "", 0, "/", "", false, false)
  currentSession = Session{}
}

func RemoveOldSession(userId int) {
  db.GetInstance().Delete(Session{}, "user_id = ?", userId)
}

func GetAuth() AuthData {
  if currentSession.UserId > 0 {
    var tmpUser user.User
    db.GetInstance().Where("id = ?", currentSession.UserId).First(&tmpUser)
    return AuthData{
      UserId: tmpUser.ID,
      Realname: tmpUser.Realname,
    }
  }
  return AuthData{}
}
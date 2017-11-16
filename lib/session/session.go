package session

import (
  "github.com/gin-gonic/gin"
  "time"
  "crypto/md5"
  "encoding/hex"
  "github.com/astahovn/ninja/lib/db"
  "github.com/astahovn/ninja/models/user"
)

type Session struct {
  ID int
  UserId int
  AuthToken string
  DateLogin time.Time
  UserAgent string
  Data string
}

const KeyHasSession = "HasSession"
const KeySession = "Session"
const KeyActiveUser = "ActiveUser"

// Auth middleware
func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {

    token, err := c.Cookie("token")
    if err != nil {
      c.Set(KeyHasSession, false)
      c.Next()
      return
    }

    var currentSession Session
    if db.GetInstance().Where("auth_token = ?", token).First(&currentSession).RecordNotFound() {
      c.Set(KeyHasSession, false)
      c.SetCookie("token", "", 0, "/", "", false, false)

    } else {
      c.Set(KeyHasSession, true)
      c.Set(KeySession, currentSession)
      activeUser, _ := user.LoadById(currentSession.UserId)
      c.Set(KeyActiveUser, activeUser)
    }

    c.Next()
  }
}

// Init new session
func Init(c *gin.Context, userId int) {
  // Remove old session, if exists
  db.GetInstance().Delete(Session{}, "user_id = ?", userId)

  // Generate auth token
  var hashMD5 = md5.New()
  hashMD5.Write([]byte(time.Now().String()))
  token := hex.EncodeToString(hashMD5.Sum(nil))

  // Create session DB record
  sessionItem := Session{
    UserId: userId,
    AuthToken: token,
    DateLogin: time.Now(),
    UserAgent: c.Request.UserAgent(),
  }
  db.GetInstance().NewRecord(sessionItem)
  db.GetInstance().Create(&sessionItem)

  // Setup cookie
  c.SetCookie("token", token, 60 * 60 * 24, "/", "", false, false)
}

// Close existed session
func Close(c *gin.Context) {
  if c.GetBool(KeyHasSession) {
    var ISession interface{}
    ISession, _ = c.Get(KeySession)
    db.GetInstance().Delete(Session{}, "auth_token = ?", ISession.(Session).AuthToken)
  }
  c.SetCookie("token", "", 0, "/", "", false, false)
  c.Set(KeyHasSession, false)
}

// Check active user is guest
func IsGuest(c *gin.Context) bool {
  return !c.GetBool(KeyHasSession)
}

// Get authenticated user
func GetActiveUser(c *gin.Context) user.User {
  var IUser interface{}
  IUser, _ = c.Get(KeyActiveUser)
  return IUser.(user.User)
}
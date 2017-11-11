package session

import (
  "github.com/gin-gonic/gin"
  "time"
  "crypto/md5"
  "encoding/hex"
  "encoding/json"
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

type AuthData struct {
  UserId int
  Username string
  Nick string
}

const KeyHasSession = "HasSession"
const KeySession = "Session"
const KeyAuthData = "AuthData"

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
      c.Set(KeySession, currentSession)
      c.Set(KeyHasSession, true)

      byt := []byte(currentSession.Data)
      var authData AuthData
      if err := json.Unmarshal(byt, &authData); err != nil {
        panic(err)
      }
      c.Set(KeyAuthData, authData)
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

  // Prepare session data
  objUser := user.LoadById(userId)
  authData := AuthData{
    UserId: userId,
    Username: objUser.Username,
    Nick: objUser.Nick,
  }
  authDataJson, _ := json.Marshal(&authData)

  // Create session DB record
  sessionItem := Session{
    UserId: userId,
    AuthToken: token,
    DateLogin: time.Now(),
    UserAgent: c.Request.UserAgent(),
    Data: string(authDataJson),
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
    currentSession := ISession.(Session)
    db.GetInstance().Delete(Session{}, "auth_token = ?", currentSession.AuthToken)
  }
  c.SetCookie("token", "", 0, "/", "", false, false)
  c.Set(KeyHasSession, false)
}

// Check active user is guest
func IsGuest(c *gin.Context) bool {
  return !c.GetBool(KeyHasSession)
}

// Get AuthData for current session if exists
func GetAuth(c *gin.Context) AuthData {
  if !c.GetBool(KeyHasSession) {
    return AuthData{}
  }

  var IAuthData interface{}
  IAuthData, _ = c.Get(KeyAuthData)
  return IAuthData.(AuthData)
}
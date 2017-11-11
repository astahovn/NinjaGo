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

// Auth middleware
func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {

    token, err := c.Cookie("token")
    if err != nil {
      c.Set("currentSession", Session{})
      return
    }
    var currentSession Session
    db.GetInstance().Where("auth_token = ?", token).First(&currentSession)
    c.Set("currentSession", currentSession)

    c.Next()
  }
}

// Init new session
func Init(c *gin.Context, userId int, userAgent string) {
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
    UserAgent: userAgent,
    Data: string(authDataJson),
  }
  db.GetInstance().NewRecord(sessionItem)
  db.GetInstance().Create(&sessionItem)

  // Setup cookie
  c.SetCookie("token", token, 60 * 60 * 24, "/", "", false, false)
}

// Close existed session
func Close(c *gin.Context) {
  var ISession interface{}
  ISession, _ = c.Get("currentSession")
  currentSession := ISession.(Session)
  if currentSession.AuthToken != "" {
    db.GetInstance().Delete(Session{}, "auth_token = ?", currentSession.AuthToken)
  }
  c.SetCookie("token", "", 0, "/", "", false, false)
}

// Get AuthData for current session if exists
func GetAuth(c *gin.Context) AuthData {
  var ISession interface{}
  ISession, _ = c.Get("currentSession")
  currentSession := ISession.(Session)
  if currentSession.UserId > 0 {
    byt := []byte(currentSession.Data)
    var authData AuthData
    if err := json.Unmarshal(byt, &authData); err != nil {
      panic(err)
    }
    return authData
  }
  return AuthData{}
}
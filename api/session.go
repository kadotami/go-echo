package api

import (
  // "fmt"
  "time"
  "crypto/rand"
  "net/http"
  "encoding/hex"
  "golang.org/x/crypto/bcrypt"
  "github.com/labstack/echo"
)

type (
  Session struct {
    UserId uint `json:"user_id"`
    Token string `json:"token"`
    ExpiredDate time.Time
  }

  LoginUser struct {
    Email string `json:"email"`
    Password string `json:"password"`
  }
)

func Login(c echo.Context) error {
  u := new(LoginUser)
  c.Bind(u)
  user_id, is_match := PasswordMatch(u)
  if !is_match {
    return c.JSON(http.StatusBadRequest, map[string]string{"message":"passwordが違います"})
  }
  s := Session{UserId:user_id, Token:CreateToken(), ExpiredDate:time.Now()}
  db.Create(&s)
  return c.JSON(http.StatusOK, map[string]string{"token":s.Token})
}

func Logout(c echo.Context) error {
  if err := Authirization(c); err != nil {
    return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  }
  db.Table("sessions").Where("user_id = ?", CurrentUserId(c)).Delete("")
  return c.String(http.StatusOK, "ok")
}

func PasswordMatch(u *LoginUser) (uid uint, success bool) {
  user := new(User)
  db.Select("id, password").Where("email = ?", u.Email).First(&user)
  password, _ := hex.DecodeString(user.Password)
  if bcrypt.CompareHashAndPassword(password, []byte(u.Password)) != nil {
    return 0, false
  }
  return user.ID, true
}

func CreateToken() string {
  token := make([]byte, 40)
  rand.Read(token)
  return hex.EncodeToString(token)
}

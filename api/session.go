package api

import (
  "fmt"
  "time"
  "net/http"
  "encoding/hex"
  "golang.org/x/crypto/bcrypt"
  "github.com/labstack/echo"
)

type (
  Sessions struct {
    UserId uint `json:"user_id"`
    Token string `json:"token"`
    ExpiredDate time.Time
  }
)

func Login(c echo.Context) error {
  u := new(User)
  if err := c.Bind(u); err != nil {
    return err
  }
  fmt.Println(PasswordMatch(u))
  if !PasswordMatch(u) {
    return c.JSON(http.StatusBadRequest, map[string]string{"message":"passwordが違います"})
  }
  return c.JSON(http.StatusOK, map[string]string{"token":"token"})
}

func Logout(c echo.Context) error {
  return c.String(http.StatusOK, "ok")
}

func PasswordMatch(u *User) bool{
  var user User
  db.Select("password").Where("email = ?", u.Email).First(&user)
  password, error := hex.DecodeString(user.Password)
  fmt.Println(error)
  if bcrypt.CompareHashAndPassword(password, []byte(u.Password)) != nil {
    return false
  }
  return true
}


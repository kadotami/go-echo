package api

import (
  // "fmt"
  "net/http"
  "encoding/hex"
  "github.com/labstack/echo"
  "golang.org/x/crypto/bcrypt"
)

type User struct {
  Model
  Name  string `json:"name"`
  Email string `json:"email"`
  Password string `json:"password,omitempty"`
  PasswordConfimation string `sql:"-" json:"password_confirmation,omitempty"`
}

func GetUser(c echo.Context) error {
  if err := Authirization(c); err != nil {
    return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  }
  var user User
  db.Select("id, name, email").Where("id = ?", c.Param("id")).First(&user)
  return c.JSON(http.StatusOK, map[string]interface{}{"data": &user})
}

func SaveUser(c echo.Context) error {
  u := new(User)
  if err := c.Bind(u); err != nil {
    return err
  }
  if u.Password != u.PasswordConfimation {
    return c.JSON(http.StatusBadRequest, map[string]string{"message":"passwordは同じものをいれてください"})
  }
  u.Password = PasswordToHash(u.Password)
  db.Create(&u)
  return c.NoContent(http.StatusCreated)
}

func UpdateUser(c echo.Context) error {
  if err := Authirization(c); err != nil {
    return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  }
  return c.NoContent(http.StatusOK)
}

func DeleteUser(c echo.Context) error {
  if err := Authirization(c); err != nil {
    return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  }
  return c.String(http.StatusOK, "ok!")
}

func PasswordToHash(pass string) string {
  converted, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
  return hex.EncodeToString(converted[:])
}
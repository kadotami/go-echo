package api

import (
  // "fmt"
  "net/http"
  "encoding/hex"
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
  "golang.org/x/crypto/bcrypt"
)

type (
  User struct {
    gorm.Model
    Name  string
    Email string
    Password string
  }

  UserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
    PasswordConfimation string `json:"password_confirmation"`
  }
)

func GetUser(c echo.Context) error {
  id := c.Param("id")
  return c.JSON(http.StatusOK, map[string]string{"data":id})
}

func SaveUser(c echo.Context) error {
  u := new(User)
  request := new(UserRequest)
  if err := c.Bind(request); err != nil {
    return err
  }
  if request.Password != request.PasswordConfimation {
    return c.JSON(http.StatusBadRequest, map[string]string{"message":"passwordは同じものをいれてください"})
  }
  u.Name, u.Email, u.Password = request.Name, request.Email, PasswordToHash(request.Password)
  db.Create(&u)
  return c.String(http.StatusOK, "ok")
}

func UpdateUser(c echo.Context) error {
  id := c.Param("id")
  jsonMap := map[string]string{
    "foo": id,
    "hoge": "fuga",
  }
  return c.JSON(http.StatusOK, jsonMap)
}

func DeleteUser(c echo.Context) error {
  return c.String(http.StatusOK, "ok!")
}

func PasswordToHash(pass string) string {
  converted, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
  return hex.EncodeToString(converted[:])
}
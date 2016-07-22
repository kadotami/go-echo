package api

import (
  "net/http"
  "github.com/labstack/echo"
)

type User struct {
  Name  string `json:"name"`
  Email string `json:"email"`
}

func GetUser(c echo.Context) error {
  id := c.Param("id")
  jsonMap := map[string]string{
    "foo": id,
    "hoge": "fuga",
  }
  return c.JSON(http.StatusOK, jsonMap)
}

func SaveUser(c echo.Context) error {
  // user := new(User)
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
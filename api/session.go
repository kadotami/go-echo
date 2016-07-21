package api

import (
  "net/http"
  "github.com/labstack/echo"
)

func Login(c echo.Context) error {
  return c.String(http.StatusOK, "ok")
}

func Logout(c echo.Context) error {
  return c.String(http.StatusOK, "ok")
}
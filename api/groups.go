package api

import (
  "net/http"
  "github.com/labstack/echo"
)

type Group struct {
  Name  string `json:"name"`
}

func SaveGroup(c echo.Context) error {
  name := c.Param("name")
  if name == "" {
    return c.JSON(http.StatusBadRequest, map[string]string{"massage":"name is required!"})
  }
//   group := Group{Name: name}
//   db.Create(&group)
  return c.JSON(http.StatusBadRequest, map[string]string{"massage":"name is required!"})
}
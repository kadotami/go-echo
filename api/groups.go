package api

import (
  // "fmt"
  "net/http"
  "github.com/labstack/echo"
)

type Group struct {
  Model
  Name string `json:"name"`
}

func GetGroup(c echo.Context) error {
  var group Group
  db.Select("id, name").Where("id = ?", c.Param("id")).First(&group)
  return c.JSON(http.StatusOK, map[string]interface{}{"group": &group})
}

func GetGroups(c echo.Context) error {
  var groups []Group
  db.Select("id, name").Find(&groups)
  return c.JSON(http.StatusOK, map[string]interface{}{"data": &groups})
}

func SaveGroup(c echo.Context) error {
  g := new(Group)
  if err := c.Bind(g); err != nil {
    return err
  }
  db.Create(&g)
  return c.NoContent(http.StatusOK)
}

func DeleteGroup(c echo.Context) error {
  var group Group
  db.Where("id = ?", c.Param("id")).Delete(&group)
  return c.NoContent(http.StatusOK)
}
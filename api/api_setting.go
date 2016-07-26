package api

import (
  "time"
  "os"
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Model struct {
  ID        uint `json:"id"`
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
  DeletedAt *time.Time   `json:"-"`
}

func init() {
  db, _ = gorm.Open("mysql", "go_user:golangyours@/yours?charset=utf8&parseTime=True&loc=Local")
}

func Authirization(c echo.Context) error {
  var count int
  token := c.Request().Header().Get(echo.HeaderAuthorization)
  db.Table("sessions").Where("token = ?", token).Count(&count)
  if count == 0 {
    return os.ErrPermission
  }
  return nil
}
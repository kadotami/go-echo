package api

import (
  "time"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

type Model struct {
  ID        uint64 `json:"id"`
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
  DeletedAt *time.Time   `json:"-"`
}

func init() {
  db, _ = gorm.Open("mysql", "go_user:golangyours@/yours?charset=utf8&parseTime=True&loc=Local")
}

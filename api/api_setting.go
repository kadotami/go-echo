package api

import (
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func init() {
  db, _ = gorm.Open("mysql", "go_user:golangyours@/yours?charset=utf8&parseTime=True&loc=Local")
}

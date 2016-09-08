package main

import (
  "fmt"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"
  "./models"
)

func main() {
  db, err := gorm.Open("mysql", "go_user:golangyours@/yours?charset=utf8&parseTime=True&loc=Local")
  fmt.Println(err)
  db.AutoMigrate(&models.Documents{}, &models.Users{}, &models.Groups{}, &models.Sessions{})
}
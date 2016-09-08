package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  _ "github.com/go-sql-driver/mysql"
)

type Documents struct {
  gorm.Model
  Title string `gorm:"not null"`
  Url string `gorm:"not null;unique"`
  ImageUrl string
  CreatorId string
}
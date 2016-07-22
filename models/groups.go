package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  _ "github.com/go-sql-driver/mysql"
)

type Groups struct {
  gorm.Model
  Name string `gorm:"not null;unique"`
}
package models

import (
  "time"
)

type Sessions struct {
  UserId int `gorm:"not null"`
  Token string `gorm:"not null"`
  ExpiredDate time.Time
}


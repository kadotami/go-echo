package api

import (
  // "fmt"
  "net/http"
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
)


type Document struct {
  Model
  Title string `json:"title"`
  Url string `json:"url"`
  ImageUrl string `json:"image_url"`
  UserId uint `json:"user_id"`
  User ReturnUserInfo `json:"user",omitempty`
}

func GetDocuments(c echo.Context) error {
  // if err := Authirization(c); err != nil {
  //   return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  // }
  var documents []Document
  db.Preload("User", func(db *gorm.DB) *gorm.DB {
    return db.Table("users")
  }).Find(&documents)
  return c.JSON(http.StatusOK, map[string]interface{}{"data": &documents})
}

func SaveDocument(c echo.Context) error {
  if err := ExtensionAuthirization(c); err != nil {
    return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  }
  d := new(Document)
  if err := c.Bind(d); err != nil {
    return err
  }
  db.Create(&d)
  return c.NoContent(http.StatusCreated)
}

func DeleteDocument(c echo.Context) error {
  if err := Authirization(c); err != nil {
    return c.JSON(http.StatusUnauthorized, map[string]string{"message":"認証期限が切れています"})
  }
  var document Document
  db.Where("id = ?", c.Param("id")).Delete(&document)
  return c.NoContent(http.StatusOK)
}
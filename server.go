package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"

  "./api"
)

func main() {
  e := echo.New()
  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
  }))

  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  // user
  e.GET("/users/:id", api.GetUser)
  e.POST("/users", api.SaveUser)
  e.PUT("/users/:id", api.UpdateUser)
  e.DELETE("/users/:id", api.DeleteUser)

  // // group
  // e.GET("/groups/:id", api.GetGroup)
  // e.POST("/groups", api.SaveGroup)

  // session(login and logout)
  e.POST("/sessions", api.Login)
  e.DELETE("/sessions", api.Logout)

  e.Run(standard.New(":1323"))
}

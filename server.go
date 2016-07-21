package main

import (
  "net/http"
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"

  "./handler"
)

func main() {
  e := echo.New()
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  e.GET("/user", user.Show())

  e.Run(standard.New(":1323"))
}

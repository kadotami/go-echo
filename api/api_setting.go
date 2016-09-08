package api

import (
  "time"
  "os"
  "github.com/labstack/echo"
  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _ "github.com/go-sql-driver/mysql"
  "encoding/hex"
  "golang.org/x/crypto/bcrypt"
  "crypto/rand"
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

func ExtensionAuthirization(c echo.Context) error {
  return nil
}

func CurrentUserId(c echo.Context) uint {
  s := new(Session)
  token := c.Request().Header().Get(echo.HeaderAuthorization)
  db.Where("token = ?", token).Count(&s)
  return s.UserId
}

func PasswordToHash(pass string) string {
  converted, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
  return hex.EncodeToString(converted[:])
}

func PasswordMatch(u *LoginUser) (uid uint, success bool) {
  user := new(User)
  db.Select("id, password").Where("email = ?", u.Email).First(&user)
  password, _ := hex.DecodeString(user.Password)
  if bcrypt.CompareHashAndPassword(password, []byte(u.Password)) != nil {
    return 0, false
  }
  return user.ID, true
}

func CreateToken() string {
  token := make([]byte, 40)
  rand.Read(token)
  return hex.EncodeToString(token)
}
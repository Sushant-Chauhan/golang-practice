package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	FullName string
	Roles    *Role
}
type Role struct {
	gorm.Model
	Name string
}

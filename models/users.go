package models

import "gorm.io/gorm"

//users table
type User struct {
	gorm.Model //includes created, deleted , unique id
	Name       string
	Email      string
}

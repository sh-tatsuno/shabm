package entity

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Name string
	Age  int
}

package entity

import "github.com/jinzhu/gorm"

type Bookmark struct {
	gorm.Model
	URL    string
	Title  string
	UserId uint
}

package models

import "github.com/jinzhu/gorm"

type TestData struct {
	gorm.Model
	SamplerId uint
	Params    string
	Separator string
	Loop      bool
	Index     int
}

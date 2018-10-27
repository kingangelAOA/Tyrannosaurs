package models

import "github.com/jinzhu/gorm"

type Script struct {
	gorm.Model
	SamplerId uint
	Data      string
}

func (s *Script) Run() {

}
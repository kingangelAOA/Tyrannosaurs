package models

import "github.com/jinzhu/gorm"

type Header struct {
	gorm.Model
	HttpRequestId uint
	Key           string
	Value         string
}

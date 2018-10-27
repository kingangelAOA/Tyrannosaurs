package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Cookie struct {
	gorm.Model
	HttpRequestId uint
	key           string
	value         string
	domain        string
	Path          string
	secure        bool
	expires       time.Time
}

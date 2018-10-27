package models

import "github.com/jinzhu/gorm"

type HttpRequest struct {
	gorm.Model
	SamplerId    uint
	Name         string
	Comment      string
	Protocol     string
	Host         string
	Port         string
	Method       string
	Path         string
	Coding       string
	Error        string
	Body         string
	Query        string
	Form         string
	PostForm     string
	Cookies      []Cookie
	Headers      []Header
}

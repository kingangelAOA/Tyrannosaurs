package models

import "github.com/jinzhu/gorm"

type ThreadGroup struct {
	gorm.Model
	PlanId uint
	Comments                string
	ActionAfterError        int
	Number                  int
	RampUp                  int
	Duration                int64
	Scenarios               []Scenario `gorm:"ForeignKey:ThreadId"`
}
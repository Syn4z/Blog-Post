package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Name   string `json:"name"`
	Number int    `json:"articles"`
}

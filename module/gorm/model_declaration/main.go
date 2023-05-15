package modeldeclaration

import "gorm.io/gorm"

type Items struct {
	Item_name     string
	Item_category string
	gorm.Model
}

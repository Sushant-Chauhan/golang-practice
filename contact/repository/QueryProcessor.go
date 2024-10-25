package repository

import "github.com/jinzhu/gorm"

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)

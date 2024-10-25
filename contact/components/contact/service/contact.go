package service

import (
	"contactApp/utils/log"

	"github.com/jinzhu/gorm"
)

type ContactService struct {
	DB  *gorm.DB
	log log.Logger
}

func NewContactService(DB *gorm.DB,
	log log.Logger) *ContactService {
	return &ContactService{
		log: log,
		DB:  DB,
	}
}

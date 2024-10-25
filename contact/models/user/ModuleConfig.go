package user

import (
	"contactApp/utils/log"

	"github.com/jinzhu/gorm"
)

type UserModuleConfig struct {
	DB  *gorm.DB
	log log.Logger
}

func NewUserModuleConfig(DB *gorm.DB, log log.Logger) *UserModuleConfig {
	return &UserModuleConfig{DB: DB, log: log}
}

func (m *UserModuleConfig) TableMigration() {

	var models []interface{} = []interface{}{
		&User{}, &Role{},
	}
	for i := 0; i < len(models); i++ {
		err := m.DB.AutoMigrate(models[i]).Error
		if err != nil {
			m.log.Error(err.Error())
		}
	}
	//addforegn keys etc...

}

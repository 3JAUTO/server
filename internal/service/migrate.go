package service

import "github.com/JEDIAC/server/internal/model"

func MigrateModel() error {
	return dbmgr.AutoMigrate(&model.DBStaff{})
}

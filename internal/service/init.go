package service

import (
	"errors"

	"github.com/JEDIAC/server/internal/db"
	"gorm.io/gorm"
)

var (
	dbmgr *gorm.DB
)

func Init() error {
	dbmgr = db.GetDB()
	if dbmgr == nil {
		return errors.New("database manager not initialized")
	}
	return nil
}

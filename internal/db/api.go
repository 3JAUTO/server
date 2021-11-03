package db

import (
	"database/sql"

	"github.com/JEDIAC/server/internal/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type M = map[string]interface{}

// Init initialize database
func Init(connStr string, dbname string) error {
	log.Debug("connecting to database", zap.String("connstr", connStr))
	rawdb, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	// Ensure database exists
	tx, err := rawdb.Begin()
	if err != nil {
		return err
	}
	_, _ = tx.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	_, _ = tx.Exec("USE " + dbname)
	if err := tx.Commit(); err != nil {
		return err
	}

	// Dial with gorm
	dbConn, err := gorm.Open(mysql.New(mysql.Config{
		Conn: rawdb,
	}), &gorm.Config{
		Logger: logger.New(log.GetDBLogger(), logger.Config{
			Colorful: true,
		}),
	})
	if err != nil {
		return err
	}
	db = dbConn
	return nil
}

// GetDB ...
func GetDB() *gorm.DB {
	return db
}

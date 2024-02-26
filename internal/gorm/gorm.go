package gorm

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDb(dbPath string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{

			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level

			IgnoreRecordNotFoundError: true,  // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,  // Don't include params in the SQL log
			Colorful:                  false, // Disable color
		},
	)

	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{Logger: newLogger},
	)
	if err != nil {
		panic(err)
	}

	return db
}

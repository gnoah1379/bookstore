package database

import (
	"bookstore/internal/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"time"
)

func Open(cfg config.Database) (*gorm.DB, error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute)
	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

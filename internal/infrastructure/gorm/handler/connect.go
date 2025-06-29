package handler

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "root:password@tcp(query_db:3306)/sample_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, DBErrHandler(err)
	}
	if db, err := conn.DB(); err != nil {
		return nil, DBErrHandler(err)
	} else {
		if err := db.Ping(); err != nil {
			return nil, DBErrHandler(err)
		}

		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(100)
		db.SetConnMaxLifetime(time.Hour)

		conn.Logger = conn.Logger.LogMode(logger.Info)
		return conn, nil
	}
}

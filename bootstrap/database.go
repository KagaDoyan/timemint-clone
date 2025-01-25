package bootstrap

import (
	"fmt"
	"go-fiber/core/logs"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection(env *Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.Database.Username,
		url.QueryEscape(env.Database.Password),
		env.Database.DBHost,
		env.Database.DBPort,
		env.Database.DBName,
	)
	// Set up a custom pool configuration
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot connect to database")
	}
	logs.Info("database connection success")
	return db
}

package mysql

import (
	"elab-backend/configs"
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewService() (*gorm.DB, error) {
	config := configs.GetConfig()
	dsn := fmt.Sprintf(
		config.MySQL.DSN,
		config.MySQL.Username,
		config.MySQL.Password,
		config.MySQL.Address,
		config.MySQL.Database,
	)
	db, err := gorm.Open(
		mysql.Open(dsn),
	)
	if err != nil {
		slog.Error("gorm.Open failed %w", err)
		return nil, err
	}
	return db, nil
}

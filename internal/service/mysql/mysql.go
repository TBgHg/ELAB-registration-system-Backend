package mysql

import (
	"elab-backend/configs"
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewService(conf *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		conf.MySQL.DSN,
		conf.MySQL.Username,
		conf.MySQL.Password,
		conf.MySQL.Address,
		conf.MySQL.Database,
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

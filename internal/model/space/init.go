package space

import (
	"elab-backend/internal/service"
)

func Init() error {
	svc := service.GetService()
	err := svc.MySQL.AutoMigrate(&Space{})
	return err
}

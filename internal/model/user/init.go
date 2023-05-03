package user

import "elab-backend/internal/service"

func Init() error {
	svc := service.GetService()
	return svc.MySQL.AutoMigrate(&User{})
}

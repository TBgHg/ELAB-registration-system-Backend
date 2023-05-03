package content

import "elab-backend/internal/service"

func Init() error {
	svc := service.GetService()
	return svc.MySQL.AutoMigrate(&Content{}, &Like{})
}

package application

import "elab-backend/internal/service"

func Init() error {
	svc := service.GetService()
	err := svc.MySQL.AutoMigrate(&LongTextForm{}, &InterviewRoomSelection{}, &InterviewRoom{})
	return err
}

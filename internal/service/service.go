package service

import "elab-backend/internal/service/oidc"

var svc Service

type Service struct {
	Oidc *oidc.Service
}

func GetService() *Service {
	return &svc
}

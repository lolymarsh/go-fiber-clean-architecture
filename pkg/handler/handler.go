package handler

import (
	"loly_api/config"
	"loly_api/pkg/service"
)

type Handler struct {
	conf *config.Config
	svc  *service.Service
}

func NewHandler(conf *config.Config, svc *service.Service) *Handler {
	return &Handler{
		conf: conf,
		svc:  svc,
	}
}

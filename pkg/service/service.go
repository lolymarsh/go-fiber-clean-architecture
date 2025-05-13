package service

import (
	"loly_api/config"
	"loly_api/pkg/repository"
)

type Service struct {
	conf *config.Config
	repo *repository.Repository
}

func NewService(conf *config.Config, repo *repository.Repository) *Service {
	return &Service{
		conf: conf,
		repo: repo,
	}
}

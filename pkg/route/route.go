package route

import (
	"loly_api/pkg/handler"
	"loly_api/pkg/repository"
	"loly_api/pkg/service"
	"loly_api/server"
)

func SetupRoutes(s *server.Server) {

	repo := repository.NewRepository(s.Db)
	svc := service.NewService(s.Conf, repo)
	hdl := handler.NewHandler(s.Conf, svc)

	setupUserRoutes(s.App, hdl)
}

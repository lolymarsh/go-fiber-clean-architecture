package server

import (
	"fmt"
	"loly_api/config"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app  *fiber.App
	conf *config.Config
}

var (
	once   sync.Once
	server *Server
)

func NewServer(conf *config.Config) *Server {

	app := fiber.New(fiber.Config{
		BodyLimit:   conf.App.BodyLimit * 1024 * 1024,
		IdleTimeout: time.Duration(conf.App.IdleTimeout) * time.Second,
	})

	once.Do(func() {
		server = &Server{
			app:  app,
			conf: conf,
		}
	})

	return server
}

func (s *Server) Run() error {

	s.app.Use(recover.New())
	s.app.Use(logger.New(logger.Config{
		TimeZone:   "Asia/Bangkok",
		TimeFormat: "2006-01-02 15:04:05",
		Format:     `${ip} - - [${time}] "${method} ${url} ${protocol}" ${status}`,
	}))
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("%v", s.conf.App.AllowOrigins),
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return s.app.Listen(":" + s.conf.App.Port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}

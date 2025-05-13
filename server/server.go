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
	"github.com/jmoiron/sqlx"
)

type Server struct {
	App  *fiber.App
	Conf *config.Config
	Db   *sqlx.DB
}

var (
	once   sync.Once
	server *Server
)

func NewServer(conf *config.Config, db *sqlx.DB) *Server {

	app := fiber.New(fiber.Config{
		BodyLimit:   conf.App.BodyLimit * 1024 * 1024,
		IdleTimeout: time.Duration(conf.App.IdleTimeout) * time.Second,
	})

	once.Do(func() {
		server = &Server{
			App:  app,
			Conf: conf,
			Db:   db,
		}
	})

	return server
}

func (s *Server) Run() error {

	s.App.Use(recover.New())
	s.App.Use(logger.New(logger.Config{
		TimeZone:   "Asia/Bangkok",
		TimeFormat: "2006-01-02 15:04:05",
		Format:     `${ip} - - [${time}] "${method} ${url} ${protocol}" ${status}`,
	}))
	s.App.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("%v", s.Conf.App.AllowOrigins),
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	return s.App.Listen(":" + s.Conf.App.Port)
}

func (s *Server) Shutdown() error {
	return s.App.Shutdown()
}

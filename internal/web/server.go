package web

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
)

//go:embed all:public
var FS embed.FS

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	configureApp(app)

	return &Server{
		app: app,
	}
}

func configureApp(app *fiber.App) {
	app.Use("/*", filesystem.New(filesystem.Config{
		Root:       http.FS(FS),
		Browse:     true,
		PathPrefix: "public",
	}))

}

func (s *Server) Start() {
	log.Fatal(s.app.Listen(":3000"))
}

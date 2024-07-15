package http

import (
	"embed"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/startcodextech/octopuslb/internal/api"
	"log"
	"net/http"
)

//go:embed all:public
var FS embed.FS

type Server struct {
	app *fiber.App
	api *api.Api
}

func NewServer(api *api.Api) *Server {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	configureApp(app)

	return &Server{
		app: app,
		api: api,
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

func (s *Server) installModule(c *fiber.Ctx) error {
	err := s.api.InstallModule("test")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Module installed successfully"})
}

package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/app/dependencies"
	"github.com/urcop/go-fiber-template/internal/app/initializers"
	"github.com/urcop/go-fiber-template/internal/repository"
	"github.com/urcop/go-fiber-template/internal/services"
)

type Application struct{}

func InitApplication(app *fiber.App) {
	initializers.InitEnv()

	repos := repository.NewLinkRepository()
	linkService := services.NewLinkService(repos)

	container := &dependencies.Container{
		LinkService: linkService,
	}

	initializers.SetupRoutes(app, container)

}

package link

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/helpers"
	"github.com/urcop/go-fiber-template/internal/model"
	"github.com/urcop/go-fiber-template/internal/services"
	"github.com/urcop/go-fiber-template/internal/web"
	"net/http"
)

var _ web.Controller = (*Controller)(nil)

type Controller struct {
	linkService services.LinkService
}

func (ctrl *Controller) CreateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var link model.Link
	err := ctx.BodyParser(&link)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error parsing JSON",
			"error":   err.Error(),
		})
	}

	if link.Random {
		link.ShortLink = helpers.RandomUrl(8)
	}

	err = ctrl.linkService.Create(&link)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error creating link in db",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(link)
}

func (ctrl *Controller) GetLinks(ctx *fiber.Ctx) error {
	result, err := ctrl.linkService.GetAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error getting all links",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(result)
}

func (ctrl *Controller) GetLink(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	links, err := ctrl.linkService.Get(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error getting link",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(links)
}

func (ctrl *Controller) UpdateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var link *model.Link

	err := ctx.BodyParser(&link)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error parsing JSON",
			"error":   err.Error(),
		})
	}
	err = ctrl.linkService.Update(link)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error update link",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(link)
}

func (ctrl *Controller) DeleteLink(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := ctrl.linkService.Delete(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "error delete link",
			"error":   err.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "link deleted",
		"error":   nil,
	})
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {
	router := app.Group("/api/v1/link")

	router.Post("/", ctrl.CreateLink)
	router.Get("/", ctrl.GetLinks)
	router.Get("/:id/", ctrl.GetLink)
	router.Delete("/:id/", ctrl.DeleteLink)
	router.Put("/", ctrl.UpdateLink)
}

func NewLinkController(link services.LinkService) *Controller {
	return &Controller{
		linkService: link,
	}
}

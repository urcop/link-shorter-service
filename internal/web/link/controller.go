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

// CreateLink создает новую ссылку.
// @Summary Создание ссылки
// @Description Создает новую ссылку с опциональным коротким URL.
// @Accept json
// @Produce json
// @Param request body Link true "Запрос с данными ссылки"
// @Success 201 {object} Link "Созданная ссылка"
// @Failure 500 {object} fiber.Map "Внутренняя ошибка сервера"
// @Router /api/v1/link/ [post]
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

// GetLinks возвращает все доступные ссылки.
// @Summary Получение всех ссылок
// @Description Возвращает список всех доступных ссылок.
// @Produce json
// @Success 200 {array} Link "Список ссылок"
// @Failure 500 {object} fiber.Map "Внутренняя ошибка сервера"
// @Router /api/v1/link/ [get]
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

// GetLink возвращает информацию о конкретной ссылке.
// @Summary Получение ссылки по идентификатору
// @Description Возвращает информацию о ссылке по указанному идентификатору.
// @Produce json
// @Param id path string true "Идентификатор ссылки"
// @Success 200 {object} Link "Информация о ссылке"
// @Failure 500 {object} fiber.Map "Внутренняя ошибка сервера"
// @Router /api/v1/link/{id}/ [get]
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

// UpdateLink обновляет существующую ссылку.
// @Summary Обновление ссылки
// @Description Обновляет информацию о существующей ссылке.
// @Accept json
// @Produce json
// @Param request body Link true "Запрос с данными обновления ссылки"
// @Success 200 {object} Link "Обновленная ссылка"
// @Failure 500 {object} fiber.Map "Внутренняя ошибка сервера"
// @Router /api/v1/link/ [put]
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

// DeleteLink удаляет существующую ссылку.
// @Summary Удаление ссылки
// @Description Удаляет существующую ссылку по идентификатору.
// @Param id path string true "Идентификатор ссылки"
// @Success 200 {object} fiber.Map "Ссылка удалена успешно"
// @Failure 500 {object} fiber.Map "Внутренняя ошибка сервера"
// @Router /api/v1/link/{id}/ [delete]
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

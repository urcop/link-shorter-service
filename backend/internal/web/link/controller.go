package link

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urcop/go-fiber-template/internal/helpers"
	"github.com/urcop/go-fiber-template/internal/model"
	"github.com/urcop/go-fiber-template/internal/services"
	"github.com/urcop/go-fiber-template/internal/web"
	"github.com/urcop/go-fiber-template/internal/web/render"
	"github.com/urcop/go-fiber-template/internal/web/validators"
	"net/http"
)

var _ web.Controller = (*Controller)(nil)

type Controller struct {
	linkService services.LinkService
}

// CreateLink Create a new Link
// @Summary Create Link
// @Description Creates a new link with an optional short URL.
// @Accept json
// @Produce json
// @Param request body Link true "Query with Link data"
// @Success 201 {object} Link "Created Link"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /api/v1/link/ [post]
func (ctrl *Controller) CreateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")

	var link model.Link
	err := ctx.BodyParser(&link)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error parsing JSON")
	}

	if link.Random {
		link.ShortLink = helpers.RandomUrl(8)
	}

	err = validators.LinkValidator(&link)
	if err != nil {
		return render.SendError(ctx, http.StatusBadRequest, err, "invalid data")
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

// GetLinks returns all available links.
// @Summary Getting all links
// @Description Returns a list of all available links.
// @Produce json
// @Success 200 {array} Link "List of links"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /api/v1/link/ [get]
func (ctrl *Controller) GetLinks(ctx *fiber.Ctx) error {
	result, err := ctrl.linkService.GetAll()
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error getting all links")
	}
	return ctx.Status(http.StatusOK).JSON(result)
}

// GetLink Returns information about a specific link.
// @Summary Getting a link by short link
// @Description Returns information about the link by the short link.
// @Produce json
// @Param shortLink path string true "short link"
// @Success 200 {object} Link "Link info"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /api/v1/link/{id}/ [get]
func (ctrl *Controller) GetLink(ctx *fiber.Ctx) error {
	id := ctx.Params("shortLink")

	links, err := ctrl.linkService.Get(id)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error getting link")
	}
	return ctx.Status(http.StatusOK).JSON(links)
}

// UpdateLink Deletes an existing link.
// @Summary Update link
// @Description Update an existing reference by ID
// @Accept json
// @Produce json
// @Param request body Link true "Query with link update data"
// @Success 200 {object} Link "Updated link"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /api/v1/link/ [put]
func (ctrl *Controller) UpdateLink(ctx *fiber.Ctx) error {
	ctx.Accepts("application/json")
	var link model.Link

	err := ctx.BodyParser(&link)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error parsing JSON")
	}

	err = validators.LinkValidator(&link)
	if err != nil {
		return render.SendError(ctx, http.StatusBadRequest, err, "invalid data")
	}

	err = ctrl.linkService.Update(&link)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error update link")
	}
	return ctx.Status(http.StatusOK).JSON(link)
}

// DeleteLink Deletes an existing link.
// @Summary Delete link from db
// @Description Deletes an existing reference by ID
// @Param id path string true "Link ID"
// @Success 200 {object} fiber.Map "Link deleted successfully"
// @Failure 500 {object} fiber.Map "Internal Server Error"
// @Router /api/v1/link/{id}/ [delete]
func (ctrl *Controller) DeleteLink(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	err := ctrl.linkService.Delete(id)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error delete link")
	}
	return render.JSONAPIPayload(ctx, http.StatusOK, "link deleted")
}

// Redirect
// @Summary Redirect to a URL
// @Description Redirects to the target URL associated with the given short link.
// @ID redirect
// @Produce json
// @Param redirect path string true "Short link for redirection"
// @Success 302 {string} string "Redirects to the target URL"
// @Failure 404 {object} fiber.Map "Error response with 404 status code"
// @Router /r/{redirect} [get]
func (ctrl *Controller) Redirect(ctx *fiber.Ctx) error {
	param := ctx.Params("redirect")

	link, err := ctrl.linkService.Get(param)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error parse link")
	}

	err = ctrl.linkService.UpdateClicked(link)
	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error update link clicked")
	}

	if err != nil {
		return render.SendError(ctx, http.StatusInternalServerError, err, "error update link model")
	}

	return ctx.Redirect(link.Link, fiber.StatusTemporaryRedirect)
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {
	router := app.Group("/api/v1/link")

	router.Post("/", ctrl.CreateLink)
	router.Get("/", ctrl.GetLinks)
	router.Get("/:shortLink/", ctrl.GetLink)
	router.Delete("/:id/", ctrl.DeleteLink)
	router.Put("/", ctrl.UpdateLink)

	// localhost:3000/r/<short link>
	app.Get("/r/:redirect/", ctrl.Redirect)
}

func NewLinkController(link services.LinkService) *Controller {
	return &Controller{
		linkService: link,
	}
}

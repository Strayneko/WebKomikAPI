package comic

import (
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var Helper interfaces.Helper
var ComicService interfaces.ComicService

type handler struct {
	controller interfaces.ComicController
}

func NewController() interfaces.ComicController {
	Helper = helpers.New()
	ComicService = comic.New()
	return &handler{}
}

func (h *handler) GetComicList(ctx *fiber.Ctx) error {
	path := "manga/"
	orderBy := ctx.Query("order")

	if !slices.Contains(configs.ComicOrderParams, orderBy) {
		orderTypes := strings.Join(configs.ComicOrderParams, ",")
		return Helper.ResponseError(ctx, fiber.NewError(http.StatusBadRequest, "Order should be in "+orderTypes))
	}

	currentPage, err := Helper.ValidatePage(ctx)
	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "?page=" + strconv.Itoa(int(currentPage))
	}
	if len(orderBy) > 0 {
		path += "&order=" + configs.GetComicOrderBy(orderBy)
	}

	return ComicService.GetComicList(ctx, path, currentPage)
}

func (h *handler) GetSearchedComics(ctx *fiber.Ctx) error {
	query := ctx.Query("query", "")
	path := "?s=" + query
	currentPage, err := Helper.ValidatePage(ctx)

	if err != nil {
		return Helper.ResponseError(ctx, err)
	}
	if currentPage > 0 {
		path = "page/" + strconv.Itoa(int(currentPage)) + "/?s=" + query
	}
	return ComicService.GetComicList(ctx, path, currentPage)
}

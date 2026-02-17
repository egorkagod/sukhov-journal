package articles

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	svc ArticleService
}

func NewArticleHandler(svc ArticleService) *ArticleHandler {
	return &ArticleHandler{svc: svc}
}

func (h *ArticleHandler) GetView(c echo.Context) error {
	articleID, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	ctx := c.Request().Context()
	article, err := h.svc.GetRepo().GetByID(ctx, articleID)

	var appError AppError
	if errors.As(err, &appError) {
		return echo.NewHTTPError(appError.code, appError.Error())
	}

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, article)
}

func (h *ArticleHandler) CreateView(c echo.Context) error {
	var data ArticleCreateSchema

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userID, ok := (c.Get("UserID")).(uint64)
	if !ok {
		return echo.NewHTTPError(500, "Не удалось идентифицировать пользователя")
	}

	dto := ArticleCreateServiceDTO{AuthorID: userID, Title: data.Title, Body: data.Body}
	ctx := c.Request().Context()
	err := h.svc.Create(ctx, dto)

	var appError AppError
	if errors.As(err, &appError) {
		return echo.NewHTTPError(appError.code, appError.Error())
	}

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]string{"message": "Статья успешно создана"})
}

func (h *ArticleHandler) EditView(c echo.Context) error {
	var data ArticleEditSchema

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userID, ok := (c.Get("UserID")).(uint64)
	if !ok {
		return echo.NewHTTPError(500, "Не удалось идентифицировать пользователя")
	}

	dto := ArticleEditServiceDTO{ID: data.ArticleID, UserID: userID, Title: data.Title, Body: data.Body}
	ctx := c.Request().Context()
	err := h.svc.Edit(ctx, dto)

	var appError AppError
	if errors.As(err, &appError) {
		return echo.NewHTTPError(appError.code, appError.Error())
	}

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]string{"message": "Статья успешно изменена"})
}

func (h *ArticleHandler) DeleteView(c echo.Context) error {
	articleID, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userID, ok := (c.Get("UserID")).(uint64)
	if !ok {
		return echo.NewHTTPError(500, "Не удалось идентифицировать пользователя")
	}

	dto := ArticleDeleteServiceDTO{ID: articleID, UserID: userID}
	ctx := c.Request().Context()
	err = h.svc.Delete(ctx, dto)

	var appError AppError
	if errors.As(err, &appError) {
		return echo.NewHTTPError(appError.code, appError.Error())
	}

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]string{"message": "Статья успешно удалена"})
}

func (h *ArticleHandler) VoiceOverView(c echo.Context) error {
	articleID, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	ctx := c.Request().Context()
	path, err := h.svc.GetAudioPath(ctx, articleID)

	var appError AppError
	if errors.As(err, &appError) {
		return echo.NewHTTPError(appError.code, appError.Error())
	}

	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	if path == "" {
		return c.JSON(200, map[string]string{"message": "Озвучка еще выполнена"})
	}

	return c.File(path)
}

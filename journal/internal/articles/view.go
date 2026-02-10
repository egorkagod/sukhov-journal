package articles

import (
	// "strconv"

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
	articleID, err := strconv.ParseInt(c.QueryParam("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	ctx := c.Request().Context()
	article, err := h.svc.GetRepo().GetByID(ctx, articleID)
	switch err {
	case ArticleNotFoundErr:
		return echo.NewHTTPError(404, ArticleNotFoundErr.Error())
	case nil:
		return c.JSON(200, article)
	default:
		return echo.NewHTTPError(500, err.Error())
	}
}

func (h *ArticleHandler) CreateView(c echo.Context) error {
	var data ArticleCreateSchema

	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	userID, ok := (c.Get("UserID")).(int64)
	if !ok {
		return echo.NewHTTPError(500, "Не удалось идентифицировать пользователя")
	}

	dto := ArticleCreateServiceDTO{AuthorID: userID, Title: data.Title, Body: data.Body}
	ctx := c.Request().Context()
	err := h.svc.Create(ctx, dto)
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

	userID, ok := (c.Get("UserID")).(int64)
	if !ok {
		return echo.NewHTTPError(500, "Не удалось идентифицировать пользователя")
	}

	dto := ArticleEditServiceDTO{ID: data.ArticleID, UserID: userID, Title: data.Title, Body: data.Body}
	ctx := c.Request().Context()
	err := h.svc.Edit(ctx, dto)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, map[string]string{"message": "Статья успешно изменена"})
}

// func (h *ArticleHandler) DeleteView(c echo.Context) error {
// 	articleID, err := strconv.ParseInt(c.QueryParam("id"), 10, 64)
// 	if err != nil {
// 		return echo.NewHTTPError(400, err.Error())
// 	}

// 	userID, ok := (c.Get("UserID")).(int64)
// 	if !ok {
// 		return echo.NewHTTPError(500, "Не удалось идентифицировать пользователя")
// 	}

// 	dto := ArticleDeleteServiceDTO{ID: articleID, UserID: userID}
// 	ctx := c.Request().Context()
// 	err = h.svc.Delete(ctx, dto)
// 	if err != nil {
// 		return echo.NewHTTPError(500, err.Error())
// 	}

// 	return c.JSON(200, map[string]string{"message": "Статья успешно удалена"})
// }

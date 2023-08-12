package articles_list

import (
	"errors"
	"incrowd-backend/application/api/apierror"
	"incrowd-backend/application/api/apiutils"
	"incrowd-backend/domain/models"
	"incrowd-backend/domain/services/articles_list_service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type ArticlesListHandler struct {
	articlesListService *articles_list_service.ArticlesListService
}

func NewArticlesListHandler(articlesListService *articles_list_service.ArticlesListService) *ArticlesListHandler {
	return &ArticlesListHandler{
		articlesListService: articlesListService,
	}
}

func (h *ArticlesListHandler) ArticlesList(c echo.Context) error {
	teamID := c.Param("teamID")

	page := apiutils.ReadIntQueryParam(c, "page", 0)
	count := apiutils.ReadIntQueryParam(c, "count", 10)
	sort := apiutils.ReadStringQueryParam(c, "sort", "_id", []string{"id", "published"})
	order := apiutils.ReadStringQueryParam(c, "order", "desc", []string{"desc", "asc"})

	md := models.MetaData{
		Page:  page,
		Count: count,
		Sort:  sort,
		Order: order,
	}

	articles, err := h.articlesListService.ArticlesList(c.Request().Context(), teamID, md)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrTeamIDNotFound):
			return apierror.Err(c, http.StatusNotFound, err)
		case os.IsTimeout(err):
			return apierror.Err(c, http.StatusBadGateway, err)
		default:
			return apierror.Err(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, convertResponse(md, articles))
}

package article_detail

import (
	"errors"
	"incrowd-backend/domain/models"
	"net/http"
	"os"

	"incrowd-backend/application/api/apierror"
	"incrowd-backend/domain/services/article_detail_service"

	"github.com/labstack/echo/v4"
)

type ArticleDetailHandler struct {
	articleDetailService *article_detail_service.ArticleDetailService
}

func NewArticleDetailHandler(articleDetailService *article_detail_service.ArticleDetailService) *ArticleDetailHandler {
	return &ArticleDetailHandler{
		articleDetailService: articleDetailService,
	}
}

func (h *ArticleDetailHandler) ArticleDetail(c echo.Context) error {
	teamID := c.Param("teamID")
	articleID := c.Param("articleID")

	article, err := h.articleDetailService.ArticleDetail(c.Request().Context(), teamID, articleID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrTeamIDNotFound) || errors.Is(err, models.ErrArticleIDNotFound):
			return apierror.Err(c, http.StatusNotFound, err)
		case os.IsTimeout(err):
			return apierror.Err(c, http.StatusBadGateway, err)
		default:
			return apierror.Err(c, http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusOK, convertResponse(article))
}

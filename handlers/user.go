package handlers

import (
	"net/http"

	"github.com/Listener430/payment_service/internal/rest/errors"
	"github.com/Listener430/payment_service/repository"
	"github.com/labstack/echo/v4"
)

// @Summary Checks user feature
// @Description Returns user feature state
// @Tags Purchase
// @Produce json
// @Param userId query string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {string} string "bad request"
// @Failure 404 {string} string "not found"
// @Router /check [get]
func CheckFeature(repo repository.Repo) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := c.QueryParam("userId")
		if u == "" {
			return c.String(http.StatusBadRequest, errors.ErrBadRequest.Error())
		}
		f, err := repo.GetUserFeature(u)
		if err != nil {
			return c.String(http.StatusNotFound, errors.ErrNotFound.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"userId":  f.UserID,
			"feature": f.Feature,
		})
	}
}

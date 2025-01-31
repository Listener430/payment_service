package handlers

import (
	"net/http"

	"github.com/Listener430/payment_service/internal/rest/errors"
	"github.com/Listener430/payment_service/models"
	"github.com/Listener430/payment_service/repository"
	"github.com/Listener430/payment_service/verify"
	"github.com/labstack/echo/v4"
)

// @Summary Apple purchase
// @Description Verifies Apple in-app purchase, unlocks feature
// @Tags Purchase
// @Accept json
// @Produce json
// @Param payload body models.PurchaseReq true "purchase"
// @Success 200 {string} string "unlocked"
// @Failure 400 {string} string "bad request"
// @Failure 401 {string} string "unauthorized"
// @Failure 500 {string} string "database operation failed"
// @Router /purchase [post]
func ApplePurchase(repo repository.Repo, v verify.Verifier) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.PurchaseReq
		if err := c.Bind(&req); err != nil {
			return c.String(http.StatusBadRequest, errors.ErrBadRequest.Error())
		}
		if !v.Check(req.ReceiptData) {
			return c.String(http.StatusUnauthorized, errors.ErrUnauthorized.Error())
		}
		if err := repo.UnlockFeature(req.UserID); err != nil {
			return c.String(http.StatusInternalServerError, errors.ErrDBFail.Error())
		}
		return c.String(http.StatusOK, "unlocked")
	}
}

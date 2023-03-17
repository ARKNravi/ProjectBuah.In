package handler

import (
	"ProjectBuahIn/models"
	"ProjectBuahIn/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CheckoutHandler interface {
	GetCheckout(*gin.Context)
	AddCheckout(*gin.Context)
}

type checkoutHandler struct {
	repo repository.CheckoutRepository
}

func NewCheckoutHandler() CheckoutHandler {
	return &checkoutHandler{
		repo: repository.NewCheckoutRepository(),
	}
}

func (h *checkoutHandler) GetCheckout(ctx *gin.Context) {
	chckStr := ctx.Param("checkout")
	chckID, err := strconv.Atoi(chckStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	checkout, err := h.repo.GetCheckout(chckID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.String(http.StatusOK, "Here is your checkout")
	ctx.JSON(http.StatusOK, checkout)

}

func (h *checkoutHandler) AddCheckout(ctx *gin.Context) {
	addrIDStr := ctx.Param("address")
	if addrID, err := strconv.Atoi(addrIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		cartIDsStr := ctx.Param("cartIDs")
		cartIDs := strings.Split(cartIDsStr, ",")
		checkouts := []models.Checkout{}
		for _, cartIDStr := range cartIDs {
			cartID, err := strconv.Atoi(cartIDStr)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			userID := ctx.GetFloat64("userID")
			checkout, err := h.repo.AddCheckout(int(userID), addrID, cartID)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			checkouts = append(checkouts, checkout)
		}
		ctx.String(http.StatusOK, "Welcome to the checkout")
		ctx.JSON(http.StatusOK, checkouts)
	}
}

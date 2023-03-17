package handler

import (
	"ProjectBuahIn/repository"
	"net/http"
	"strconv"

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
	ctx.JSON(http.StatusOK, checkout)

}

func (h *checkoutHandler) AddCheckout(ctx *gin.Context) {
	cartIDStr := ctx.Param("cart")
	if cartID, err := strconv.Atoi(cartIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		addrIDStr := ctx.Param("address")
		if addrID, err := strconv.Atoi(addrIDStr); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			userID := ctx.GetFloat64("userID")
			if checkout, err := h.repo.AddCheckout(int(userID), cartID, addrID); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, checkout)
			}
		}
	}
}

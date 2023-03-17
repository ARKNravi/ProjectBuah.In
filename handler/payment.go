package handler

import (
	"ProjectBuahIn/models"
	"ProjectBuahIn/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentHandler interface {
	GetPayment(*gin.Context)
	AddPayment(*gin.Context)
}

type paymentHandler struct {
	repo repository.PaymentRepository
}

// NewBuahHandler --> returns new handler for Buah entity
func NewPaymentHandler() PaymentHandler {
	return &paymentHandler{
		repo: repository.NewPaymentRepository(),
	}
}

func (h *paymentHandler) GetPayment(ctx *gin.Context) {
	payStr := ctx.Param("payment")
	payID, err := strconv.Atoi(payStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	payment, err := h.repo.GetPayment(payID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, payment)
}

func (h *paymentHandler) AddPayment(ctx *gin.Context) {
	var payment models.Payment
	chckIDStr := ctx.Param("checkout")
	if chckID, err := strconv.Atoi(chckIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		userID := ctx.GetFloat64("userID")
		if _, err := h.repo.AddPayment(payment, int(userID), int(chckID)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, "Payment added successfully")
		}
	}

}

package handler

import (
	"ProjectBuahIn/models"
	"ProjectBuahIn/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	GetCart(*gin.Context)
	GetAllCart(*gin.Context)
	AddCart(*gin.Context)
	UpdateCart(*gin.Context)
	DeleteCart(*gin.Context)
}

type cartHandler struct {
	repo repository.CartRepository
}

func NewCartHandler() CartHandler {
	return &cartHandler{
		repo: repository.NewCartRepository(),
	}
}

func (h *cartHandler) GetAllCart(ctx *gin.Context) {
	cart, err := h.repo.GetAllCart()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, cart)

}

func (h *cartHandler) GetCart(ctx *gin.Context) {
	cartStr := ctx.Param("cart")
	cartID, err := strconv.Atoi(cartStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cart, err := h.repo.GetCart(cartID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, cart)

}

func (h *cartHandler) AddCart(ctx *gin.Context) {
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	buahIDStr := ctx.Param("buah")
	if buahID, err := strconv.Atoi(buahIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		userID := ctx.GetFloat64("userID")
		if err := h.repo.AddCart(int(userID), buahID); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.String(http.StatusOK, "Product Successfully added to cart")
		}
		ctx.JSON(http.StatusOK, cart)
	}

}

func (h *cartHandler) UpdateCart(ctx *gin.Context) {
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	buahIDStr := ctx.Param("buah")
	if buahID, err := strconv.Atoi(buahIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		userID := ctx.GetFloat64("userID")
		if err := h.repo.UpdateCart(int(userID), buahID, cart); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.String(http.StatusOK, "Product Successfully Updated")
		}
		ctx.JSON(http.StatusOK, cart)
	}
}

func (h *cartHandler) DeleteCart(ctx *gin.Context) {
	var cart models.Cart
	if err := ctx.ShouldBindJSON(&cart); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	buahIDStr := ctx.Param("buah")
	if buahID, err := strconv.Atoi(buahIDStr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		userID := ctx.GetFloat64("userID")
		if err := h.repo.DeleteCart(int(userID), buahID, cart); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.String(http.StatusOK, "Product Successfully Deleted")
		}
		ctx.JSON(http.StatusOK, cart)
	}
}

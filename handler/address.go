package handler

//handler
import (
	"ProjectBuahIn/models"
	"ProjectBuahIn/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandler interface {
	GetAddress(*gin.Context)
	GetAllAddress(*gin.Context)
	AddAddress(*gin.Context)
	UpdateAddress(*gin.Context)
	DeleteAddress(*gin.Context)
}

type addressHandler struct {
	repo repository.AddressRepository
}

// NewBuahHandler --> returns new handler for Buah entity
func NewAddressHandler() AddressHandler {
	return &addressHandler{
		repo: repository.NewAddressRepository(),
	}
}

func (h *addressHandler) GetAddress(ctx *gin.Context) {
	id := ctx.Param("address")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address, err := h.repo.GetAddress(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, address)
}

func (h *addressHandler) GetAllAddress(ctx *gin.Context) {
	fmt.Println(ctx.Get("addressID"))
	address, err := h.repo.GetAllAddress()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, address)

}

func (h *addressHandler) AddAddress(ctx *gin.Context) {
	var address models.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	address, err := h.repo.AddAddress(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, address)

}

func (h *addressHandler) UpdateAddress(ctx *gin.Context) {
	var address models.Address
	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("address")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	address.ID = uint(intID)
	address, err = h.repo.UpdateAddress(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, address)

}

func (h *addressHandler) DeleteAddress(ctx *gin.Context) {

	var address models.Address
	addrStr := ctx.Param("address")
	addrID, _ := strconv.Atoi(addrStr)
	address.ID = uint(addrID)
	address, err := h.repo.DeleteAddress(address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, address)

}

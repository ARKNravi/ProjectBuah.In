package handler

//handler
import (
	"ProjectBuahIn/models"
	"ProjectBuahIn/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BuahHandler --> interface to Buah handler
type BuahHandler interface {
	GetBuah(*gin.Context)
	GetAllBuah(*gin.Context)
	AddBuah(*gin.Context)
	UpdateBuah(*gin.Context)
	DeleteBuah(*gin.Context)
	GetBuahByKondisi(*gin.Context)
	GetBuahByPriceDescending(*gin.Context)
	GetBuahByPriceAscending(*gin.Context)
}

type buahHandler struct {
	repo repository.BuahRepository
}

// NewBuahHandler --> returns new handler for Buah entity
func NewBuahHandler() BuahHandler {
	return &buahHandler{
		repo: repository.NewBuahRepository(),
	}
}

func (h *buahHandler) GetAllBuah(ctx *gin.Context) {
	Buah, err := h.repo.GetAllBuah()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Buah)

}

func (h *buahHandler) GetBuah(ctx *gin.Context) {
	prodStr := ctx.Param("buah")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Buah, err := h.repo.GetBuah(prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Buah)

}

func (h *buahHandler) AddBuah(ctx *gin.Context) {
	var Buah models.Buah
	if err := ctx.ShouldBindJSON(&Buah); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Buah, err := h.repo.AddBuah(Buah)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Buah)

}
func (h *buahHandler) UpdateBuah(ctx *gin.Context) {

	var Buah models.Buah
	if err := ctx.ShouldBindJSON(&Buah); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("buah")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	Buah.ID = uint(intID)
	Buah, err = h.repo.UpdateBuah(Buah)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Buah)

}
func (h *buahHandler) DeleteBuah(ctx *gin.Context) {

	var Buah models.Buah
	prodStr := ctx.Param("buah")
	prodID, _ := strconv.Atoi(prodStr)
	Buah.ID = uint(prodID)
	Buah, err := h.repo.DeleteBuah(Buah)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Buah)

}

func (h *buahHandler) GetBuahByKondisi(ctx *gin.Context) {
	kondisi := ctx.Param("kondisi")
	buahs, err := h.repo.GetBuahByKondisi(kondisi)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, buahs)
}

func (h *buahHandler) GetBuahByPriceDescending(ctx *gin.Context) {
	buahs, err := h.repo.GetBuahByPriceDescending()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, buahs)
}

func (h *buahHandler) GetBuahByPriceAscending(ctx *gin.Context) {
	buahs, err := h.repo.GetBuahByPriceAscending()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, buahs)
}

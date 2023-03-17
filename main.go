package main

//main
import (
	"ProjectBuahIn/handler"
	"ProjectBuahIn/initializer"
	"ProjectBuahIn/middleware"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
}

func main() {

	r := gin.Default()

	userHandler := handler.NewUserHandler()
	buahHandler := handler.NewBuahHandler()
	cartHandler := handler.NewCartHandler()
	orderHandler := handler.NewOrderHandler()
	addressHandler := handler.NewAddressHandler()
	checkoutHandler := handler.NewCheckoutHandler()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := r.Group("/api")
	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", userHandler.AddUser)
		userRoutes.POST("/signin", userHandler.SignInUser)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", userHandler.GetAllUser)
		userProtectedRoutes.GET("/:user", userHandler.GetUser)
		userProtectedRoutes.GET("/:user/products", userHandler.GetProductOrdered)
		userProtectedRoutes.PUT("/:user", userHandler.UpdateUser)
		userProtectedRoutes.DELETE("/:user", userHandler.DeleteUser)
	}

	addressRoutes := apiRoutes.Group("/address", middleware.AuthorizeJWT())
	{
		addressRoutes.GET("/", addressHandler.GetAllAddress)
		addressRoutes.GET("/:address", addressHandler.GetAddress)
		addressRoutes.POST("/", addressHandler.AddAddress)
		addressRoutes.PUT("/:address", addressHandler.UpdateAddress)
		addressRoutes.DELETE("/:address", addressHandler.DeleteAddress)
	}

	productRoutes := apiRoutes.Group("/buahs", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", buahHandler.GetAllBuah)
		productRoutes.GET("/:buah", buahHandler.GetBuah)
		productRoutes.GET("/kondisi/:kondisi", buahHandler.GetBuahByKondisi)
		productRoutes.GET("/pricedesc", buahHandler.GetBuahByPriceDescending)
		productRoutes.GET("/priceasc", buahHandler.GetBuahByPriceAscending)
		productRoutes.POST("/", buahHandler.AddBuah)
		productRoutes.PUT("/:buah", buahHandler.UpdateBuah)
		productRoutes.DELETE("/:buah", buahHandler.DeleteBuah)
	}
	AnonproductRoutes := apiRoutes.Group("/anonbuah")
	{
		AnonproductRoutes.GET("/", buahHandler.GetAllBuah)
		AnonproductRoutes.GET("/:buah", buahHandler.GetBuah)
		AnonproductRoutes.GET("kondisi/:kondisi", buahHandler.GetBuahByKondisi)
		AnonproductRoutes.GET("/pricedesc", buahHandler.GetBuahByPriceDescending)
		AnonproductRoutes.GET("/priceasc", buahHandler.GetBuahByPriceAscending)
	}

	cartRoutes := apiRoutes.Group("/cart", middleware.AuthorizeJWT())
	{
		cartRoutes.GET("/:cart", cartHandler.GetCart)
		cartRoutes.GET("user/:user", cartHandler.GetAllCart)
		cartRoutes.POST("/buah/:buah/quantity/:quantity", cartHandler.AddCart)
		cartRoutes.PUT("/:cart", cartHandler.UpdateCart)
		cartRoutes.DELETE("/:cart", cartHandler.DeleteCart)
	}
	checkoutRoutes := apiRoutes.Group("/checkout", middleware.AuthorizeJWT())
	{
		checkoutRoutes.GET("/:checkout", checkoutHandler.GetCheckout)
		checkoutRoutes.POST("/carts/:cart/address/:address", checkoutHandler.AddCheckout)
	}
	orderRoutes := apiRoutes.Group("/order", middleware.AuthorizeJWT())
	{
		orderRoutes.POST("/buah/:buah/quantity/:quantity", orderHandler.OrderProduct)
	}

	fileRoutes := r.Group("/file")
	{
		fileRoutes.POST("/single", handler.SingleFile)
		fileRoutes.POST("/multi", handler.MultipleFile)
	}

	port := os.Getenv("PORT")
	r.Run(":" + port)
}

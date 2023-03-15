package main

import (
	"ProjectBuahIn/handler"
	"ProjectBuahIn/initializer"
	"ProjectBuahIn/middleware"

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
		userProtectedRoutes.GET("/:user/buah/:buah/products", userHandler.GetProductOrdered)
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
		productRoutes.POST("/", buahHandler.AddBuah)
		productRoutes.PUT("/:buah", buahHandler.UpdateBuah)
		productRoutes.DELETE("/:buah", buahHandler.DeleteBuah)
	}

	cartRoutes := apiRoutes.Group("/cart", middleware.AuthorizeJWT())
	{
		cartRoutes.GET("/", cartHandler.GetAllCart)
		cartRoutes.POST("/buah/:buah/quantity/:quantity", cartHandler.AddCart)
		cartRoutes.PUT("/buah/:buah", cartHandler.UpdateCart)
		cartRoutes.DELETE("/:cart", cartHandler.DeleteCart)
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

	r.Run(":8090")

}

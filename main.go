package main

import (
	//"ProjectBuahIn/controllers"
	"ProjectBuahIn/handler"
	"ProjectBuahIn/initializer"
	"ProjectBuahIn/middleware"

	//"ProjectBuahIn/repository"
	//"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
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
	r.Run(":8070")

}

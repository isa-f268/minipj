package main

import (
	"main/config"
	"main/middleware"

	"main/handler"
	"main/repository"
	"main/service"
	"os"

	_ "main/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Book Rent App
// @version 1.0
// @description API untuk aplikasi book rent
// @termsOfService http://swagger.io/terms/

// @contact.name wapp
// @contact.url http://wapp.support.local
// @contact.email support@wapp.local

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load()
	db := config.DBInit()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	e := echo.New()
	e.HTTPErrorHandler = handler.ErrorHandler

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//users
	e.POST("/api/users/register", userHandler.RegisterUser)
	e.POST("/api/users/login", userHandler.LoginUser)
	e.POST("/api/users/topup", userHandler.TopUp, middleware.AuthMiddleware)
	e.GET("/api/users/payment-detail", userHandler.GetPaymentDetails, middleware.AuthMiddleware)
	e.GET("/api/users/book", userHandler.GetBook, middleware.AuthMiddleware)
	e.GET("/api/users/interbook", userHandler.GetInterBooks, middleware.AuthMiddleware)

	//Orders
	orderGroup := e.Group("/api/order")
	orderGroup.Use(middleware.AuthMiddleware)
	orderGroup.POST("", orderHandler.CreateOrder)
	orderGroup.POST("/payment", orderHandler.CreatePayment)
	orderGroup.POST("/midtrans", orderHandler.PayMidtrans)
	orderGroup.PATCH("/midtrans-status", orderHandler.UpdateStatusPayment)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}

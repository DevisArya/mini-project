package routes

import (
	"log"
	"miniproject/controller"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Pre(mid.AddTrailingSlash())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("SECRET_KEY")
	jwt := mid.JWT([]byte(secretKey))

	gArea := e.Group("/area")
	gArea.GET("/", controller.GetAreas, jwt)
	gArea.POST("/", controller.CreateArea, jwt)
	gArea.PUT("/:id/", controller.UpdateArea, jwt)
	gArea.DELETE("/:id/", controller.DeleteArea, jwt)

	gCleaner := e.Group("/cleaner")
	gCleaner.GET("/", controller.GetCleaners, jwt)
	gCleaner.POST("/", controller.CreateCleaner, jwt)
	gCleaner.GET("/:id/", controller.GetCleaner, jwt)
	gCleaner.PUT("/:id/", controller.UpdateCleaner, jwt)
	gCleaner.DELETE("/:id/", controller.DeleteCleaner, jwt)

	gCustomer := e.Group("/customer")
	gCustomer.GET("/", controller.GetCustomers, jwt)
	gCustomer.POST("/", controller.CreateCustomer)
	gCustomer.GET("/:id/", controller.GetCustomer, jwt)
	gCustomer.PUT("/:id/", controller.UpdateCustomer, jwt)
	gCustomer.DELETE("/:id/", controller.DeleteCustomer, jwt)

	gPayment := e.Group("/payment")
	gPayment.GET("/", controller.GetPayments, jwt)
	gPayment.POST("/", controller.CreatePayment, jwt)
	gPayment.GET("/:id/", controller.GetPayment, jwt)
	gPayment.PUT("/:id/", controller.UpdatePayment, jwt)
	gPayment.DELETE("/:id/", controller.DeletePayment, jwt)

	gServiceType := e.Group("/service-type")
	gServiceType.GET("/", controller.GetServiceTypes, jwt)
	gServiceType.POST("/", controller.CreateServiceType, jwt)
	gServiceType.GET("/:id/", controller.GetServiceType, jwt)
	gServiceType.PUT("/:id/", controller.UpdateServiceType, jwt)
	gServiceType.DELETE("/:id/", controller.DeleteServiceType, jwt)

	gStore := e.Group("/store")
	gStore.GET("/", controller.GetStores, jwt)
	gStore.POST("/", controller.CreateStore, jwt)
	gStore.GET("/:id/", controller.GetStore, jwt)
	gStore.PUT("/:id/", controller.UpdateStore, jwt)
	gStore.DELETE("/:id/", controller.DeleteStore, jwt)

	gTeam := e.Group("/team")
	gTeam.GET("/", controller.GetTeams, jwt)
	gTeam.POST("/", controller.CreateTeam, jwt)
	gTeam.GET("/:id/", controller.GetTeam, jwt)
	gTeam.PUT("/:id/", controller.UpdateTeam, jwt)
	gTeam.DELETE("/:id/", controller.DeleteTeam, jwt)

	gTransaction := e.Group("/transaction")
	gTransaction.GET("/", controller.GetTransactions, jwt)
	gTransaction.POST("/", controller.CreateTransaction, jwt)
	gTransaction.GET("/:id/", controller.GetTransaction, jwt)
	gTransaction.PUT("/:id/", controller.UpdateTransaction, jwt)
	gTransaction.DELETE("/:id/", controller.DeleteTransaction, jwt)

	gTransactionDetail := e.Group("/Transaction-Detail")
	gTransactionDetail.GET("/", controller.GetTransactionDetails, jwt)
	gTransactionDetail.POST("/", controller.CreateTransactionDetail, jwt)
	gTransactionDetail.GET("/:id/", controller.GetTransactionDetail, jwt)
	gTransactionDetail.PUT("/:id/", controller.UpdateTransactionDetail, jwt)
	gTransactionDetail.DELETE("/:id/", controller.DeleteTransactionDetail, jwt)

	return e
}

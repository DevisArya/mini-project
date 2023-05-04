package routes

import (
	c "miniproject/controller"
	m "miniproject/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Pre(mid.AddTrailingSlash())

	gArea := e.Group("/area")
	gArea.GET("/", c.GetAreas, m.IsloggedIn, m.IsAdmin)
	gArea.POST("/", c.CreateArea, m.IsloggedIn)
	gArea.PUT("/:id/", c.UpdateArea, m.IsloggedIn)
	gArea.DELETE("/:id/", c.DeleteArea, m.IsloggedIn)

	gCleaner := e.Group("/cleaner")
	gCleaner.GET("/", c.GetCleaners, m.IsloggedIn)
	gCleaner.POST("/", c.CreateCleaner, m.IsloggedIn)
	gCleaner.GET("/:id/", c.GetCleaner, m.IsloggedIn)
	gCleaner.PUT("/:id/", c.UpdateCleaner, m.IsloggedIn)
	gCleaner.DELETE("/:id/", c.DeleteCleaner, m.IsloggedIn)

	gCustomer := e.Group("/customer")
	gCustomer.GET("/", c.GetCustomers, m.IsloggedIn)
	gCustomer.POST("/", c.CreateCustomer)
	gCustomer.GET("/:id/", c.GetCustomer, m.IsloggedIn)
	gCustomer.PUT("/:id/", c.UpdateCustomer, m.IsloggedIn)
	gCustomer.DELETE("/:id/", c.DeleteCustomer, m.IsloggedIn)
	gCustomer.POST("/login/", c.LoginCustomer)

	gAdmin := e.Group("/admin")
	gAdmin.GET("/", c.GetAdmins, m.IsloggedIn)
	gAdmin.POST("/", c.CreateAdmin)
	gAdmin.GET("/:id/", c.GetAdmin, m.IsloggedIn)
	gAdmin.PUT("/:id/", c.UpdateAdmin, m.IsloggedIn)
	gAdmin.DELETE("/:id/", c.DeleteAdmin, m.IsloggedIn)
	gAdmin.POST("/login/", c.LoginAdmin)

	gPayment := e.Group("/payment")
	gPayment.GET("/", c.GetPayments, m.IsloggedIn)
	gPayment.POST("/", c.CreatePayment, m.IsloggedIn)
	gPayment.GET("/:id/", c.GetPayment, m.IsloggedIn)
	gPayment.PUT("/:id/", c.UpdatePayment, m.IsloggedIn)
	gPayment.DELETE("/:id/", c.DeletePayment, m.IsloggedIn)

	gServiceType := e.Group("/service-type")
	gServiceType.GET("/", c.GetServiceTypes, m.IsloggedIn)
	gServiceType.POST("/", c.CreateServiceType, m.IsloggedIn)
	gServiceType.GET("/:id/", c.GetServiceType, m.IsloggedIn)
	gServiceType.PUT("/:id/", c.UpdateServiceType, m.IsloggedIn)
	gServiceType.DELETE("/:id/", c.DeleteServiceType, m.IsloggedIn)

	gStore := e.Group("/store")
	gStore.GET("/", c.GetStores, m.IsloggedIn)
	gStore.POST("/", c.CreateStore, m.IsloggedIn)
	gStore.GET("/:id/", c.GetStore, m.IsloggedIn)
	gStore.PUT("/:id/", c.UpdateStore, m.IsloggedIn)
	gStore.DELETE("/:id/", c.DeleteStore, m.IsloggedIn)

	gTeam := e.Group("/team")
	gTeam.GET("/", c.GetTeams, m.IsloggedIn)
	gTeam.POST("/", c.CreateTeam, m.IsloggedIn)
	gTeam.GET("/:id/", c.GetTeam, m.IsloggedIn)
	gTeam.PUT("/:id/", c.UpdateTeam, m.IsloggedIn)
	gTeam.DELETE("/:id/", c.DeleteTeam, m.IsloggedIn)

	gTransaction := e.Group("/transaction")
	gTransaction.GET("/", c.GetTransactions, m.IsloggedIn)
	gTransaction.POST("/", c.CreateTransaction, m.IsloggedIn)
	gTransaction.GET("/:id/", c.GetTransaction, m.IsloggedIn)
	gTransaction.PUT("/:id/", c.UpdateTransaction, m.IsloggedIn)
	gTransaction.DELETE("/:id/", c.DeleteTransaction, m.IsloggedIn)

	gTransactionDetail := e.Group("/transaction-detail")
	gTransactionDetail.GET("/", c.GetTransactionDetails, m.IsloggedIn)
	gTransactionDetail.POST("/", c.CreateTransactionDetail, m.IsloggedIn)
	gTransactionDetail.GET("/:id/", c.GetTransactionDetail, m.IsloggedIn)
	gTransactionDetail.PUT("/:id/", c.UpdateTransactionDetail, m.IsloggedIn)
	gTransactionDetail.DELETE("/:id/", c.DeleteTransactionDetail, m.IsloggedIn)

	return e
}

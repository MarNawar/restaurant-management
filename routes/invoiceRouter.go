package routes

import (
	"github.com/MarNawar/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/invoices", controllers.GetInvoices())
	incommingRoutes.GET("/invoices/:invoice_id", controllers.GetInvoice())
	incommingRoutes.POST("/invoices", controllers.CreateInvoice())
	incommingRoutes.PATCH("/invoices/:menu_id", controllers.UpdateInvoice())
}
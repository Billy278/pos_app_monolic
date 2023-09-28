package server

import (
	routeCategories "github.com/Billy278/pos_app_monolic/modules/router/v1/categories"
	routePayment "github.com/Billy278/pos_app_monolic/modules/router/v1/payment"
	routeProduct "github.com/Billy278/pos_app_monolic/modules/router/v1/products"
	routeUser "github.com/Billy278/pos_app_monolic/modules/router/v1/users"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	g := gin.Default()
	g.Use(gin.Recovery())
	handler := initHandler()
	v1 := g.Group("api/v1")
	routePayment.NewPaymentRouter(v1, handler.PaymentCtrl)
	routeCategories.NewCategoriesRouter(v1, handler.CategoriesCtrl)
	routeProduct.NewProductRouter(v1, handler.ProductsCtrl)
	routeUser.NewUserRouter(v1, handler.UserCtrl)
	g.Run(":9090")

}

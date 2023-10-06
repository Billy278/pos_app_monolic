package payment

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/orders"
	"github.com/Billy278/pos_app_monolic/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewPaymentRouter(v1 *gin.RouterGroup, orderCtrl controllers.CtrlOrders) {

	g := v1.Group("/orders", middleware.AuthenticationBearer())

	// // register all router
	g.GET("", orderCtrl.List)
	g.POST("", orderCtrl.Add)
	g.GET(":id", orderCtrl.Detail)
}

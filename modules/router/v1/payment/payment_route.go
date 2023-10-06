package payment

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/payment"
	"github.com/Billy278/pos_app_monolic/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewPaymentRouter(v1 *gin.RouterGroup, paymentCtrl controllers.CtrlPayment) {

	g := v1.Group("/payments", middleware.AuthenticationBearer())

	// // register all router
	g.GET("", paymentCtrl.List)
	g.POST("", middleware.AuthorizationUser(), paymentCtrl.Created)
	g.GET(":id", paymentCtrl.FindByid)
	g.PUT(":id", middleware.AuthorizationUser(), paymentCtrl.Updated)
	g.DELETE(":id", middleware.AuthorizationUser(), paymentCtrl.Deleted)
}

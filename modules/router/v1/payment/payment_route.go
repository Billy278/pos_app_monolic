package payment

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/payment"
	"github.com/gin-gonic/gin"
)

func NewPaymentRouter(v1 *gin.RouterGroup, paymentCtrl controllers.CtrlPayment) {
	g := v1.Group("/payments")

	// // register all router
	g.GET("", paymentCtrl.List)
	g.POST("", paymentCtrl.Created)
	g.GET(":id", paymentCtrl.FindByid)
	g.PUT(":id", paymentCtrl.Updated)
	g.DELETE(":id", paymentCtrl.Deleted)
}

package products

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/products"
	"github.com/gin-gonic/gin"
)

func NewProductRouter(v1 *gin.RouterGroup, productCtrl controllers.CtrlProduct) {
	g := v1.Group("/products")

	// // register all router
	g.GET("", productCtrl.List)
	g.POST("", productCtrl.Created)
	g.GET(":id", productCtrl.FindByid)
	g.PUT(":id", productCtrl.Updated)
	g.DELETE(":id", productCtrl.Deleted)
}

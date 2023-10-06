package products

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/products"
	"github.com/Billy278/pos_app_monolic/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewProductRouter(v1 *gin.RouterGroup, productCtrl controllers.CtrlProduct) {
	g := v1.Group("/products", middleware.AuthenticationBearer())

	// // register all router
	g.GET("", productCtrl.List)
	g.POST("", middleware.AuthorizationUser(), productCtrl.Created)
	g.GET(":id", productCtrl.FindByid)
	g.PUT(":id", middleware.AuthorizationUser(), productCtrl.Updated)
	g.DELETE(":id", middleware.AuthorizationUser(), productCtrl.Deleted)
}

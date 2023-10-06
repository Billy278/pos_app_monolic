package categories

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/categories"
	"github.com/Billy278/pos_app_monolic/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewCategoriesRouter(v1 *gin.RouterGroup, categoriesCtrl controllers.CtrlCategories) {

	g := v1.Group("/categories", middleware.AuthenticationBearer())

	// // register all router
	g.GET("", categoriesCtrl.List)
	g.POST("", middleware.AuthorizationUser(), categoriesCtrl.Created)
	g.GET(":id", categoriesCtrl.FindByid)
	g.PUT(":id", middleware.AuthorizationUser(), categoriesCtrl.Updated)
	g.DELETE(":id", middleware.AuthorizationUser(), categoriesCtrl.Deleted)
}

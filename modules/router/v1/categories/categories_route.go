package categories

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/categories"
	"github.com/gin-gonic/gin"
)

func NewCategoriesRouter(v1 *gin.RouterGroup, categoriesCtrl controllers.CtrlCategories) {
	g := v1.Group("/categories")

	// // register all router
	g.GET("", categoriesCtrl.List)
	g.POST("", categoriesCtrl.Created)
	g.GET(":id", categoriesCtrl.FindByid)
	g.PUT(":id", categoriesCtrl.Updated)
	g.DELETE(":id", categoriesCtrl.Deleted)
}

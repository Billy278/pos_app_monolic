package users

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/users"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(v1 *gin.RouterGroup, userCtrl controllers.UserCtrl) {
	g := v1.Group("/users")

	// // register all router
	g.GET("", userCtrl.List)
	g.POST("", userCtrl.Created)
	g.GET(":id", userCtrl.FindByid)
	g.PUT(":id", userCtrl.Updated)
	g.DELETE(":id", userCtrl.Deleted)
}

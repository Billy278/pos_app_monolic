package users

import (
	controllers "github.com/Billy278/pos_app_monolic/modules/controllers/users"
	"github.com/Billy278/pos_app_monolic/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(v1 *gin.RouterGroup, userCtrl controllers.UserCtrl) {
	a := v1.Group("/auth/login")
	g := v1.Group("/users")
	// // register all router
	g.GET("", middleware.AuthenticationBearer(), middleware.AuthorizationUser(), userCtrl.List)
	g.POST("", userCtrl.Register)
	g.GET(":id", middleware.AuthenticationBearer(), middleware.AuthorizationUser(), userCtrl.FindByid)
	g.PUT(":id", middleware.AuthenticationBearer(), middleware.AuthorizationUser(), userCtrl.Updated)
	g.DELETE(":id", middleware.AuthenticationBearer(), middleware.AuthorizationUser(), userCtrl.Deleted)

	//==========================
	//group auth
	a.POST("", userCtrl.Login)

}

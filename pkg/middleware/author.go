package middleware

import (
	"net/http"
	"strconv"

	"github.com/Billy278/pos_app_monolic/db"
	modelsToken "github.com/Billy278/pos_app_monolic/modules/models/tokens"
	repository "github.com/Billy278/pos_app_monolic/modules/repository/users"
	"github.com/Billy278/pos_app_monolic/pkg/crypto"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
)

func AuthorizationUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accClaimIn, ok := ctx.Get("access_claim")
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
				Code:    http.StatusInternalServerError,
				Success: false,
				Message: responses.SomethingWentWrong,
			})
			return
		}
		var accessClaim modelsToken.AccessClaim
		crypto.ObjectMapper(accClaimIn, &accessClaim)
		dbIn := db.NewDBPostges()
		repoUser := repository.NewUserRepoImpl(dbIn)
		id, err := strconv.ParseUint(accessClaim.UserId, 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
				Code:    http.StatusInternalServerError,
				Success: false,
				Message: responses.SomethingWentWrong,
			})
			return
		}

		resUser, err := repoUser.RepoFindByid(ctx, id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
				Code:    http.StatusInternalServerError,
				Success: false,
				Message: responses.SomethingWentWrong,
			})
			return
		}
		if resUser.Role != "ADMIN" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Success: false,
				Message: responses.Unauthorized,
			})
			return
		}
		ctx.Next()

	}

}

package middleware

import (
	"net/http"
	"strings"

	models "github.com/Billy278/pos_app_monolic/modules/models/tokens"
	"github.com/Billy278/pos_app_monolic/pkg/crypto"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
)

const (
	Bearer        string = "Bearer "
	Authorization string = "Authorization"
)

func AuthenticationBearer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		author := ctx.GetHeader(Authorization)
		if author == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Success: false,
				Message: responses.Unauthorized,
			})
			return
		}
		token := strings.Split(author, Bearer)
		if len(token) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Success: false,
				Message: responses.Unauthorized,
			})
			return
		}
		var claim models.AccessClaim

		err := crypto.ParseAndVerifyToken(token[1], &claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				Code:    http.StatusUnauthorized,
				Success: false,
				Message: "invalid Token",
			})
			return
		}
		ctx.Set("access_claim", claim)
		ctx.Next()
	}

}

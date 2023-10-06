package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	models "github.com/Billy278/pos_app_monolic/modules/models/orders"
	modelToken "github.com/Billy278/pos_app_monolic/modules/models/tokens"
	services "github.com/Billy278/pos_app_monolic/modules/services/orders"
	"github.com/Billy278/pos_app_monolic/pkg/crypto"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlOrdersImpl struct {
	OrderSrv services.OrderSrv
	Validate *validator.Validate
}

func NewCtrlOrdersImpl(ordersrv services.OrderSrv, v *validator.Validate) CtrlOrders {
	return &CtrlOrdersImpl{
		OrderSrv: ordersrv,
		Validate: v,
	}
}

func (ctrl *CtrlOrdersImpl) List(ctx *gin.Context) {
	//get query
	limit, offset, err := ctrl.GetQuery(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidQuery,
		})
		return
	}
	resOrder, err := ctrl.OrderSrv.SrvList(ctx, limit, offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, responses.Response{
		Code:    http.StatusAccepted,
		Success: true,
		Message: responses.Success,
		Data:    resOrder,
	})
}
func (ctrl *CtrlOrdersImpl) Detail(ctx *gin.Context) {
	//get and validate param
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	resOrder, err := ctrl.OrderSrv.SrvDetail(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, responses.Response{
		Code:    http.StatusAccepted,
		Success: true,
		Message: responses.Success,
		Data:    resOrder,
	})
}
func (ctrl *CtrlOrdersImpl) Add(ctx *gin.Context) {
	// bind req
	reqOrder := models.Order{}
	err := ctx.ShouldBindJSON(&reqOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}
	//validate reqOrder with validator
	err = ctrl.Validate.Struct(&reqOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	//get user id from ctx
	accClaimIn, ok := ctx.Get("access_claim")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	var AccessClaim modelToken.AccessClaim
	err = crypto.ObjectMapper(accClaimIn, &AccessClaim)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	fmt.Println(AccessClaim.UserId)
	reqOrder.UserId, err = strconv.ParseUint(AccessClaim.UserId, 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	resOrder, err := ctrl.OrderSrv.SrvCreate(ctx, reqOrder)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, responses.Response{
		Code:    http.StatusCreated,
		Success: true,
		Message: responses.Success,
		Data:    resOrder,
	})

}
func (ctrl *CtrlOrdersImpl) GetQuery(ctx *gin.Context) (limit, offset uint64, err error) {
	limitQry, ok := ctx.GetQuery("limit")
	if !ok {
		limitQry = "10"
	}
	offsetQry, ok := ctx.GetQuery("page")
	if !ok {
		offsetQry = "0"
	}
	if offsetQry == "1" {
		offsetQry = "0"
	}
	limit, err = strconv.ParseUint(limitQry, 10, 64)
	if err != nil {
		return
	}

	offset_, err := strconv.ParseUint(offsetQry, 10, 64)
	if err != nil {
		return
	}

	if offset_ > 1 {

		offset_ *= 10
	}
	offset = offset_

	return
}
func (ctrl *CtrlOrdersImpl) ConvertId(ctx *gin.Context) (id uint64, err error) {
	reqId := ctx.Param("id")
	id, err = strconv.ParseUint(reqId, 10, 64)
	if err != nil {
		return
	}
	return
}

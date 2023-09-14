package controllers

import (
	"net/http"
	"strconv"

	modelsPayment "github.com/Billy278/pos_app_monolic/modules/models/payment"
	servicesPayment "github.com/Billy278/pos_app_monolic/modules/services/payment"
	"github.com/Billy278/pos_app_monolic/pkg/helper"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlPaymentImpl struct {
	PaymentSrv servicesPayment.PaymentSrv
	Validate   *validator.Validate
}

func NewCtrlPaymentImpl(patmentsrv servicesPayment.PaymentSrv, v *validator.Validate) CtrlPayment {
	return &CtrlPaymentImpl{
		PaymentSrv: patmentsrv,
		Validate:   v,
	}
}

func (ctrl *CtrlPaymentImpl) List(ctx *gin.Context) {
	limit, offset, err := ctrl.GetQuery(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	resPayment, err := ctrl.PaymentSrv.SrvList(ctx, limit, offset)
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
		Data:    resPayment,
	})

}
func (ctrl *CtrlPaymentImpl) FindByid(ctx *gin.Context) {
	idReq, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	resPayment, err := ctrl.PaymentSrv.SrvFindByid(ctx, idReq)
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
		Data:    resPayment,
	})
}
func (ctrl *CtrlPaymentImpl) Created(ctx *gin.Context) {
	reqPayment := modelsPayment.Payment{}
	err := ctx.ShouldBindJSON(&reqPayment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	//validasi request with validator
	err = ctrl.Validate.Struct(reqPayment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	//validasi type req
	err = helper.CekReqType(reqPayment.Type)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	resPayment, err := ctrl.PaymentSrv.SrvCreate(ctx, reqPayment)
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
		Data:    resPayment,
	})
}
func (ctrl *CtrlPaymentImpl) Updated(ctx *gin.Context) {
	idReq, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	reqPayment := modelsPayment.Payment{}
	err = ctx.ShouldBindJSON(&reqPayment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	//validasi req
	err = ctrl.Validate.Struct(reqPayment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	//validasi type req
	err = helper.CekReqType(reqPayment.Type)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	reqPayment.Id = idReq
	resPayment, err := ctrl.PaymentSrv.SrvUpdate(ctx, reqPayment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Response{
			Code:    http.StatusInternalServerError,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: responses.Success,
		Data:    resPayment,
	})

}
func (ctrl *CtrlPaymentImpl) Deleted(ctx *gin.Context) {
	reqId, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	err = ctrl.PaymentSrv.SrvDelete(ctx, reqId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, responses.Response{
		Code:    http.StatusOK,
		Success: true,
		Message: responses.Success,
	})
}
func (ctrl *CtrlPaymentImpl) GetQuery(ctx *gin.Context) (limit, offset uint64, err error) {
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
func (ctrl *CtrlPaymentImpl) ConvertId(ctx *gin.Context) (id uint64, err error) {
	reqId := ctx.Param("id")
	id, err = strconv.ParseUint(reqId, 10, 64)
	if err != nil {
		return
	}
	return
}

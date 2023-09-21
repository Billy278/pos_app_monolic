package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Billy278/pos_app_monolic/modules/models/products"
	sercives "github.com/Billy278/pos_app_monolic/modules/services/products"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlProductimpl struct {
	ProductSrv sercives.SrvProduct
	Validate   *validator.Validate
}

func NewCtrlProductimpl(proudctsrv sercives.SrvProduct, validate *validator.Validate) CtrlProduct {
	return &CtrlProductimpl{
		ProductSrv: proudctsrv,
		Validate:   validate,
	}
}

func (ctrl *CtrlProductimpl) List(ctx *gin.Context) {
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
	resProduct, err := ctrl.ProductSrv.SrvList(ctx, limit, offset)
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
		Data:    resProduct,
	})
}
func (ctrl *CtrlProductimpl) FindByid(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	resProduct, err := ctrl.ProductSrv.SrvFindByid(ctx, id)
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
		Data:    resProduct,
	})
}
func (ctrl *CtrlProductimpl) Created(ctx *gin.Context) {
	reqProduct := models.Products{}
	err := ctx.ShouldBindJSON(&reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	//validasi req
	err = ctrl.Validate.Struct(reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}

	resProduct, err := ctrl.ProductSrv.SrvCreate(ctx, reqProduct)
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
		Data:    resProduct,
	})
}
func (ctrl *CtrlProductimpl) Updated(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}

	reqProduct := models.Products{}
	err = ctx.ShouldBindJSON(&reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	err = ctrl.Validate.Struct(&reqProduct)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}
	reqProduct.Id = id
	resProduct, err := ctrl.ProductSrv.SrvUpdate(ctx, reqProduct)
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
		Data:    resProduct,
	})
}
func (ctrl *CtrlProductimpl) Deleted(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	err = ctrl.ProductSrv.SrvDelete(ctx, id)
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
func (ctrl *CtrlProductimpl) GetQuery(ctx *gin.Context) (limit, offset uint64, err error) {
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
func (ctrl *CtrlProductimpl) ConvertId(ctx *gin.Context) (id uint64, err error) {
	reqId := ctx.Param("id")
	id, err = strconv.ParseUint(reqId, 10, 64)
	if err != nil {
		return
	}
	return
}

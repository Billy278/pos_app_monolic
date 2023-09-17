package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Billy278/pos_app_monolic/modules/models/categories"
	services "github.com/Billy278/pos_app_monolic/modules/services/categories"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CtrlCategoriesimpl struct {
	CategoriesServ services.SrvCategories
	Validate       *validator.Validate
}

func NewCtrlCategoriesimpl(categorysrv services.SrvCategories, validate *validator.Validate) CtrlCategories {
	return &CtrlCategoriesimpl{
		CategoriesServ: categorysrv,
		Validate:       validate,
	}
}

func (ctrl *CtrlCategoriesimpl) List(ctx *gin.Context) {
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

	resCategories, err := ctrl.CategoriesServ.SrvList(ctx, limit, offset)
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
		Data:    resCategories,
	})
}
func (ctrl *CtrlCategoriesimpl) FindByid(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	resCategories, err := ctrl.CategoriesServ.SrvFindByid(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responses.Response{
			Code:    http.StatusNotFound,
			Success: false,
			Message: responses.NotFound,
		})
		return
	}
	ctx.JSON(http.StatusAccepted, responses.Response{
		Code:    http.StatusAccepted,
		Success: true,
		Message: responses.Success,
		Data:    resCategories,
	})
}
func (ctrl *CtrlCategoriesimpl) Created(ctx *gin.Context) {
	reqcategories := models.Categories{}
	err := ctx.ShouldBindJSON(&reqcategories)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	//validasi req
	err = ctrl.Validate.Struct(&reqcategories)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		})
		return
	}

	resCategories, err := ctrl.CategoriesServ.SrvCreate(ctx, reqcategories)
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
		Data:    resCategories,
	})
}
func (ctrl *CtrlCategoriesimpl) Updated(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}

	reqCategories := models.Categories{}
	err = ctx.ShouldBindJSON(&reqCategories)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}
	//validasi req

	err = ctrl.Validate.Struct(&reqCategories)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	reqCategories.Id = id
	resCategories, err := ctrl.CategoriesServ.SrvUpdate(ctx, reqCategories)
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
		Data:    resCategories,
	})
}
func (ctrl *CtrlCategoriesimpl) Deleted(ctx *gin.Context) {
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	err = ctrl.CategoriesServ.SrvDelete(ctx, id)
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
	})

}
func (ctrl *CtrlCategoriesimpl) GetQuery(ctx *gin.Context) (limit, offset uint64, err error) {
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
func (ctrl *CtrlCategoriesimpl) ConvertId(ctx *gin.Context) (id uint64, err error) {
	reqId := ctx.Param("id")
	id, err = strconv.ParseUint(reqId, 10, 64)
	if err != nil {
		return
	}
	return
}

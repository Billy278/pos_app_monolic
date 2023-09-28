package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Billy278/pos_app_monolic/modules/models/users"
	services "github.com/Billy278/pos_app_monolic/modules/services/users"
	"github.com/Billy278/pos_app_monolic/pkg/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCtrlimpl struct {
	UserSrv  services.UserSrv
	Validate *validator.Validate
}

func NewUserCtrlimpl(usersrv services.UserSrv, v *validator.Validate) UserCtrl {
	return &UserCtrlimpl{
		UserSrv:  usersrv,
		Validate: v,
	}
}

func (ctrl *UserCtrlimpl) List(ctx *gin.Context) {
	limit, offset, err := ctrl.GetQuery(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidQuery,
		})
		return
	}
	resUser, err := ctrl.UserSrv.SrvList(ctx, limit, offset)
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
		Data:    resUser,
	})
}
func (ctrl *UserCtrlimpl) FindByid(ctx *gin.Context) {
	// validasi param
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	resUser, err := ctrl.UserSrv.SrvFindByid(ctx, id)
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
		Data:    resUser,
	})
}
func (ctrl *UserCtrlimpl) Created(ctx *gin.Context) {
	//validasi req
	reqUser := models.User{}
	err := ctx.ShouldBindJSON(&reqUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	//validasi req with validator
	err = ctrl.Validate.Struct(reqUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}
	//cek apakah username sudah ada yg menggunakan
	err = ctrl.UserSrv.RepoFindUser(ctx, reqUser.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: "Username sudah digunakan",
		})
		return
	}
	resUser, err := ctrl.UserSrv.SrvCreate(ctx, reqUser)
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
		Data:    resUser,
	})
}
func (ctrl *UserCtrlimpl) Updated(ctx *gin.Context) {
	//validasi req
	reqUser := models.UserUpdated{}
	err := ctx.ShouldBindJSON(&reqUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	//validasi req with validator
	err = ctrl.Validate.Struct(reqUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidBody,
		})
		return
	}

	//validasi param
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}

	reqUser.Id = id
	resUser, err := ctrl.UserSrv.SrvUpdate(ctx, models.User{
		Id:    reqUser.Id,
		Name:  reqUser.Name,
		Email: reqUser.Email,
	})
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
		Data:    resUser,
	})
}
func (ctrl *UserCtrlimpl) Deleted(ctx *gin.Context) {
	//validasi param
	id, err := ctrl.ConvertId(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
			Success: false,
			Message: responses.InvalidParam,
		})
		return
	}
	err = ctrl.UserSrv.SrvDelete(ctx, id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responses.Response{
			Code:    http.StatusBadRequest,
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
func (ctrl *UserCtrlimpl) GetQuery(ctx *gin.Context) (limit, offset uint64, err error) {
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
func (ctrl *UserCtrlimpl) ConvertId(ctx *gin.Context) (id uint64, err error) {
	reqId := ctx.Param("id")
	id, err = strconv.ParseUint(reqId, 10, 64)
	if err != nil {
		return
	}
	return
}

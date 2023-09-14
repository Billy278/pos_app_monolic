package controllers

import "github.com/gin-gonic/gin"

type CtrlPayment interface {
	List(ctx *gin.Context)
	FindByid(ctx *gin.Context)
	Created(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Deleted(ctx *gin.Context)
}

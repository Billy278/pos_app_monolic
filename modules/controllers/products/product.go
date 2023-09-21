package controllers

import "github.com/gin-gonic/gin"

type CtrlProduct interface {
	List(ctx *gin.Context)
	FindByid(ctx *gin.Context)
	Created(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Deleted(ctx *gin.Context)
}

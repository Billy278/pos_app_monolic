package controllers

import "github.com/gin-gonic/gin"

type UserCtrl interface {
	List(ctx *gin.Context)
	FindByid(ctx *gin.Context)
	Register(ctx *gin.Context)
	Updated(ctx *gin.Context)
	Deleted(ctx *gin.Context)
	Login(ctx *gin.Context)
}

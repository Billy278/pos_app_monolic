package controllers

import "github.com/gin-gonic/gin"

type CtrlOrders interface {
	List(ctx *gin.Context)
	Detail(ctx *gin.Context)
	Add(ctx *gin.Context)
}

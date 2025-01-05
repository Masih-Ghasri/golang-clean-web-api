package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type HealthHandler struct{}

func (h HealthHandler) HealthGET(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "On Working GET!")
	return
}

func (h HealthHandler) HealthPOST(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "On Working POST!")
	return
}

func (h HealthHandler) HealthPOSTById(ctx *gin.Context){
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, fmt.Sprintf("On Working POST by id %s!",id))
	return
}

func NewHealthHandler() *HealthHandler{
	return &HealthHandler{}
}
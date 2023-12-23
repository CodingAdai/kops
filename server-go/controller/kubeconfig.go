package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wonderivan/logger"
	"net/http"
	"server-go/service"
)

var KubeConfig kubeconfig

type kubeconfig struct{}

type updateKubeConfigContentParam struct {
	Content string `json:"content"`
}

func (kc *kubeconfig) updateKubeConfigContent(ctx *gin.Context) {

	params := &updateKubeConfigContentParam{}

	if err := ctx.BindJSON(params); err != nil {
		logger.Error("bind json error, err: " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	data, err := service.KubeConfig.SetKubeConfigContent(params.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})

}

package http

import (
	"net/http"

	"github.com/Pdhenrique/GoNeoway/domain"
	"github.com/gin-gonic/gin"
)

func (handler *handler) getClient(context *gin.Context) {
	cpf := context.Param("cpf")

	client, err := handler.clientService.Get(cpf)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, client)
}

func (handler *handler) postClient(context *gin.Context) {
	client := &domain.Client{}
	if err := context.BindJSON(&client); err != nil {
		return
	}

	client, err := handler.clientService.Create(client)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusCreated, client)
}

func (handler *handler) putClient(context *gin.Context) {
	client := &domain.Client{}
	if err := context.BindJSON(&client); err != nil {
		return
	}

	err := handler.clientService.Update(client)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusNoContent, client)
}

func (handler *handler) deleteClient(context *gin.Context) {
	cpf := context.Param("cpf")

	err := handler.clientService.Delete(cpf)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.AbortWithStatus(http.StatusNoContent)
}

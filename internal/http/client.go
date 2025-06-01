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
	cpf := context.Param("cpf")
	client := &domain.Client{}

	if err := context.ShouldBindJSON(&client); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	client.CPF = cpf

	err := handler.clientService.Update(client)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, client)
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

// func (handler *handler) importClients(context *gin.Context) {
// 	var clients []*domain.Client
// 	if err := context.BindJSON(&clients); err != nil {
// 		context.AbortWithStatus(http.StatusBadRequest)
// 		return
// 	}

// 	err := handler.clientService.ImportClients(clients)
// 	if err != nil {
// 		context.AbortWithStatus(http.StatusInternalServerError)
// 		return
// 	}

// 	context.JSON(http.StatusCreated, clients)
// }

package http

import (
	"net/http"

	"github.com/Pdhenrique/GoNeoway/domain"
	"github.com/gin-gonic/gin"
)

type handler struct {
	clientService domain.ClientService
}

func NewHandler(clientService domain.ClientService) http.Handler {
	h := &handler{
		clientService: clientService,
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	v1 := router.Group("/v1")

	v1.GET("/clients/:cpf", h.getClient)
	v1.POST("/clients", h.postClient)
	v1.PUT("/clients/:cpf", h.putClient)
	v1.DELETE("/clients/:cpf", h.deleteClient)

	return router
}

package domain

import (
	"fmt"
	"time"
)

type FixedFloat float64

func (f FixedFloat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.2f", f)), nil
}

type Client struct {
	CPF                   string      `json:"cpf"`                  // 0 -> 18
	PRIVATE               int         `json:"private" `             // 18 -> 30
	INCOMPLETO            int         `json:"incompleto"`           // 30 -> 43
	DATA_ULTIMA_COMPRA    *time.Time  `json:"data_ultima_compra"`   // 43 -> 65
	TICKET_MEDIO          *FixedFloat `json:"ticket_medio"`         // 65 -> 87
	TICKET_ULTIMA_COMPRA  *FixedFloat `json:"ticket_ultima_compra"` // 87 -> 111
	LOJA_MAIS_FREQUENTADA string      `json:"loja_mais_frequente"`  // 111 -> 131
	LOJA_ULTIMA_COMPRA    string      `json:"loja_ultima_compra"`   // 131 -> 150
}

type ClientService interface {
	Get(cpf string) (*Client, error)
	Update(client *Client) error
	Create(client *Client) (*Client, error)
	Delete(cpf string) error
}

type ClientStorage interface {
	Insert(client *Client) (*Client, error)
	FindByCPF(cpf string) (*Client, error)
	Update(client *Client) error
	Delete(cpf string) error

	ImportClients(clients []*Client) error
}

func NewClient(
	cpf string,
	private, incompleto int,
	dataUltimaCompra *time.Time,
	ticketMedio, ticketUltimaCompra *FixedFloat,
	lojaMaisFrequentada, lojaUltimaCompra string) *Client {
	return &Client{
		CPF:                   cpf,
		PRIVATE:               private,
		INCOMPLETO:            incompleto,
		DATA_ULTIMA_COMPRA:    dataUltimaCompra,
		TICKET_MEDIO:          ticketMedio,
		TICKET_ULTIMA_COMPRA:  ticketUltimaCompra,
		LOJA_MAIS_FREQUENTADA: lojaMaisFrequentada,
		LOJA_ULTIMA_COMPRA:    lojaUltimaCompra,
	}
}

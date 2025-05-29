package model

import "time"

type Client struct {
	CPF                   string     // 0 -> 18
	PRIVATE               int        // 18 -> 30
	INCOMPLETO            int        // 30 -> 43
	DATA_ULTIMA_COMPRA    *time.Time //43 -> 65
	TICKET_MEDIO          *float64   // 65 -> 87
	TICKET_ULTIMA_COMPRA  *float64   // 87 -> 111
	LOJA_MAIS_FREQUENTADA string     // 111 -> 131
	LOJA_ULTIMA_COMPRA    string     // 131 -> 150
}

package repository

import (
	"database/sql"
	"github.com/Pdhenrique/GoNeoway/pkg/model"
)

func SaveClients(db *sql.DB, clients []model.Client) error {
		query := `INSERT INTO clients (
		cpf, private, incompleto, data_ultima_compra, 
		ticket_medio, ticket_ultima_compra, loja_mais_frequentada, loja_ultima_compra
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	
		for _, client := range clients{
			_, err := db.Exec(query,
				client.CPF,
				client.PRIVATE,
				client.INCOMPLETO,
				client.DATA_ULTIMA_COMPRA,
				client.TICKET_MEDIO,
				client.TICKET_ULTIMA_COMPRA,
				client.LOJA_MAIS_FREQUENTADA,
				client.LOJA_ULTIMA_COMPRA)
			if err != nil {
				return err
			}
		}

	return nil
}


package db

import (
	"database/sql"
	"log"

	"github.com/Pdhenrique/GoNeoway/domain"
)

type clientStorage struct {
	DB *sql.DB
}

func NewClientStorage(db *sql.DB) domain.ClientStorage {
	return &clientStorage{
		DB: db,
	}
}

func (c *clientStorage) Delete(cpf string) error {
	_, err := c.DB.Exec(
		`DELETE FROM clients WHERE cpf = $1`, cpf)

	if err != nil {
		log.Printf("error deleting client with CPF %s: %v", cpf, err)
	}

	return err
}

func (c *clientStorage) FindByCPF(cpf string) (*domain.Client, error) {
	var client domain.Client

	err := c.DB.QueryRow(
		`SELECT cpf, private, incompleto, data_ultima_compra, 
		ticket_medio, ticket_ultima_compra, loja_mais_frequentada, loja_ultima_compra 
		FROM clients WHERE cpf = $1`, cpf).Scan(
		&client.CPF,
		&client.PRIVATE,
		&client.INCOMPLETO,
		&client.DATA_ULTIMA_COMPRA,
		&client.TICKET_MEDIO,
		&client.TICKET_ULTIMA_COMPRA,
		&client.LOJA_MAIS_FREQUENTADA,
		&client.LOJA_ULTIMA_COMPRA)

	if err != nil {
		log.Printf("error finding client with CPF %s: %v", cpf, err)
		return nil, err
	}

	return &client, nil
}

func (c *clientStorage) Insert(client *domain.Client) (*domain.Client, error) {
	err := c.DB.QueryRow(
		`INSERT INTO clients (
		cpf, 
		private, 
		incompleto, 
		data_ultima_compra, 
		ticket_medio, 
		ticket_ultima_compra, 
		loja_mais_frequentada, 
		loja_ultima_compra) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING cpf`,
		client.CPF,
		client.PRIVATE,
		client.INCOMPLETO,
		client.DATA_ULTIMA_COMPRA,
		client.TICKET_MEDIO,
		client.TICKET_ULTIMA_COMPRA,
		client.LOJA_MAIS_FREQUENTADA,
		client.LOJA_ULTIMA_COMPRA).Scan(&client.CPF)

	if err != nil {
		log.Printf("error inserting client with CPF %s: %v", client.CPF, err)
	}
	return client, nil
}

func (c *clientStorage) Update(client *domain.Client) error {
	_, err := c.DB.Exec(
		`UPDATE clients SET 
		private = $1, 
		incompleto = $2, 
		data_ultima_compra = $3, 
		ticket_medio = $4, 
		ticket_ultima_compra = $5, 
		loja_mais_frequentada = $6, 
		loja_ultima_compra = $7 
		WHERE cpf = $8`,
		client.PRIVATE,
		client.INCOMPLETO,
		client.DATA_ULTIMA_COMPRA,
		client.TICKET_MEDIO,
		client.TICKET_ULTIMA_COMPRA,
		client.LOJA_MAIS_FREQUENTADA,
		client.LOJA_ULTIMA_COMPRA,
		client.CPF)

	if err != nil {
		log.Printf("error updating client with CPF %s: %v", client.CPF, err)
		return err
	}

	return err
}

func (c *clientStorage) ImportClients(clients []*domain.Client) error {
	if len(clients) == 0 {
		return nil
	}

	query := `INSERT INTO clients (
		cpf, private, incompleto, data_ultima_compra, 
		ticket_medio, ticket_ultima_compra, loja_mais_frequentada, loja_ultima_compra
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	tx, err := c.DB.Begin()
	if err != nil {
		log.Printf("error starting transaction: %v", err)
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Printf("error preparing statement: %v", err)
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for _, client := range clients {
		_, err := stmt.Exec(
			client.CPF,
			client.PRIVATE,
			client.INCOMPLETO,
			client.DATA_ULTIMA_COMPRA,
			client.TICKET_MEDIO,
			client.TICKET_ULTIMA_COMPRA,
			client.LOJA_MAIS_FREQUENTADA,
			client.LOJA_ULTIMA_COMPRA)
		if err != nil {
			log.Printf("error inserting client with CPF %s: %v", client.CPF, err)
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

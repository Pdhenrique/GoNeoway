package main

import (
	"fmt"
	"log"

	"github.com/Pdhenrique/GoNeoway/internal/db"
	"github.com/Pdhenrique/GoNeoway/pkg/parser"
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("ERROR connecting to the db", err)
	}
	defer conn.Close()
	fmt.Println("API running and successfuly connected!")

	clients, err := parser.Parse()
	if err != nil {
		log.Fatal("erro ao realizar parse do arquivo", err)
	}

	fmt.Println("Total de arquivos lidos", len(clients))
}

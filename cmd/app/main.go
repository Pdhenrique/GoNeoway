package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Pdhenrique/GoNeoway/internal/db"
	"github.com/Pdhenrique/GoNeoway/pkg/parser"
	"github.com/Pdhenrique/GoNeoway/pkg/sanitizer"
	"github.com/Pdhenrique/GoNeoway/pkg/repository"
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("ERROR connecting to the db", err)
	}
	defer conn.Close()
	fmt.Println("API running and successfuly connected!")

	file, _ := os.Open("base_teste.txt")

	clients, err := parser.Parse(file)
	if err != nil {
		log.Fatal("erro ao realizar parse do arquivo", err)
	}

	sanitized, err := sanitizer.Sanitize(clients)
	if err != nil {
		log.Fatal("erro ao realizar limpeza dos valores", err)
	}

	err = repository.SaveClients(conn, sanitized)
	if err != nil {
		log.Fatal("Erro ao persistir no banco", err)
	}

	fmt.Println("Importação finalizada!")

	fmt.Printf("Clientes validos após sanitização: %d\n", len(sanitized))

}

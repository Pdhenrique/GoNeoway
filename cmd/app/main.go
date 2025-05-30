package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Pdhenrique/GoNeoway/internal/db"
	"github.com/Pdhenrique/GoNeoway/pkg/parser"
	"github.com/Pdhenrique/GoNeoway/pkg/sanitizer"
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("ERROR connecting to the db", err)
	}
	defer conn.Close()
	fmt.Println("API running and successfuly connected!")

	file, _ := os.Open("base_teste_2.txt")

	clients, err := parser.Parse(file)
	if err != nil {
		log.Fatal("erro ao realizar parse do arquivo", err)
	}

	sanitized, err := sanitizer.Sanitize(clients)
	if err != nil {
		log.Fatal("erro ao realizar limpeza dos valores", err)
	}

	fmt.Println("sanitized", sanitized)

	http.ListenAndServe(":8080", nil)

}

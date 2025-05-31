package main

import (
	"log"
	"os"

	"github.com/Pdhenrique/GoNeoway/internal/db"
	"github.com/Pdhenrique/GoNeoway/internal/http"
	"github.com/Pdhenrique/GoNeoway/pkg/client"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("env DATABASE_URL não definido")
	}

	conn, err := db.Connect(dbURL)
	if err != nil {
		log.Fatal("ERROR conectando ao db:", err)
	}
	defer conn.Close()

	clientStorage := db.NewClientStorage(conn)
	clientService := client.NewService(clientStorage)
	handler := http.NewHandler(clientService)

	server := http.NewServer(handler, "8080")
	server.Start()
	defer server.Stop()

}

// file, _ := os.Open("base_teste.txt")

// clients, err := parser.Parse(file)
// if err != nil {
// 	log.Fatal("erro ao realizar parse do arquivo", err)
// }

// sanitized, err := sanitizer.Sanitize(clients)
// if err != nil {
// 	log.Fatal("erro ao realizar limpeza dos valores", err)
// }

// err = db.SaveClients(conn, sanitized)
// if err != nil {
// 	log.Fatal("Erro ao persistir no banco", err)
// }

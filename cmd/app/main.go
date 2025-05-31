package main

import (
	"log"
	"os"

	"github.com/Pdhenrique/GoNeoway/internal/db"
	"github.com/Pdhenrique/GoNeoway/internal/http"
	"github.com/Pdhenrique/GoNeoway/pkg/client"
	"github.com/Pdhenrique/GoNeoway/pkg/importer"
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

	importer := importer.New(clientStorage)
	if err := importer.ImportFromFile("base_teste.txt"); err != nil {
		log.Fatal("Erro ao importar dados:", err)
	}

	server := http.NewServer(handler, "8080")
	server.Start()
	defer server.Stop()

	select {}
}

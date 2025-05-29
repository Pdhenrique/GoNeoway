package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pdhenrique/GoNeoway/internal/db"
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("ERROR connecting to the db", err)
	}
	defer conn.Close()
	fmt.Println("API running and successfuly connected!")

	// clients, err := parser.Parse()
	// if err != nil {
	// 	log.Fatal("erro ao realizar parse do arquivo", err)
	// }

	http.ListenAndServe(":8080", nil)

}

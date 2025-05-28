package main

import (
	"fmt"
	"log"

	"github.com/Pdhenrique/GoNeoway/internal/db"
)

func main() {

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("ERROR connecting to the db", err)
	}
	defer conn.Close()

	fmt.Println("API running and successfuly connected!")

}

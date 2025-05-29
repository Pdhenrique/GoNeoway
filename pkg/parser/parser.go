package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func Parse() ([]Client, error) {
	var clients []Client

	file, err := os.Open("base_teste.txt")
	if err != nil {
		log.Fatalf("Erro ao abrir txt", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		fmt.Println("Pulando cabeçalho:", scanner.Text())
	} else {
		fmt.Errorf("arquivo vazio ou erro na leitura")
	}

	for scanner.Scan() {
		linha := scanner.Text()
		fmt.Println(len(linha))

		cpf := strings.TrimSpace(linha[0:18])
		privateStr := strings.TrimSpace(linha[18:30])
		incompletoStr := strings.TrimSpace(linha[30:43])
		dataCompraStr := strings.TrimSpace(linha[43:65])
		ticketMedioStr := strings.TrimSpace(linha[65:87])
		ticketUltimaStr := strings.TrimSpace(linha[87:111])
		lojaMaisFreq := strings.TrimSpace(linha[111:131])
		lojaUltima := strings.TrimSpace(linha[131:150])

		private, err := strconv.Atoi(privateStr)
		if err != nil {
			private = 0
		}

		incompleto, err := strconv.Atoi(incompletoStr)
		if err != nil {
			incompleto = 0
		}

		dataCompra := parseNullableDate(dataCompraStr)
		ticketMedio := parseNullableFloat(ticketMedioStr)
		ticketUltima := parseNullableFloat(ticketUltimaStr)

		clients = append(clients, Client{
			CPF:                   cpf,
			PRIVATE:               private,
			INCOMPLETO:            incompleto,
			DATA_ULTIMA_COMPRA:    dataCompra,
			TICKET_MEDIO:          ticketMedio,
			TICKET_ULTIMA_COMPRA:  ticketUltima,
			LOJA_MAIS_FREQUENTADA: lojaMaisFreq,
			LOJA_ULTIMA_COMPRA:    lojaUltima})

		fmt.Printf("CPF: %s | PRIVATE: %s | INCOMPLETO: %s | MAISFREQ: %s | ULTIMA: %s\n", cpf, privateStr, incompletoStr, lojaMaisFreq, lojaUltima)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("erro na linha: %v", err)
	}

	return clients, nil
}

func parseNullableFloat(s string) *float64 {
	if s == "" || strings.ToUpper(s) == "NULL" {
		return nil
	}
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &f
}

func parseNullableDate(s string) *time.Time {
	if s == "" || strings.ToUpper(s) == "NULL" {
		return nil
	}
	layout := "2006-01-02" // ajuste se necessário
	t, err := time.Parse(layout, s)
	if err != nil {
		return nil
	}
	return &t
}

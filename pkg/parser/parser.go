package parser

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Pdhenrique/GoNeoway/pkg/model"
)

func Parse(reader io.Reader) ([]model.Client, error) {
	var clients []model.Client

	scanner := bufio.NewScanner(reader)

	if scanner.Scan() {
		fmt.Println("Pulando cabeçalho:", scanner.Text())
	} else {
		fmt.Errorf("arquivo vazio ou erro na leitura")
	}

	for scanner.Scan() {
		linha := scanner.Text()

		// Define tamanho do campo em colunas.
		cpf := strings.TrimSpace(safeSlice(linha, 0, 18))
		privateStr := strings.TrimSpace(safeSlice(linha, 18, 30))
		incompletoStr := strings.TrimSpace(safeSlice(linha, 30, 43))
		dataCompraStr := strings.TrimSpace(safeSlice(linha, 43, 65))
		ticketMedioStr := strings.TrimSpace(safeSlice(linha, 65, 87))
		ticketUltimaStr := strings.TrimSpace(safeSlice(linha, 87, 111))
		lojaMaisFreq := strings.TrimSpace(safeSlice(linha, 111, 131))
		lojaUltima := strings.TrimSpace(safeSlice(linha, 131, 150))

		// Transformar String em Int
		private, err := strconv.Atoi(privateStr)
		if err != nil {
			private = 0
		}

		// Transformar String em Int
		incompleto, err := strconv.Atoi(incompletoStr)
		if err != nil {
			incompleto = 0
		}

		// Transforma String em data e valida caso seja null
		dataCompra := parseNullableDate(dataCompraStr)

		// Transforma String em Float e valida caso seja null
		ticketMedio := parseNullableFloat(ticketMedioStr)
		ticketUltima := parseNullableFloat(ticketUltimaStr)

		clients = append(clients, model.Client{
			CPF:                   cpf,
			PRIVATE:               private,
			INCOMPLETO:            incompleto,
			DATA_ULTIMA_COMPRA:    dataCompra,
			TICKET_MEDIO:          ticketMedio,
			TICKET_ULTIMA_COMPRA:  ticketUltima,
			LOJA_MAIS_FREQUENTADA: lojaMaisFreq,
			LOJA_ULTIMA_COMPRA:    lojaUltima})

		fmt.Printf("%+v\n", clients[len(clients)-1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("erro na linha: %v", err)
	}

	return clients, nil
}

// Transforma String em Float e valida caso seja null
func parseNullableFloat(arg string) *float64 {
	if arg == "" || strings.ToUpper(arg) == "NULL" {
		return nil
	}
	float, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		return nil
	}
	return &float
}

// Transforma String em data e valida caso seja null
func parseNullableDate(arg string) *time.Time {
	if arg == "" || strings.ToUpper(arg) == "NULL" {
		return nil
	}
	layout := "2006-01-02"
	time, err := time.Parse(layout, arg)
	if err != nil {
		return nil
	}
	return &time
}

// Verifica tamanho da
func safeSlice(arg string, start int, end int) string {
	if len(arg) < start {
		return ""
	}

	if len(arg) < end {
		end = len(arg)
	}

	return arg[start:end]
}

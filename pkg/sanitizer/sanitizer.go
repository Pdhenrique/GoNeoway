package sanitizer

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/Pdhenrique/GoNeoway/pkg/model"
)

func Sanitize(client []model.Client) ([]model.Client, error) {

	for i := range client {
		client[i].LOJA_MAIS_FREQUENTADA = sanitizeText(client[i].LOJA_MAIS_FREQUENTADA)
		client[i].LOJA_ULTIMA_COMPRA = sanitizeText(client[i].LOJA_ULTIMA_COMPRA)

		cpf := client[i].CPF
		if !validateCPF(cpf) {
			fmt.Printf("⚠️ CPF inválido: %s\n", client[i].CPF)
		}
		client[i].CPF = unmaskCPF(cpf)

		lojaUltimaCompra := client[i].LOJA_ULTIMA_COMPRA
		if lojaUltimaCompra != "" && lojaUltimaCompra != "NULL" {
			if !validateCNPJ(lojaUltimaCompra) {
				fmt.Printf("⚠️ CNPJ inválido: %s\n", lojaUltimaCompra)
			}
		}
		client[i].LOJA_ULTIMA_COMPRA = unmaskCNPJ(lojaUltimaCompra)

		lojaMaisFrequentada := client[i].LOJA_MAIS_FREQUENTADA
		if lojaMaisFrequentada != "" && lojaMaisFrequentada != "NULL" {
			if !validateCNPJ(lojaMaisFrequentada) {
				fmt.Printf("⚠️ CNPJ inválido: %s\n", lojaMaisFrequentada)
			}
		}
		client[i].LOJA_MAIS_FREQUENTADA = unmaskCNPJ(lojaMaisFrequentada)

	}

	return client, nil

}

func unmaskCPF(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	return cpf
}

func validateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	// CNPJs repetidos não são válidos
	invalids := []string{
		"00000000000", "11111111111", "22222222222",
		"33333333333", "44444444444", "55555555555",
		"66666666666", "77777777777", "88888888888",
		"99999999999",
	}
	for _, inv := range invalids {
		if cpf == inv {
			return false
		}
	}

	// cálculo do primeiro dígito
	sum := 0
	for i := 0; i < 9; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (10 - i)
	}
	d1 := 11 - (sum % 11)
	if d1 >= 10 {
		d1 = 0
	}
	if d1 != int(cpf[9]-'0') {
		return false
	}

	// cálculo do segundo dígito
	sum = 0
	for i := 0; i < 10; i++ {
		digit := int(cpf[i] - '0')
		sum += digit * (11 - i)
	}
	d2 := 11 - (sum % 11)
	if d2 >= 10 {
		d2 = 0
	}
	return d2 == int(cpf[10]-'0')
}

func unmaskCNPJ(cnpj string) string {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	return cnpj
}

func validateCNPJ(cnpj string) bool {
	// Remove máscara
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")

	if len(cnpj) != 14 {
		return false
	}

	// CNPJs repetidos não são válidos
	invalids := []string{
		"00000000000000", "11111111111111", "22222222222222",
		"33333333333333", "44444444444444", "55555555555555",
		"66666666666666", "77777777777777", "88888888888888",
		"99999999999999",
	}
	for _, inv := range invalids {
		if cnpj == inv {
			return false
		}
	}

	// Cálculo dos dígitos verificadores
	var calcDigit = func(cnpj string, multipliers []int) int {
		sum := 0
		for i, m := range multipliers {
			d := int(cnpj[i] - '0')
			sum += d * m
		}
		remainder := sum % 11
		if remainder < 2 {
			return 0
		}
		return 11 - remainder
	}

	// Multiplicadores para 1º e 2º dígito
	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	d1 := calcDigit(cnpj, multipliers1)
	d2 := calcDigit(cnpj+fmt.Sprint(d1), multipliers2)

	return cnpj[12] == byte(d1+'0') && cnpj[13] == byte(d2+'0')
}

func removeAccents(input string) string {
	t := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	result, _, _ := transform.String(t, input)
	return result
}

func sanitizeText(input string) string {
	input = strings.TrimSpace(input)
	input = removeAccents(input)
	input = strings.ToUpper(input)
	return input
}

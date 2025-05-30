package sanitizer

import (
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
		client[i].CPF = unmaskCPF(cpf)

		lojaUltimaCompra := client[i].LOJA_ULTIMA_COMPRA
		client[i].LOJA_ULTIMA_COMPRA = unmaskCNPJ(lojaUltimaCompra)

		lojaMaisFrequentada := client[i].LOJA_MAIS_FREQUENTADA
		client[i].LOJA_MAIS_FREQUENTADA = unmaskCNPJ(lojaMaisFrequentada)

	}

	return client, nil

}

func unmaskCPF(cpf string) string {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	return cpf
}

func unmaskCNPJ(cnpj string) string {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	return cnpj
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

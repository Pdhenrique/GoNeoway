package importer

import (
	"fmt"
	"os"

	"github.com/Pdhenrique/GoNeoway/domain"
	"github.com/Pdhenrique/GoNeoway/pkg/parser"
	"github.com/Pdhenrique/GoNeoway/pkg/sanitizer"
)

type Service struct {
	storage domain.ClientStorage
}

func New(storage domain.ClientStorage) *Service {
	return &Service{storage: storage}
}

func (s *Service) ImportFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}
	defer file.Close()

	clients, err := parser.Parse(file)
	if err != nil {
		return fmt.Errorf("erro no parser: %w", err)
	}

	sanitized, err := sanitizer.Sanitize(clients)
	if err != nil {
		return fmt.Errorf("erro na sanitização: %w", err)
	}

	return s.storage.ImportClients(pointerSlice(sanitized))
}

func pointerSlice(clients []domain.Client) []*domain.Client {
	result := make([]*domain.Client, len(clients))
	for i := range clients {
		result[i] = &clients[i]
	}
	return result
}

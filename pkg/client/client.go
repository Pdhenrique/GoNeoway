package client

import "github.com/Pdhenrique/GoNeoway/domain"

type service struct {
	clientStorage domain.ClientStorage
}

func NewService(clientStorage domain.ClientStorage) *service {
	return &service{
		clientStorage: clientStorage,
	}
}

// Create implements domain.ClientService.
func (s *service) Create(client *domain.Client) (*domain.Client, error) {
	return s.clientStorage.Insert(client)
}

// Delete implements domain.ClientService.
func (s *service) Delete(cpf string) error {
	return s.clientStorage.Delete(cpf)
}

// Get implements domain.ClientService.
func (s *service) Get(cpf string) (*domain.Client, error) {
	return s.clientStorage.FindByCPF(cpf)
}

// Update implements domain.ClientService.
func (s *service) Update(client *domain.Client) error {
	return s.clientStorage.Update(client)
}

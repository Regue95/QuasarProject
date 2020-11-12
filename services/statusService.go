package services

import (
	"quasarproject/repositories"
	"quasarproject/responses"
)

// StatusServiceInterface is an interface
type StatusServiceInterface interface {
	GetPing() (responses.Status, error)
	PingDB() error
}

type statusService struct {
	repository repositories.RepositoryInterface
}

// NewStatusService implements StatusServiceInterface
func NewStatusService(repo repositories.RepositoryInterface) StatusServiceInterface {
	return &statusService{
		repository: repo,
	}
}

func (s statusService) GetPing() (responses.Status, error) {
	return responses.Status{Status: "OK"}, nil
}

func (s statusService) PingDB() error {
	err := s.repository.Ping()
	return s.checkError(err)
}

func (s statusService) checkError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

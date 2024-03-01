package service

import (
	"task1/pkg/repository"
)

type GetResults interface {
	GetResults(data []int64) ([]int64, error)
}

type Service struct {
	GetResults
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		GetResults: NewGetResultsSevice(repository.GetResults),
	}
}

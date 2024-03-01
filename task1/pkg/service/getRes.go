package service

import (
	"task1/pkg/repository"
)

type GetResultsSevice struct {
	repository repository.GetResults
}

func NewGetResultsSevice(repo repository.GetResults) *GetResultsSevice {
	return &GetResultsSevice{repository: repo}
}

func (s *GetResultsSevice) GetResults(data []int64) ([]int64, error) {
	return s.repository.GetResults(data)
}

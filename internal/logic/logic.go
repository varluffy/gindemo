package logic

import "github.com/varluffy/gindemo/internal/repository"

type Logic struct {
}

func NewLogic(repo *repository.Repository) *Logic {
	return &Logic{}
}

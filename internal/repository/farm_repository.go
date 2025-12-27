package repository

import "looker.com/neutral-farming/internal/domain"

type FarmRepository interface {
	FindByID(ID uint) (*domain.Farm, error)
}

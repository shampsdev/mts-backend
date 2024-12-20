package adapter

import "api.mts.shamps.dev/internal/domain"

type Adapter interface {
	Parse(data []byte) (*domain.Person, error)
	GetAll(data []byte) ([]*domain.Person, error)
}

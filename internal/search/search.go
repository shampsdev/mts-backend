package search

import "api.mts.shamps.dev/internal/domain"

type Filter struct {
	Key string
	Val string
}

type Engine interface {
	GetAll() []*domain.PersonNode
	Filter(filters []Filter) []*domain.PersonNode
	Search(text string) []*domain.PersonNode
}

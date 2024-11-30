package search

import "api.mts.shamps.dev/internal/domain"

type Filter struct {
	Key string
	Val string
}

type Engine interface {
	AllPersons() []*domain.Person
	SearchPersons(text string, filters []Filter) []*domain.Person
}

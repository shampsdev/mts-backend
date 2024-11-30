package search

import (
	"strings"

	"api.mts.shamps.dev/internal/domain"
	"api.mts.shamps.dev/internal/search/data"
)

type JSONEngine struct {
	persons []*domain.Person
}

func NewJSONEngine() *JSONEngine {
	return &JSONEngine{persons: data.LoadPersons()}
}

func (e *JSONEngine) AllPersons() []*domain.Person {
	return e.persons
}

func (e *JSONEngine) SearchPersons(text string, filters []Filter) []*domain.Person {
	var results []*domain.Person
	for _, person := range e.persons {
		if strings.Contains(person.Name, text) || strings.Contains(person.JobTitle, text) {
			results = append(results, person)
		}
	}
	return results
}

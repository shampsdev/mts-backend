package search

import (
	"strings"

	"api.mts.shamps.dev/internal/domain"
)

type JSONEngine struct {
	data []*domain.PersonNode
}

func NewJSONEngine() *JSONEngine {
	return &JSONEngine{data: data}
}

func (e *JSONEngine) GetAll() []*domain.PersonNode {
	return e.data
}

func (e *JSONEngine) Filter(filters []Filter) []*domain.PersonNode {
	var results []*domain.PersonNode
	for _, person := range e.data {
		matches := true
		for _, filter := range filters {
			switch filter.Key {
			case "id":
				if person.ID != filter.Val {
					matches = false
				}
			case "name":
				if person.Name != filter.Val {
					matches = false
				}
			case "status":
				if person.Status != filter.Val {
					matches = false
				}
			// Добавьте другие фильтры, если необходимо
			default:
				matches = false
			}
			if !matches {
				break
			}
		}
		if matches {
			results = append(results, person)
		}
	}
	return results
}

func (e *JSONEngine) Search(text string) []*domain.PersonNode {
	var results []*domain.PersonNode
	for _, person := range e.data {
		if strings.Contains(person.Name, text) || strings.Contains(person.JobTitle, text) {
			results = append(results, person)
		}
	}
	return results
}

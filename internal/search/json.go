package search

import (
	"log"

	"api.mts.shamps.dev/external/adapter"
	"api.mts.shamps.dev/internal/domain"
	"github.com/blevesearch/bleve/v2"
)

type BleveEngine struct {
	persons map[string]*domain.Person
	index   bleve.Index
}

func NewBleveEngine(a adapter.Adapter) *BleveEngine {
	personsSlice := loadData(a)

	persons := make(map[string]*domain.Person)
	for _, person := range personsSlice {
		persons[person.ID] = person
	}

	indexMapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(indexMapping)
	if err != nil {
		log.Fatalf("Error creating index: %v", err)
	}

	for _, person := range persons {
		err := index.Index(person.ID, person)
		if err != nil {
			log.Fatalf("Error indexing person: %v", err)
		}
	}

	return &BleveEngine{
		persons: persons,
		index:   index,
	}
}

func (e *BleveEngine) AllPersons() []*domain.Person {
	var result []*domain.Person
	for _, person := range e.persons {
		result = append(result, person)
	}
	return result
}

func (e *BleveEngine) SearchPersons(text string, filters []Filter) []*domain.Person {
	query := bleve.NewQueryStringQuery(text)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, err := e.index.Search(searchRequest)
	if err != nil {
		log.Printf("Error searching for persons: %v", err)
		return nil
	}

	results := make([]*domain.Person, 0, len(searchResult.Hits))
	for _, hit := range searchResult.Hits {
		if person, exists := e.persons[hit.ID]; exists {
			results = append(results, person)
		}
	}

	return results
}

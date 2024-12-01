package search

import (
	"errors"
	"log"

	"api.mts.shamps.dev/external/adapter"
	"api.mts.shamps.dev/internal/domain"
	"github.com/blevesearch/bleve/v2"
	"golang.org/x/exp/slices"
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

	for _, person := range persons {
		if person.Head != nil {
			parent, exists := persons[*person.Head]
			if exists {
				parent.Children = append(parent.Children, person.ID)
			}
		}
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

var ErrNotFound = errors.New("person not found")

func (e *BleveEngine) NodeByID(id string) (*domain.PersonNode, error) {
	person, exists := e.persons[id]
	if !exists {
		return nil, ErrNotFound
	}

	return domain.PersonToNode(person), nil
}

func (e *BleveEngine) FindPathByIDs(from, to string) ([]*domain.PersonNode, error) {
	lca, err := e.lessCommonAncestor(from, to)
	if err != nil {
		return nil, err
	}

	pathFromLca := make([]*domain.PersonNode, 0, 2)
	cur := from
	for cur != lca {
		pathFromLca = append(pathFromLca, domain.PersonToNode(e.persons[cur]))
		cur = *e.persons[cur].Head
	}

	pathToLca := make([]*domain.PersonNode, 0, 2)
	cur = to
	for cur != lca {
		pathToLca = append(pathToLca, domain.PersonToNode(e.persons[cur]))
		cur = *e.persons[cur].Head
	}

	path := make([]*domain.PersonNode, 0, len(pathFromLca)+len(pathToLca)+1)
	path = append(path, pathFromLca...)
	path = append(path, domain.PersonToNode(e.persons[lca]))
	slices.Reverse(pathToLca)
	path = append(path, pathToLca...)

	return path, nil
}

func (e *BleveEngine) personHeight(id string) (int, error) {
	person, exists := e.persons[id]
	if !exists {
		return 0, ErrNotFound
	}

	height := 0
	parent := person.Head
	for parent != nil {
		height++
		p, exists := e.persons[*parent]
		if !exists {
			return 0, ErrNotFound
		}
		parent = p.Head
	}

	return height, nil
}

func (e *BleveEngine) lessCommonAncestor(from, to string) (string, error) {
	hFrom, err := e.personHeight(from)
	if err != nil {
		return "", err
	}

	hTo, err := e.personHeight(to)
	if err != nil {
		return "", err
	}

	for hFrom != hTo {
		if hFrom > hTo {
			from = *e.persons[from].Head
			hFrom--
		} else {
			to = *e.persons[to].Head
			hTo--
		}
	}

	for from != to {
		from = *e.persons[from].Head
		to = *e.persons[to].Head
	}

	return from, nil
}

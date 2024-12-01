package search

import (
	"errors"
	"log"

	"api.mts.shamps.dev/external/adapter"
	"api.mts.shamps.dev/internal/domain"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/search/query"
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

func (e *BleveEngine) SearchPersons(text string, filters []Filter) ([]*domain.Person, error) {
	prefixQuery := bleve.NewPrefixQuery(text)
	fuzzyQuery := bleve.NewFuzzyQuery(text)
	fuzzyQuery.Fuzziness = 2
	translitQuery := bleve.NewFuzzyQuery(transliterate(text))
	translitQuery.Fuzziness = 2

	fieldQuery := bleve.NewMatchQuery(text)

	query := bleve.NewDisjunctionQuery(prefixQuery, fuzzyQuery, translitQuery, fieldQuery)

	persons, err := e.findPersons(query)
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (e *BleveEngine) findPersons(q query.Query) ([]*domain.Person, error) {
	sr := bleve.NewSearchRequest(q)
	searchResult, err := e.index.Search(sr)
	if err != nil {
		log.Printf("Error searching for persons: %v", err)
		return nil, err
	}
	return e.hitsToPersons(searchResult), nil
}

func (e *BleveEngine) hitsToPersons(sr *bleve.SearchResult) []*domain.Person {
	persons := make([]*domain.Person, 0, len(sr.Hits))
	for _, hit := range sr.Hits {
		if person, exists := e.persons[hit.ID]; exists {
			persons = append(persons, person)
		}
	}
	return persons
}

var ErrNotFound = errors.New("person not found")

func (e *BleveEngine) NodeByID(id string) (*domain.PersonNode, error) {
	log.Printf("NodeByID called with ID: %s", id)

	person, exists := e.persons[id]
	if !exists {
		log.Printf("Person with ID %s not found", id)
		return nil, ErrNotFound
	}

	return domain.PersonToNode(person), nil
}

func (e *BleveEngine) PersonById(id string) (*domain.Person, error) {
	log.Printf("PersonByID called with ID: %s", id)

	person, exists := e.persons[id]
	if !exists {
		log.Printf("Person with ID %s not found", id)
		return nil, ErrNotFound
	}

	return person, nil
}

func (e *BleveEngine) AllDepartments() []string {
	departments := make(map[string]struct{})
	for _, person := range e.persons {
		if person.Department != "" {
			departments[person.Department] = struct{}{}
		}
	}

	result := make([]string, 0, len(departments))
	for department := range departments {
		result = append(result, department)
	}
	return result
}

func (e *BleveEngine) AllDivisions() []string {
	divisions := make(map[string]struct{})
	for _, person := range e.persons {
		if person.Division != nil && *person.Division != "" {
			divisions[*person.Division] = struct{}{}
		}
	}

	result := make([]string, 0, len(divisions))
	for division := range divisions {
		result = append(result, division)
	}
	return result
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
	log.Printf("Calculating height for person ID: %s", id)
	person, exists := e.persons[id]
	if !exists {
		log.Printf("Person with ID %s not found", id)
		return 0, ErrNotFound
	}

	height := 0
	parent := person.Head
	for parent != nil && *parent != "" {
		height++
		log.Printf("Current height: %d, current parent ID: %s", height, *parent)
		p, exists := e.persons[*parent]
		if !exists {
			log.Printf("Parent with ID %s not found", *parent)
			return 0, ErrNotFound
		}
		parent = p.Head
	}

	log.Printf("Final height for person ID %s: %d", id, height)
	return height, nil
}

func (e *BleveEngine) lessCommonAncestor(from, to string) (string, error) {
	log.Printf("Finding least common ancestor for IDs: from=%s, to=%s", from, to)

	hFrom, err := e.personHeight(from)
	if err != nil {
		log.Printf("Error calculating height for 'from' ID %s: %v", from, err)
		return "", err
	}

	hTo, err := e.personHeight(to)
	if err != nil {
		log.Printf("Error calculating height for 'to' ID %s: %v", to, err)
		return "", err
	}

	log.Printf("Initial heights - from: %d, to: %d", hFrom, hTo)

	for hFrom != hTo {
		if hFrom > hTo {
			log.Printf("Moving up from ID %s", from)
			from = *e.persons[from].Head
			hFrom--
		} else {
			log.Printf("Moving up to ID %s", to)
			to = *e.persons[to].Head
			hTo--
		}
	}

	for from != to {
		log.Printf("Moving up both IDs: from=%s, to=%s", from, to)
		from = *e.persons[from].Head
		to = *e.persons[to].Head
	}

	log.Printf("Least common ancestor found: %s", from)
	return from, nil
}

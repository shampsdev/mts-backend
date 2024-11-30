package adapter

import (
	"encoding/json"
	"fmt"

	"api.mts.shamps.dev/internal/domain"
)

type JsonAdapter struct{}

func (ja *JsonAdapter) Parse(data []byte) (*domain.Person, error) {
	var persons []Person
	if err := json.Unmarshal(data, &persons); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	if len(persons) > 0 {
		return convertToDomainPerson(&persons[0]), nil
	}

	return nil, fmt.Errorf("no persons found in data")
}

func convertToDomainPerson(p *Person) *domain.Person {
	return &domain.Person{
		ID:            p.ID,
		Surname:       p.Surname,
		Name:          p.Name,
		MiddleNameRus: p.MiddleNameRus,
		JobTitle:      p.JobTitle,
		Status:        p.Status,
		Contacts: domain.ContactInfo{
			Email: p.Contacts.Email,
			Phone: p.Contacts.Phone,
		},
		WorkingHour: p.WorkingHour,
		Workplace:   p.Workplace,
		Head:        p.Head,
		Department:  p.Department,
		Division:    p.Division,
		Team:        p.Team,
		About:       p.About,
	}
}

func (ja *JsonAdapter) GetAll(data []byte) ([]*domain.Person, error) {
	var persons []Person
	if err := json.Unmarshal(data, &persons); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	var result []*domain.Person
	for _, person := range persons {
		result = append(result, convertToDomainPerson(&person))
	}
	return result, nil
}

func NewJsonAdapter() *JsonAdapter {
	return &JsonAdapter{}
}

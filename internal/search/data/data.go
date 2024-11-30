package data

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"api.mts.shamps.dev/internal/domain"
)

func LoadPersons() []*domain.Person {
	file, err := os.Open("internal/search/data/data.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var persons []*domain.Person
	if err := json.Unmarshal(bytes, &persons); err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	log.Printf("Loaded %d persons", len(persons))

	return persons
}

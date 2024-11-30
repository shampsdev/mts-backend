package search

import (
	"io"
	"log"
	"os"

	"api.mts.shamps.dev/external/adapter"
	"api.mts.shamps.dev/internal/domain"
)

func loadData(a adapter.Adapter) []*domain.Person {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}
	persons, err := a.GetAll(bytes)
	if err != nil {
		log.Fatalf("failed to get all persons: %v", err)
	}

	return persons
}

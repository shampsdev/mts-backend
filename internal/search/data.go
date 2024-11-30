package search

import "api.mts.shamps.dev/internal/domain"

var data = []*domain.PersonNode{
	{
		ID:        "unique-id-1",
		Name:      "Mike de Geofroy",
		GroupID:   "group-1",
		GroupName: "Engineering",
		Status:    "Working",
		Image:     "https://t.me/i/userpic/320/mikedegeofroy.jpg",
		JobTitle:  "Team Lead",
		Children:  []string{"unique-id-2", "unique-id-3", "unique-id-4"},
		Parents:   []string{},
	},
	{
		ID:        "unique-id-2",
		Name:      "Mike de Geofroy",
		GroupID:   "group-1",
		GroupName: "Engineering",
		Status:    "Working",
		Image:     "https://t.me/i/userpic/320/mikedegeofroy.jpg",
		JobTitle:  "Frontend Developer",
		Children:  []string{},
		Parents:   []string{"unique-id-1"},
	},
	{
		ID:        "unique-id-3",
		Name:      "John Smith",
		GroupID:   "group-2",
		GroupName: "Engineering",
		Status:    "Working",
		Image:     "https://thispersondoesnotexist.com/",
		JobTitle:  "Backend Developer",
		Children:  []string{},
		Parents:   []string{"unique-id-1"},
	},
	{
		ID:        "unique-id-4",
		Name:      "Sarah Connor",
		GroupID:   "group-1",
		GroupName: "Quality Assurance",
		Status:    "Working",
		Image:     "https://thispersondoesnotexist.com/",
		JobTitle:  "QA / DevOps",
		Children:  []string{},
		Parents:   []string{"unique-id-1"},
	},
}

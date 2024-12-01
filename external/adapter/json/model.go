package adapter

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Person struct {
	ID            string      `json:"id"`
	Surname       string      `json:"surname"`
	Name          string      `json:"name"`
	MiddleName string         `json:"middle_name"`
	JobTitle      string      `json:"jobtitle"`
	Status        string      `json:"status"`
	Contacts      ContactInfo `json:"contacts"`
	WorkingHour   string      `json:"working_hour"`
	Workplace     string      `json:"workplace"`
	Head          *string     `json:"head"`
	Department    string      `json:"department"`
	Division      *string     `json:"division"`
	Team          *string     `json:"team"`
	Image         string      `json:"image"`
	About         string      `json:"about"`
}

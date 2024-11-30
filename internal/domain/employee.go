package domain

type PersonNode struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	GroupID   string   `json:"groupid"`
	GroupName string   `json:"groupname"`
	Status    string   `json:"status"`
	Image     string   `json:"image"`
	JobTitle  string   `json:"jobtitle"`
	Children  []string `json:"children"`
	Parents   []string `json:"parents"`
}

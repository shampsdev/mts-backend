package domain

type ContactInfo struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Person struct {
	ID            string      `json:"id"`
	Surname       string      `json:"surname"`
	Name          string      `json:"name"`
	MiddleNameRus string      `json:"middle_name_rus"`
	JobTitle      string      `json:"jobtitle"`
	Status        string      `json:"status"`
	Contacts      ContactInfo `json:"contacts"`
	WorkingHour   string      `json:"working_hour"`
	Workplace     string      `json:"workplace"`
	Head          *string     `json:"head"`
	Children      []string    `json:"children"`
	Department    string      `json:"department"`
	Division      *string     `json:"division"`
	Team          *string     `json:"team"`
	Image         string      `json:"image"`
	About         string      `json:"about"`
}

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

func PersonToNode(p *Person) *PersonNode {
	node := &PersonNode{
	  ID:        p.ID,
	  Name:      p.Surname + " " + p.Name + " " + p.MiddleNameRus,
	  GroupID:   p.Department,
	  GroupName: p.Department,
	  Status:    p.Status,
	  JobTitle:  p.JobTitle,
	  Children:  p.Children,
	  Parents:   []string{},
	  Image:     p.Image,
	}
  
	if node.Parents == nil || len(node.Parents) == 0 || (len(node.Parents) == 1 && node.Parents[0] == "") {
	  node.Parents = []string{}
	}
  
	if p.Head != nil && *p.Head != "" {
	  node.Parents = []string{*p.Head}
	}
  
	if node.Children == nil {
	  node.Children = []string{}
	}
  
	return node
  }
  
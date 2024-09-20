package domain

import "fmt"

// Group struct is a struct to represent a Gitlab group
type Group struct {
	// ID is the group identifier
	ID int
	// Name is the group name
	Name string
	// Description is the group description
	Description string
	// Path is the group path
	Path string
	// WebURL is the group web url
	WebURL string
}

// NewGroup creates a new Group instance
func NewGroup(id int, name, description, path, url string) Group {
	return Group{
		ID:          id,
		Name:        name,
		Description: description,
		Path:        path,
		WebURL:      url,
	}
}

// String returns a string representation of a Group
func (g Group) String() string {
	return fmt.Sprintf("%d:\t%s %s %s %s", g.ID, g.Name, g.Description, g.Path, g.WebURL)
}

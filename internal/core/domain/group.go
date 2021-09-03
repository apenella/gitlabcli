package domain

import "fmt"

type Group struct {
	Id          int
	Name        string
	Description string
	Path        string
	WebUrl      string
}

func NewGroup(id int, name, description, path, url string) Group {
	return Group{
		Id:          id,
		Name:        name,
		Description: description,
		Path:        path,
		WebUrl:      url,
	}
}

func (g Group) String() string {
	return fmt.Sprintf("%d:\t%s %s %s %s", g.Id, g.Name, g.Description, g.Path, g.WebUrl)
}

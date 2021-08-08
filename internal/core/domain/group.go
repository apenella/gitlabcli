package domain

import "fmt"

type Group struct {
	Id   int
	Name string
}

func NewGroup(id int, name string) Group {
	return Group{
		Id:   id,
		Name: name,
	}
}

func (g Group) String() string {
	return fmt.Sprintf("%d: %s", g.Id, g.Name)
}

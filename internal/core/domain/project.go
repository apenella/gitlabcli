package domain

import "fmt"

// Project struct is a struct to represent a Gitlab project
type Project struct {
	// ID is the project identifier
	ID int
	// Name is the project name
	Name string
	// Description is the project description
	Description string
	// Path is the project path
	Path string
	// Sshurl is the project ssh url
	Sshurl string
	// Httpurl is the project http url
	Httpurl string
	// DefaultBranch is the project default branch
	DefaultBranch string
}

// NewProject creates a new Project instance
func NewProject(id int, name, description, branch, path, sshurl, httpurl string) Project {
	return Project{
		ID:            id,
		Name:          name,
		Description:   description,
		Path:          path,
		Sshurl:        sshurl,
		Httpurl:       httpurl,
		DefaultBranch: branch,
	}
}

// String returns a string representation of a Project
func (p Project) String() string {
	return fmt.Sprintf("%d: %s %s [%s]", p.ID, p.Name, p.Path, p.Sshurl)
}

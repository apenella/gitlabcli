package domain

import "fmt"

type Project struct {
	Id            int
	Name          string
	Path          string
	Sshurl        string
	Httpurl       string
	DefaultBranch string
}

func NewProject(id int, name, branch, path, sshurl, httpurl string) Project {
	return Project{
		Id:            id,
		Name:          name,
		Path:          path,
		Sshurl:        sshurl,
		Httpurl:       httpurl,
		DefaultBranch: branch,
	}
}

func (p Project) String() string {
	return fmt.Sprintf("%d: %s %s [%s]", p.Id, p.Name, p.Path, p.Sshurl)
}

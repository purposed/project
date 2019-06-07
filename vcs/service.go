package vcs

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

// ProjectService regroups project collection functions.
type ProjectService struct {
	RootPath string
}

// NewService returns a new service.
func NewService(path string) *ProjectService {
	return &ProjectService{path}
}

// Clone clones a repository.
func (s *ProjectService) Clone(repoInfoStr string) error {
	repo, err := ParseRepository(repoInfoStr)
	if err != nil {
		return err
	}

	repo.SetRoot(s.RootPath)

	if err := repo.InitDirectories(s.RootPath); err != nil {
		return err
	}

	if err := repo.Clone(); err != nil {
		return err
	}

	return nil
}

// Create creates a new repository.
func (s *ProjectService) Create(name, owner, provider string) error {
	repo := &Repository{Name: name, Owner: owner, Provider: provider}
	repo.SetRoot(s.RootPath)

	pPath := repo.ProjectPath(s.RootPath)

	if _, err := os.Stat(pPath); os.IsNotExist(err) {
		if err := os.MkdirAll(pPath, os.ModePerm); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("directory [%s] already exists", pPath)
	}

	if err := os.Chdir(pPath); err != nil {
		return err
	}

	// Git init
	cmd := exec.Command("git", "init")
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}

func (s *ProjectService) listForProvider(providerName, owner string) ([]*Repository, error) {
	var out []*Repository

	providerPath := path.Join(s.RootPath, SourceDirPrefix, providerName)

	files, err := ioutil.ReadDir(providerPath)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if !f.IsDir() || strings.HasPrefix(f.Name(), ".") || f.Name() != owner {
			continue
		}

		projects, err := ioutil.ReadDir(path.Join(providerPath, f.Name()))
		if err != nil {
			return nil, err
		}

		for _, p := range projects {
			if !p.IsDir() || strings.HasPrefix(f.Name(), ".") {
				continue
			}

			out = append(out, newRepository(p.Name(), owner, providerName, s.RootPath))
		}
	}

	return out, nil
}

// List lists all projects for a given owner.
func (s *ProjectService) List(owner string) ([]*Repository, error) {
	srcDir := path.Join(s.RootPath, SourceDirPrefix)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		return nil, err
	}

	var out []*Repository

	for _, f := range files {
		if !f.IsDir() || strings.HasPrefix(f.Name(), ".") {
			continue
		}

		projects, err := s.listForProvider(f.Name(), owner)
		if err != nil {
			return nil, err
		}

		out = append(out, projects...)
	}

	return out, nil

}

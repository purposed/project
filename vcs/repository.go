package vcs

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	providerSSHTemplate = "git@%s:%s/%s.git"
	SourceDirPrefix     = "src"
)

// Repository represents a remote project repository.
type Repository struct {
	Name     string
	Owner    string
	Provider string
	URL      string
}

func newRepository(name, owner, provider string) *Repository {
	return &Repository{name, owner, provider, fmt.Sprintf(providerSSHTemplate, provider, owner, name)}
}

// Pretty returns a prettified version of the repository.
func (r *Repository) Pretty() string {
	return fmt.Sprintf("%s/%s/%s", r.Provider, r.Owner, r.Name)
}

// ParseRepository generates a repository object from an URL string.
func ParseRepository(repoInfo string) (Repository, error) {
	splitted := strings.Split(repoInfo, "/")
	if len(splitted) != 3 {
		return Repository{}, fmt.Errorf("could not parse repo [%s]", repoInfo)
	}

	url := fmt.Sprintf(providerSSHTemplate, splitted[0], splitted[1], splitted[2])

	return Repository{Name: splitted[2], Owner: splitted[1], Provider: splitted[0], URL: url}, nil
}

// ProjectDir returns the folder containing projects for this provider/owner combination.
func (r *Repository) ProjectDir(rootPath string) string {
	return path.Join(rootPath, SourceDirPrefix, r.Provider, r.Owner)
}

// ProjectPath returns the project path.
func (r *Repository) ProjectPath(rootPath string) string {
	return path.Join(r.ProjectDir(rootPath), r.Name)
}

// InitDirectories creates all the directories that have to exist in order to clone the repo.
func (r *Repository) InitDirectories(rootPath string) error {
	projectPath := r.ProjectDir(rootPath)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return os.MkdirAll(projectPath, os.ModePerm)
	}
	return nil
}

// Clone clones the repository in the specified dir.
func (r *Repository) Clone(rootPath string) error {
	// TODO: Support HTTPS authentication
	cmd := exec.Command("git", "clone", r.URL, r.ProjectPath(rootPath))
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}

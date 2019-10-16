package vcs

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	providerSSHTemplate = "git@%s:%s/%s.git"

	sourceDirPrefix = "src"
)

// Repository represents a remote project repository.
type Repository struct {
	Name     string
	Owner    string
	Provider string
	URL      string

	rootPath string
}

func newRepository(name, owner, provider, rootPath string) *Repository {
	return &Repository{
		Name:     name,
		Owner:    owner,
		Provider: provider,
		URL:      fmt.Sprintf(providerSSHTemplate, provider, owner, name),
		rootPath: rootPath,
	}
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

// ProjectsListDir returns the folder containing projects for this provider/owner combination.
func (r *Repository) ProjectsListDir(rootPath string) string {
	return path.Join(rootPath, sourceDirPrefix, r.Provider, r.Owner)
}

// ProjectPath returns the project path.
func (r *Repository) ProjectPath(rootPath string) string {
	return path.Join(r.ProjectsListDir(rootPath), r.Name)
}

// InitDirectories creates all the directories that have to exist in order to clone the repo.
func (r *Repository) InitDirectories(rootPath string) error {
	projectPath := r.ProjectsListDir(rootPath)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return os.MkdirAll(projectPath, os.ModePerm)
	}
	return nil
}

// SetRoot sets the root path of the repository.
func (r *Repository) SetRoot(rootPath string) {
	r.rootPath = rootPath
}

// Fetch runs "git fetch" in the repo's root.
func (r *Repository) Fetch() error {
	curPath, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := os.Chdir(r.ProjectPath(r.rootPath)); err != nil {
		return err
	}
	defer os.Chdir(curPath)

	cmd := exec.Command("git", "fetch")
	_, err = cmd.CombinedOutput()
	if code := cmd.ProcessState.ExitCode(); code == 128 {
		return errors.New("no git repository or remote")
	}
	return err
}

// Clone clones the repository in the specified dir.
func (r *Repository) Clone() error {
	// TODO: Support HTTPS authentication
	cmd := exec.Command("git", "clone", r.URL, r.ProjectPath(r.rootPath))
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
	}
	return err
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

const (
	bitbucketHTTPS = "https://([a-zA-Z]+)@(?P<provider>bitbucket.org)/[a-zA-Z]+/([a-zA-Z]+)\\.git"
	bitbucketSSH   = "git@(?P<provider>bitbucket.org):(?P<username>[a-zA-Z]+)/(?P<project>[a-zA-Z]+)\\.git"

	githubHTTPS = "https://(?P<provider>github.com)/(?P<username>[a-zA-Z]+)/(?P<project>[a-zA-Z]+)\\.git"
	githubSSH   = "git@(?P<provider>github.com):(?P<username>[a-zA-Z]+)/(?P<project>[a-zA-Z]+)\\.git"

	gitlabHTTPS = "https://(?P<provider>gitlab.com)/(?P<username>[a-zA-Z]+)/(&P<project>[a-zA-Z]+)\\.git"
	gitlabSSH   = "git@(?P<provider>gitlab.com):(?P<username>[a-zA-Z]+)/(?P<project>[a-zA-Z]+).git"
)

// Repository represents a remote project repository.
type Repository struct {
	Name     string
	Owner    string
	Provider string
	URL      string
}

// ParseRepository generates a repository object from an URL string.
func ParseRepository(url string) (Repository, error) {

	repo := Repository{URL: url}
	for _, pattern := range []string{bitbucketHTTPS, bitbucketSSH, gitlabHTTPS, gitlabSSH, githubHTTPS, githubSSH} {
		exp := regexp.MustCompile(pattern)
		if matches := exp.FindStringSubmatch(url); matches != nil {
			for i, name := range exp.SubexpNames() {
				switch name {
				case "provider":
					repo.Provider = matches[i]
				case "username":
					repo.Owner = strings.ToLower(matches[i])
				case "project":
					repo.Name = matches[i]
				default:
					continue
				}
			}
		} else {
			continue
		}
	}

	return repo, nil
}

func (r *Repository) projectDir(rootPath string) string {
	return path.Join(rootPath, "src", r.Provider, r.Owner)
}

func (r *Repository) projectPath(rootPath string) string {
	return path.Join(r.projectDir(rootPath), r.Name)
}

// InitDirectories creates all the directories that have to exist in order to clone the repo.
func (r *Repository) InitDirectories(rootPath string) error {
	projectPath := r.projectDir(rootPath)
	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		return os.MkdirAll(projectPath, os.ModePerm)
	}
	return nil
}

// Clone clones the repository in the specified dir.
func (r *Repository) Clone(rootPath string) error {
	// TODO: Support HTTPS authentication
	cmd := exec.Command("git", "clone", r.URL, r.projectPath(rootPath))
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	return err
}

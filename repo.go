package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Repo struct {
	Description string
	Dir         string
	Name        string
	Username    string
}

func (r Repo) Path() string {
	return path.Join(r.Dir, r.Name)
}

func (r Repo) Remote() string {
	return fmt.Sprintf("git@github.com:%s/%s", r.Username, r.Name)
}

func (r Repo) Create(c Config) (err error) {
	if err = r.createGithubRepo(c.Github.Token); err != nil {
		return err
	}

	command := exec.Command("git", "clone", r.Remote(), r.Path())

	return command.Run()
}

func (r Repo) Finish(d Docker, c CircleCI, l Licence) (err error) {
	var f *os.File

	p := path.Join(r.Path(), "README.md")

	if f, err = os.Create(p); err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	if _, err = w.WriteString(r.Readme(d.HubURL(), c.CircleURL(), l.Type, l.Text)); err != nil {
		return
	}
	w.Flush()

	addCommand := exec.Command("git", "add", ".")
	addCommand.Dir = r.Path()
	if err = addCommand.Run(); err != nil {
		return
	}

	commitCommand := exec.Command("git", "commit", "-m", "Initialise Repo")
	commitCommand.Dir = r.Path()
	if err = commitCommand.Run(); err != nil {
		return
	}

	pushCommand := exec.Command("git", "push")
	pushCommand.Dir = r.Path()
	err = pushCommand.Run()

	return
}

func (r Repo) createGithubRepo(token string) (err error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	repo := &github.Repository{
		Name:        github.String(r.Name),
		Private:     github.Bool(false),
		Description: github.String(r.Description),
	}

	_, _, err = client.Repositories.Create("", repo)

	return
}

func (r Repo) Readme(hubUrl, circleUrl, licenceType, licenceText string) string {
	return fmt.Sprintf(`%s
==

| who       | what |
|-----------|------|
| dockerhub | %s   |
| circleci  | %s   |
| licence   | %s   |


Licence
--

%s
`, r.Name, hubUrl, circleUrl, licenceType, licenceText)
}

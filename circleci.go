package main

import (
    "crypto/tls"
    "fmt"
    "net/http"
)

type CircleCI struct {
    Username string
    Repo string
    Token string
}

func NewCircleCI(config Config, repo string) (c CircleCI, err error) {
    c.Username = config.Github.Username
    c.Repo = repo
    c.Token = config.CircleCI.Token

    return
}

func (c CircleCI) Follow() (err error) {
    tr := &http.Transport{
        TLSClientConfig:    &tls.Config{},
        DisableCompression: true,
    }

    client := &http.Client{Transport: tr}

    _, err = client.Post(c.followURL(), "", nil)

    return
}

func (c CircleCI) followURL() string {
    return fmt.Sprintf("https://circleci.com/api/v1.1/project/github/%s/%s?circle-token=%s", c.Username, c.Repo, c.Token)
}

func (c CircleCI) CircleURL() string {
    return fmt.Sprintf("https://circleci.com/gh/%s/%s", c.Username, c.Repo)
}

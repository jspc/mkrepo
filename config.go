package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    // Basic Options
    Name string
    EmailAddress string

    // Git repos
    Git struct {
        Licence string
    }

    // Auth and the like
    Github struct {
        Token string
        Username string
    }

    // Docker
    Docker struct {
        DefaultDockerImage string
        Username string
    }

    // Use circleci for tests
    CircleCI struct {
        Token string
    }
}

func LoadConfig(path string) (c Config, err error) {
    var f *os.File

    if f, err = os.Open(path); err != nil {
        return err
    }
    defer f.Close()

    decoder := json.NewDecoder(f)
    err = decoder.Decode(&c)

    return
}

func (c Config) Maintainer() string {
    return fmt.Sprintf("%s <%s>", c.Name, c.EmailAddress)
}

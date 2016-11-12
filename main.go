package main

import (
    "flag"
    "log"
    "os"
    "os/user"
    "path"
)

var noDocker *bool

var circleCI CircleCI
var config Config
var docker Docker
var licence Licence
var repo Repo

func init() {
    noDocker = flag.Bool("nd", false, "Don't Configure Docker")
}

func main() {
    var u *user.User
    var err error

    flag.Parse()

    repo.Name = flag.Arg(0)
    repo.Description = flag.Arg(1)

    if u, err = user.Current(); err != nil {
        log.Fatal(err)
    }

    if err = config.Load( path.Join(u.HomeDir, ".config", "mkrepo.json") ); err != nil {
        log.Fatal(err)
    }

    repo.Username = config.Github.Username

    if repo.Dir, err = os.Getwd(); err != nil {
        log.Fatalf("Could not get current directory: %s", err)
    }

    log.Printf("Creating %s", repo.Path())

    if err = repo.Create(config); err != nil {
        log.Fatalf("Error creating repo: %s", err)
    }

    if licence, err = NewLicence(config, repo.Path()); err != nil {
        log.Fatalf("Error configuring licence: %s", err)
    }

    if err = licence.Write(); err != nil {
        log.Fatalf("Error writing licence file: %s", err)
    }

    if !*noDocker {
        if docker, err = NewDocker(config, repo.Name, repo.Path()); err != nil {
            log.Fatalf("Error configuring Docker: %s", err)
        }

        if err = docker.Write(); err != nil {
            log.Fatalf("Error writing Dockerfile: %s", err)
        }
    }

    if circleCI, err = NewCircleCI(config, repo.Name); err != nil {
        log.Fatalf("Error connecting to CircleCI: %s", err)
    }

    repo.Finish(docker, circleCI, licence)


    if err = circleCI.Follow(); err != nil {
        log.Fatalf("Error following project on CircleCI: %s", err)
    }

}

package main

import (
    "flag"
    "log"
    "os"
    "os/user"
    "path"
)

var (
    noDocker *bool
    configPath *string

    circleCI CircleCI
    config Config
    docker Docker
    licence Licence
    repo Repo
)

func init() {
    var u *user.User
    var err error
    if u, err = user.Current(); err != nil {
        log.Fatal(err)
    }

    noDocker = flag.Bool("nd", false, "Don't Configure Docker")
    configPath= flag.String("config", path.Join(u.HomeDir, ".config", "mkrepo.json"), "Path to mkrepo config")
}

func main() {
    var err error

    flag.Parse()

    repo.Name = flag.Arg(0)
    if repo.Name == "" {
        log.Fatalf("Missing project name")
    }

    repo.Description = flag.Arg(1)


    if config,err = LoadConfig( *configPath ); err != nil {
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

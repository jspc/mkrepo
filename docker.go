package main

import (
    "bufio"
    "fmt"
    "os"
    "path"
)

type Docker struct {
    BaseImage string
    Maintainer string
    ContainerName string
    Dest string
}

func NewDocker(c Config, repo, dest string) (d Docker, err error) {
    d.BaseImage = c.Docker.DefaultDockerImage
    d.Maintainer = c.Maintainer()
    d.ContainerName = fmt.Sprintf("%s/%s", c.Docker.Username, repo)
    d.Dest = dest

    return
}

func (d Docker) Write() (err error) {
    var f *os.File

    p := path.Join(d.Dest, "Dockerfile")

    if f, err = os.Create(p); err != nil {
        return
    }
    defer f.Close()

    w := bufio.NewWriter(f)
    if _, err = w.WriteString( d.Dockerfile() ); err != nil {
        return
    }
    w.Flush()

    return

}

func (d Docker) HubURL() string {
    return fmt.Sprintf("https://hub.docker.com/r/%s/", d.ContainerName)
}

func (d Docker) Dockerfile() string {
    return fmt.Sprintf("FROM %s\nMAINTAINER %s\n\n", d.BaseImage, d.Maintainer)
}

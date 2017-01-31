package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var expectedDockerfile = `FROM some-base
MAINTAINER some-maintainer

`

func TestNewDocker(t *testing.T) {
	p := "./.test/mkrepo.json"
	c, err := LoadConfig(p)

	if err != nil {
		t.Errorf("LoadConfig() error: %q", err)
	}

	repo := "hello-world"
	dest := ".test/repo"
	d, err := NewDocker(c, repo, dest)
	if err != nil {
		t.Errorf("NewDocker() error: %q", err)
	}

	t.Run("Test BaseImage value", func(t *testing.T) {
		if d.BaseImage != "alpine" {
			t.Errorf("d.BaseImage error: received %q, expected 'alpine'", d.BaseImage)
		}
	})

	t.Run("Test Maintainer value", func(t *testing.T) {
		if d.Maintainer != "gopher <gopher@example.com>" {
			t.Errorf("d.Maintainer error: received %q, expected 'gopher <gopher@example.com>'", d.Maintainer)
		}
	})

	t.Run("Test Dest value", func(t *testing.T) {
		if d.Dest != dest {
			t.Errorf("d.Dest error: received %q, expected %q", d.Dest, dest)
		}
	})

	t.Run("Test ContainerName value", func(t *testing.T) {
		if d.ContainerName != fmt.Sprintf("admin/%s", repo) {
			t.Errorf("d.ContainerName error: received %q, expected %q", d.ContainerName, fmt.Sprintf("admin/%s", repo))
		}
	})
}

func TestHubURL(t *testing.T) {
	d := Docker{"some-base", "some-maintainer", "container-name", "dir"}

	if d.HubURL() != "https://hub.docker.com/r/container-name/" {
		t.Errorf("d.HubURL() error: received %q, expected 'https://hub.docker.com/r/container-name/'", d.HubURL())
	}
}

func TestDockerFile(t *testing.T) {
	d := Docker{"some-base", "some-maintainer", "container-name", "dir"}

	if d.Dockerfile() != expectedDockerfile {
		t.Errorf("d.Dockerfile() error: received %q, expected %q", d.Dockerfile(), expectedDockerfile)
	}
}

func TestWrite(t *testing.T) {
	d := Docker{"some-base", "some-maintainer", "container-name", ".test/repo"}

	err := d.Write()
	if err != nil {
		t.Errorf("d.Write() error: %q", err)
	}

	writtenDockerfile, _ := ioutil.ReadFile(".test/repo/Dockerfile")
	if string(writtenDockerfile) != expectedDockerfile {
		t.Errorf("d.Write() error: wrote %q, expected %q", string(writtenDockerfile), expectedDockerfile)
	}
}

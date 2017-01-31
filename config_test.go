package main

import (
    "testing"
)

func TestLoadConfig(t *testing.T) {
    p := "./.test/mkrepo.json"
    c,err := LoadConfig(p)

    if err != nil {
        t.Errorf("LoadConfig() error: %q", err)
    }

    t.Run("Test Name value", func(t *testing.T){
        if c.Name != "gopher" {
            t.Errorf("config.Name error: received %q, expected 'gopher'", c.Name)
        }
    })

    t.Run("Test EmailAddress value", func(t *testing.T){
        if c.EmailAddress != "gopher@example.com" {
            t.Errorf("config.EmailAddress error: received %q, expected 'gopher@example.com'", c.EmailAddress)
        }
    })

    t.Run("Test Git.Licence value", func(t *testing.T){
        if c.Git.Licence != "MIT" {
            t.Errorf("config.Git.Licence error: received %q, expected 'MIT'", c.Git.Licence)
        }
    })

    t.Run("Test Github.Username value", func(t *testing.T){
        if c.Github.Username != "admin" {
            t.Errorf("config.Github.Username error: received %q, expected 'admin'", c.Github.Token)
        }
    })

    t.Run("Test Github.Token value", func(t *testing.T){
        if c.Github.Token != "abcdefgh" {
            t.Errorf("config.Github.Token error: received %q, expected 'abcdefgh'", c.Github.Username)
        }
    })

    t.Run("Test Docker.DefaultDockerImage value", func(t *testing.T){
        if c.Docker.DefaultDockerImage != "alpine" {
            t.Errorf("config.Docker.DefaultDockerImage error: received %q, expected 'alpine'", c.Docker.DefaultDockerImage)
        }
    })

    t.Run("Test Docker.Username value", func(t *testing.T){
        if c.Docker.Username != "admin" {
            t.Errorf("config.Docker.Username error: received %q, expected 'admin'", c.Docker.Username)
        }
    })

    t.Run("Test CircleCI.Token value", func(t *testing.T){
        if c.CircleCI.Token != "123456" {
            t.Errorf("config.CircleCI.Token error: received %q, expected '123456'", c.CircleCI.Token)
        }
    })
}

func TestMaintainer(t *testing.T) {
    p := "./.test/mkrepo.json"
    c,err := LoadConfig(p)

    if err != nil {
        t.Errorf("LoadConfig() error: %q", err)
    }

    if c.Maintainer() != "gopher <gopher@example.com>" {
        t.Errorf("c.Maintainer() error: received %q, expected 'gopher <gopher@example.com>'")
    }
}

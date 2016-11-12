package main

import "testing"

func TestConfig_Load(t *testing.T) {
    type fields struct {
        Name         string
        EmailAddress string
        Git          struct{ Licence string }
        Github       struct {
            Token    string
            Username string
        }
        Docker struct {
            DefaultDockerImage string
            Username           string
        }
        CircleCI struct{ Token string }
    }
    type args struct {
        path string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := &Config{
                Name:         tt.fields.Name,
                EmailAddress: tt.fields.EmailAddress,
                Git:          tt.fields.Git,
                Github:       tt.fields.Github,
                Docker:       tt.fields.Docker,
                CircleCI:     tt.fields.CircleCI,
            }
            if err := c.Load(tt.args.path); (err != nil) != tt.wantErr {
                t.Errorf("Config.Load() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestConfig_Maintainer(t *testing.T) {
    type fields struct {
        Name         string
        EmailAddress string
        Git          struct{ Licence string }
        Github       struct {
            Token    string
            Username string
        }
        Docker struct {
            DefaultDockerImage string
            Username           string
        }
        CircleCI struct{ Token string }
    }
    tests := []struct {
        name   string
        fields fields
        want   string
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := &Config{
                Name:         tt.fields.Name,
                EmailAddress: tt.fields.EmailAddress,
                Git:          tt.fields.Git,
                Github:       tt.fields.Github,
                Docker:       tt.fields.Docker,
                CircleCI:     tt.fields.CircleCI,
            }
            if got := c.Maintainer(); got != tt.want {
                t.Errorf("Config.Maintainer() = %v, want %v", got, tt.want)
            }
        })
    }
}

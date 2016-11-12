package main

import (
    "reflect"
    "testing"
)

func TestNewDocker(t *testing.T) {
    type args struct {
        c    Config
        repo string
        dest string
    }
    tests := []struct {
        name    string
        args    args
        wantD   Docker
        wantErr bool
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotD, err := NewDocker(tt.args.c, tt.args.repo, tt.args.dest)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewDocker() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotD, tt.wantD) {
                t.Errorf("NewDocker() = %v, want %v", gotD, tt.wantD)
            }
        })
    }
}

func TestDocker_Write(t *testing.T) {
    type fields struct {
        BaseImage     string
        Maintainer    string
        ContainerName string
        Dest          string
    }
    tests := []struct {
        name    string
        fields  fields
        wantErr bool
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            d := Docker{
                BaseImage:     tt.fields.BaseImage,
                Maintainer:    tt.fields.Maintainer,
                ContainerName: tt.fields.ContainerName,
                Dest:          tt.fields.Dest,
            }
            if err := d.Write(); (err != nil) != tt.wantErr {
                t.Errorf("Docker.Write() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestDocker_HubURL(t *testing.T) {
    type fields struct {
        BaseImage     string
        Maintainer    string
        ContainerName string
        Dest          string
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
            d := Docker{
                BaseImage:     tt.fields.BaseImage,
                Maintainer:    tt.fields.Maintainer,
                ContainerName: tt.fields.ContainerName,
                Dest:          tt.fields.Dest,
            }
            if got := d.HubURL(); got != tt.want {
                t.Errorf("Docker.HubURL() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestDocker_Dockerfile(t *testing.T) {
    type fields struct {
        BaseImage     string
        Maintainer    string
        ContainerName string
        Dest          string
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
            d := Docker{
                BaseImage:     tt.fields.BaseImage,
                Maintainer:    tt.fields.Maintainer,
                ContainerName: tt.fields.ContainerName,
                Dest:          tt.fields.Dest,
            }
            if got := d.Dockerfile(); got != tt.want {
                t.Errorf("Docker.Dockerfile() = %v, want %v", got, tt.want)
            }
        })
    }
}

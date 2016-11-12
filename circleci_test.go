package main

import (
    "reflect"
    "testing"
)

func TestNewCircleCI(t *testing.T) {
    type args struct {
        config Config
        repo   string
    }
    tests := []struct {
        name    string
        args    args
        wantC   CircleCI
        wantErr bool
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotC, err := NewCircleCI(tt.args.config, tt.args.repo)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewCircleCI() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotC, tt.wantC) {
                t.Errorf("NewCircleCI() = %v, want %v", gotC, tt.wantC)
            }
        })
    }
}

func TestCircleCI_Follow(t *testing.T) {
    type fields struct {
        Username string
        Repo     string
        Token    string
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
            c := CircleCI{
                Username: tt.fields.Username,
                Repo:     tt.fields.Repo,
                Token:    tt.fields.Token,
            }
            if err := c.Follow(); (err != nil) != tt.wantErr {
                t.Errorf("CircleCI.Follow() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestCircleCI_followURL(t *testing.T) {
    type fields struct {
        Username string
        Repo     string
        Token    string
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
            c := CircleCI{
                Username: tt.fields.Username,
                Repo:     tt.fields.Repo,
                Token:    tt.fields.Token,
            }
            if got := c.followURL(); got != tt.want {
                t.Errorf("CircleCI.followURL() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestCircleCI_CircleURL(t *testing.T) {
    type fields struct {
        Username string
        Repo     string
        Token    string
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
            c := CircleCI{
                Username: tt.fields.Username,
                Repo:     tt.fields.Repo,
                Token:    tt.fields.Token,
            }
            if got := c.CircleURL(); got != tt.want {
                t.Errorf("CircleCI.CircleURL() = %v, want %v", got, tt.want)
            }
        })
    }
}

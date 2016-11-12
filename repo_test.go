package main

import "testing"

func TestRepo_Path(t *testing.T) {
    type fields struct {
        Description string
        Dir         string
        Name        string
        Username    string
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
            r := Repo{
                Description: tt.fields.Description,
                Dir:         tt.fields.Dir,
                Name:        tt.fields.Name,
                Username:    tt.fields.Username,
            }
            if got := r.Path(); got != tt.want {
                t.Errorf("Repo.Path() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRepo_Remote(t *testing.T) {
    type fields struct {
        Description string
        Dir         string
        Name        string
        Username    string
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
            r := Repo{
                Description: tt.fields.Description,
                Dir:         tt.fields.Dir,
                Name:        tt.fields.Name,
                Username:    tt.fields.Username,
            }
            if got := r.Remote(); got != tt.want {
                t.Errorf("Repo.Remote() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestRepo_Create(t *testing.T) {
    type fields struct {
        Description string
        Dir         string
        Name        string
        Username    string
    }
    type args struct {
        c Config
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
            r := Repo{
                Description: tt.fields.Description,
                Dir:         tt.fields.Dir,
                Name:        tt.fields.Name,
                Username:    tt.fields.Username,
            }
            if err := r.Create(tt.args.c); (err != nil) != tt.wantErr {
                t.Errorf("Repo.Create() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRepo_Finish(t *testing.T) {
    type fields struct {
        Description string
        Dir         string
        Name        string
        Username    string
    }
    type args struct {
        d Docker
        c CircleCI
        l Licence
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
            r := Repo{
                Description: tt.fields.Description,
                Dir:         tt.fields.Dir,
                Name:        tt.fields.Name,
                Username:    tt.fields.Username,
            }
            if err := r.Finish(tt.args.d, tt.args.c, tt.args.l); (err != nil) != tt.wantErr {
                t.Errorf("Repo.Finish() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRepo_createGithubRepo(t *testing.T) {
    type fields struct {
        Description string
        Dir         string
        Name        string
        Username    string
    }
    type args struct {
        token string
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
            r := Repo{
                Description: tt.fields.Description,
                Dir:         tt.fields.Dir,
                Name:        tt.fields.Name,
                Username:    tt.fields.Username,
            }
            if err := r.createGithubRepo(tt.args.token); (err != nil) != tt.wantErr {
                t.Errorf("Repo.createGithubRepo() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRepo_Readme(t *testing.T) {
    type fields struct {
        Description string
        Dir         string
        Name        string
        Username    string
    }
    type args struct {
        hubUrl      string
        circleUrl   string
        licenceType string
        licenceText string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
        want   string
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            r := Repo{
                Description: tt.fields.Description,
                Dir:         tt.fields.Dir,
                Name:        tt.fields.Name,
                Username:    tt.fields.Username,
            }
            if got := r.Readme(tt.args.hubUrl, tt.args.circleUrl, tt.args.licenceType, tt.args.licenceText); got != tt.want {
                t.Errorf("Repo.Readme() = %v, want %v", got, tt.want)
            }
        })
    }
}

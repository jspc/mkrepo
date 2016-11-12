package main

import (
    "reflect"
    "testing"
)

func TestNewLicence(t *testing.T) {
    type args struct {
        c    Config
        dest string
    }
    tests := []struct {
        name    string
        args    args
        wantL   Licence
        wantErr bool
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotL, err := NewLicence(tt.args.c, tt.args.dest)
            if (err != nil) != tt.wantErr {
                t.Errorf("NewLicence() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(gotL, tt.wantL) {
                t.Errorf("NewLicence() = %v, want %v", gotL, tt.wantL)
            }
        })
    }
}

func TestLicence_Write(t *testing.T) {
    type fields struct {
        Type    string
        RawText string
        Args    LicenceArgs
        Text    string
        Dest    string
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
            l := &Licence{
                Type:    tt.fields.Type,
                RawText: tt.fields.RawText,
                Args:    tt.fields.Args,
                Text:    tt.fields.Text,
                Dest:    tt.fields.Dest,
            }
            if err := l.Write(); (err != nil) != tt.wantErr {
                t.Errorf("Licence.Write() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestLicence_addDate(t *testing.T) {
    type fields struct {
        Type    string
        RawText string
        Args    LicenceArgs
        Text    string
        Dest    string
    }
    tests := []struct {
        name   string
        fields fields
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l := &Licence{
                Type:    tt.fields.Type,
                RawText: tt.fields.RawText,
                Args:    tt.fields.Args,
                Text:    tt.fields.Text,
                Dest:    tt.fields.Dest,
            }
            l.addDate()
        })
    }
}

func TestLicence_addName(t *testing.T) {
    type fields struct {
        Type    string
        RawText string
        Args    LicenceArgs
        Text    string
        Dest    string
    }
    tests := []struct {
        name   string
        fields fields
    }{
    // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            l := &Licence{
                Type:    tt.fields.Type,
                RawText: tt.fields.RawText,
                Args:    tt.fields.Args,
                Text:    tt.fields.Text,
                Dest:    tt.fields.Dest,
            }
            l.addName()
        })
    }
}

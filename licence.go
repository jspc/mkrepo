package main

import (
	"bufio"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/nishanths/license/pkg/license"
)

type Licence struct {
	Type    string
	RawText string
	Args    LicenceArgs
	Text    string
	Dest    string
}

type LicenceArgs struct {
	Year     int
	Fullname string
}

func NewLicence(c Config, dest string) (l Licence, err error) {
	l.Type = c.Git.Licence
	l.Dest = dest

	client := license.NewClient()

	client.Config.Token = c.Github.Token
	client.Config.Username = c.Github.Username

	var lic license.License
	if lic, err = client.Info(l.Type); err != nil {
		return
	}

	year := time.Now().Year()
	l.Args = LicenceArgs{year, c.Name}
	l.RawText = lic.Body
	l.Text = lic.Body

	l.addDate()
	l.addName()

	return
}

func (l *Licence) Write() (err error) {
	var f *os.File

	p := path.Join(l.Dest, "LICENCE.md")

	if f, err = os.Create(p); err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	if _, err = w.WriteString(l.Text); err != nil {
		return
	}
	w.Flush()

	return
}

func (l *Licence) addDate() {
	l.Text = strings.Replace(l.Text, "[year]", strconv.Itoa(l.Args.Year), -1)
}

func (l *Licence) addName() {
	l.Text = strings.Replace(l.Text, "[fullname]", l.Args.Fullname, -1)
}

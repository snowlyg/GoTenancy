// +build ignore

package main

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Depado/articles/code/qor/admin/bindatafs"
)

func main() {
	var err error
	var vp string
	var ok bool
	afs := bindatafs.AssetFS

	logrus.SetLevel(logrus.DebugLevel)
	if vp, ok = detectQORdir(); !ok {
		logrus.Fatal("Could not detect a QOR Admin directory with assets. Aborting.")
	}
	logrus.WithField("path", vp).Debug("Highest candidate found")

	if err = afs.NameSpace("admin").RegisterPath(vp); err != nil {
		logrus.WithError(err).Fatal("Couldn't register path")
	}
	if err = afs.NameSpace("login").RegisterPath("admin/templates/"); err != nil {
		logrus.WithError(err).Fatal("Couldn't register templates with login directory")
	}
	if err := afs.Compile(); err != nil {
		logrus.WithError(err).Fatal("Couldn't compile templates")
	}
}

func detectViewsDir(path string) (string, bool) {
	var foundp string
	var found bool

	pkgorg := "github.com/qor"
	pkgname := "admin"
	ppath := filepath.Join(path, pkgorg)
	if _, err := os.Stat(ppath); err == nil {
		filepath.Walk(ppath, func(p string, f os.FileInfo, err error) error { // nolint: errcheck, gosec, unparam
			if found {
				return nil
			}
			if strings.HasPrefix(filepath.Base(p), pkgname) {
				vp := filepath.Join(p, "views")
				if _, err := os.Stat(vp); err == nil {
					logrus.WithField("path", vp).Debug("Found QOR Views Directory")
					foundp = vp
					found = true
				}
			}
			return nil
		})
	}
	return foundp, found
}

func detectQORdir() (string, bool) {
	var err error
	var found, d string
	var ok bool

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	if d, err = os.Getwd(); err != nil {
		logrus.WithError(err).Fatal("Couldn't get current working directory")
	}

	var candidates = []string{
		filepath.Join(gopath, "/pkg/mod/"),
		filepath.Join(gopath, "/src/"),
		filepath.Join(d, "/vendor/"),
	}
	for _, c := range candidates {
		if found, ok = detectViewsDir(c); ok {
			return found, ok
		}
	}
	return "", false
}

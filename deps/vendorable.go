// Copyright 2016 govend. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

package deps

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/talalashraf/govend/deps/semver"
)

// Vendorable ensures the current local setup is conducive to vendoring.
//
// If the current version of Go cannot be parsed, then trust it supports
// vendoring, but display a message if verbose is true.
func Vendorable(verbose bool) error {
	err := checkGopath()
	if err != nil {
		return err
	}

	go15, _ := semver.New("1.5.0")
	go16, _ := semver.New("1.6.0")
	go17, _ := semver.New("1.7.0")

	version, err := semver.New(strings.TrimPrefix(runtime.Version(), "go"))
	if err != nil {
		if verbose {
			fmt.Printf("\n%s\n", err)
		}
		return nil
	}

	if version.LessThan(go15) {
		return errors.New("vendoring requires Go versions 1.5+")
	}

	if version.GreaterThanEqual(go15) && version.LessThan(go16) {
		if os.Getenv("GO15VENDOREXPERIMENT") != "1" {
			return errors.New("Go 1.5.x requires 'GO15VENDOREXPERIMENT=1'")
		}
	}

	if version.GreaterThanEqual(go16) && version.LessThan(go17) {
		if os.Getenv("GO15VENDOREXPERIMENT") == "0" {
			return errors.New("Go 1.6.x cannot vendor with 'GO15VENDOREXPERIMENT=0'")
		}
	}

	return nil
}

// checkGopath checks if the current working directory has $GOPATH/src as a prefix.
func checkGopath() error {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		return errors.New("please set your $GOPATH")
	}

	// determine the current working directory and coerce it to an absolute
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cwd, err = filepath.Abs(cwd)
	if err != nil {
		return err
	}

	cwd, err = filepath.EvalSymlinks(cwd)
	if err != nil {
		return err
	}
	sep := string(filepath.Separator)
	cwd = cwd + sep

	// for each filepath in $GOPATH, check path/src
	paths := filepath.SplitList(gopath)
	for _, path := range paths {
		gosrc := filepath.Join(path, "src") + sep
		if strings.HasPrefix(cwd, gosrc) {
			return nil
		}
	}

	return errors.New("you cannot vendor packages outside of your $GOPATH/src")
}

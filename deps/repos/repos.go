// Copyright 2016 govend. All rights reserved.
// Use of this source code is governed by an Apache 2.0
// license that can be found in the LICENSE file.

// Package repos provides methods for repositories.
package repos

import "github.com/talalashraf/govend/deps/vcs"

// Repo RepoRoot
type Repo struct {
	VCS *VCS

	// URL is the repository URL, including scheme
	URL string

	// ImportPath is the import path corresponding to the root of the repository
	ImportPath string
}

// New returns a new Repo.
func New(v *VCS, url, importpath string) *Repo {
	return &Repo{
		VCS:        v,
		URL:        url,
		ImportPath: importpath,
	}
}

// ImportPath returns a new Repo.
func ImportPath(importpath string, verbose bool) (*Repo, error) {
	rr, err := vcs.RepoRootForImportPath(importpath, vcs.Secure, verbose)
	if err != nil {
		return nil, err
	}
	vcs, err := NewVCS(rr.VCS)
	if err != nil {
		return nil, err
	}
	return New(vcs, rr.Repo, rr.Root), nil
}

// ImportDynamic returns a new Repo.
func ImportDynamic(importpath string, verbose bool) (*Repo, error) {
	rr, err := vcs.RepoRootForImportDynamic(importpath, vcs.Secure, verbose)
	if err != nil {
		return nil, err
	}
	vcs, err := NewVCS(rr.VCS)
	if err != nil {
		return nil, err
	}
	return New(vcs, rr.Repo, rr.Root), nil
}

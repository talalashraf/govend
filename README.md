![govend](art/govend.png)

# govend [![GoDoc](http://godoc.org/github.com/talalashraf/govend?status.png)](http://godoc.org/github.com/talalashraf/govend) [![Build Status](http://beta.drone.io/api/badges/govend/govend/status.svg)](http://beta.drone.io/govend/govend) [![Go Report Card](https://goreportcard.com/badge/github.com/talalashraf/govend)](https://goreportcard.com/report/github.com/talalashraf/govend) [![Join the chat at https://gitter.im/govend/govend](https://badges.gitter.im/govend/govend.svg)](https://gitter.im/govend/govend?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) ![](https://img.shields.io/badge/windows-ready-green.svg)

`govend` is a simple tool to vendor Go package dependencies.
It's like `go get`, but for vendoring external or third party packages.

**govend is:**
* like `go get`, but for vendoring packages
* compatible with any project directory structure
* designed to vendor nested dependencies to the nth degree
* compatible with Go versions 1.5+.

**govend does not:**
* wrap the `go` command
* try to create a new project
* force you to lock dependency versions
* generate temporary directories or files
* alter any Go environment variables, including `$GOPATH`

# Install

```bash
$ go get -u github.com/talalashraf/govend
```

# Usage

### Small Project Usage

```bash
# ...                                        # code, code, code
$ govend github.com/gorilla/mux              # vendor a dependency
# ...                                        # code, code, code
$ go build                                   # go tools work normally
$ govend -u github.com/gorilla/mux           # updated a dependency
# ...                                        # code, code, code
$ go build                                   # go tools work normally
```

### Big Project Usage

```bash
# ...                  # code, code, code
$ govend -v            # scan your project and download all dependencies
# ...                  # code, code, code
$ go build             # go tools work normally
$ govend -u            # scan your project and update all dependencies
# ...                  # code, code, code
$ go build             # go tools work normally
```

### Team Usage
Sarah:
```bash
$ git init                     # start git project
# ...                          # code, code, code
$ govend -v -l                 # scan your project, download all dependencies,
                               # and create a vendor.yml file to lock
                               # dependency versions
# ...                          # code, code, code
$ go build                     # go tools work normally
$ govend -v -u                 # scan your project, update all dependencies,
                               # and update the vendor.yml revision versions
# ...                          # code, code, code
$ git push                     # push code to github
$ go build                     # go tools work normally
```

Mike:
```bash
$ git clone url       # grab all the code Sarah pushed
$ govend -v           # download all the dependencies in the vendor.yml file
                      # and use the same revision versions Sarah is using
$ go build            # build the exact same binary Sarah has
# ...                 # code, code, code
```

# Verbose Mode
As with most unixy programs, no news is good news.
Therefore, unless something goes wrong `govend` will not print anything to the
terminal.
If you want to see progress or proof something is happening use the `-v` flag
to print out package import paths as they are downloaded and vendored.

# The Dependency Tree
You can use the `-t` or `--tree` flag to view a rendition of the package
dependency tree. This is helpful to visualize and understand what packages your
dependencies rely on as well.

# Explicitly Vendor A Package
You can explicitly tell `govend` to vendor one or more packages.
It works the same way as `go get` but instead of running:

```Bash
$ go get github.com/gorilla/mux
```

which will download the gorilla `mux` package into your `$GOPATH`, run:

```Bash
$ govend github.com/gorilla/mux
```

which will download the gorilla `mux` package into your local project `vendor/`
directory.
If you want `govend` to download more than one package, just tack them on.
For example, you might want to vendor the gorilla `mux`, `http`, and
`securecookie` packages like so:

```Bash
$ govend github.com/gorilla/mux github.com/gorilla/http github.com/gorilla/securecookie
```

# Explicitly Update Packages

To update a package that has already been vendored, simply use the `-u` network
update flag.
This flag has the same meaning as `go get -u` and will always use the network
to download a fresh copy of the dependency.

To update the gorilla `mux` package in your `$GOPATH` you would run:

```Bash
$ go get -u github.com/gorilla/mux
```

To update the gorilla `mux` package in your local project `vendor/` directory
run:

```Bash
$ govend -u github.com/gorilla/mux
```

# Vendor Packages Automatically

It would get old to ask `govend` to download and vendor each individual package
when working on large Go projects.
Thankfully `govend` can scan your project source code and identify dependencies
for you.

`govend` assumes you want this behavior when no packages are explicitly
provided:

```Bash
$ cd project/root

$ govend
```

You can also show dependencies as they are vendored with the `-v` flag:

```Bash
$ cd project/root

$ govend -v
github.com/kr/fs
github.com/BurntSushi/toml
github.com/spf13/cobra
github.com/inconshreveable/mousetrap
github.com/spf13/pflag
gopkg.in/yaml.v2
gopkg.in/check.v1
```

If you would like to update all vendored packages in a project use the `-u`
flag:

```Bash
$ cd project/root

$ govend -v -u
github.com/kr/fs
github.com/BurntSushi/toml
github.com/spf13/cobra
github.com/inconshreveable/mousetrap
github.com/spf13/pflag
gopkg.in/yaml.v2
gopkg.in/check.v1
```

# Lock Packages

The command `govend` only scans for external packages and downloads them to the
`vendor/` directory in your project. You may need more control over versioning
your dependencies so that reliable reproducible builds are possible.

`govend` can save the path and commit revisions of each repository downloaded
in a `vendor.yml` file.
This is called vendor locking.
The format of the file can be specified to be `JSON` or `TOML`, `YAML` is used
by default.
Usually this file is located in the root directory of your project and should
be included in your version control system.

To generate a `vendor.yml` file use the `-l` flag:

```Bash
$ cd project/root

$ govend -v -l
github.com/kr/fs
github.com/BurntSushi/toml
github.com/spf13/cobra
github.com/inconshreveable/mousetrap
github.com/spf13/pflag
gopkg.in/yaml.v2
gopkg.in/check.v1
```

The resulting project structure should look something like:

```Bash
.
├── ...
├── code
├── README.md
├── vendor
└── vendor.yml
```

The contents of the generated `vendor.yml` file in this example would be:

```yaml
vendors:
- path: github.com/BurntSushi/toml
  rev: f772cd89eb0b33743387f826d1918df67f99cc7a
- path: github.com/inconshreveable/mousetrap
  rev: 76626ae9c91c4f2a10f34cad8ce83ea42c93bb75
- path: github.com/kr/fs
  rev: 2788f0dbd16903de03cb8186e5c7d97b69ad387b
- path: github.com/spf13/cobra
  rev: 65a708cee0a4424f4e353d031ce440643e312f92
- path: github.com/spf13/pflag
  rev: 7f60f83a2c81bc3c3c0d5297f61ddfa68da9d3b7
- path: gopkg.in/check.v1
  rev: 4f90aeace3a26ad7021961c297b22c42160c7b25
- path: gopkg.in/yaml.v2
  rev: f7716cbe52baa25d2e9b0d0da546fcf909fc16b4
```


You can now ignore the large `vendor/` directory and pass the small `vendor.yml`
file to your buddy.
Your buddy can run `$ govend` and will get the exact same dependency versions
as specified by `vendor.yml`.

This is how a team of developers can ensure reproducible builds if they do not
want to check the `vendor/` directory into a version control system.

> Note: It is still a best practice to check in the `vendor/` directory to your
VCS.

# Update Locked Packages

If you would like to update a particular vendored package to its latest version
use the `-u` flag:

```Bash
$ govend -u github.com/gorilla/mux
```

If you would like to update all the vendored packages to their latest versions
run:

```Bash
$ govend -u
```

If you want to update a particular vendored package to a particular revision,
update the relevant `rev:` value inside the `vendor.yml` file.
Then to update to that specific revision hash run:

```Bash
$ govend -l
```

# Prune Packages
Sometimes large repositories must be downloaded to satisfy a singular package
dependency. This generates a lot of dead files if you are checking in `vendor/`
to your VCS.

The `--prune` flag will use the known dependency tree to remove or prune out
all unused package leaving you with only relevant code.

> Note: When pruning occurs it removes not only unused packages, but also files
that start with `.`, `_` and files that end with `_test.go`. This is because
pruning is highly likely to break third party tests.

# Hold onto Packages
If you stop using or importing a package path in your project code, `govend`
will remove that package from your `vendor.yml`. Its how `govend` cleans up
after you and keeps `vendor.yml` tidy.

The `--hold` flag will tell `govend` to keep that dependency, even if its not
being used as an import by your project. This is great for versioning tooling
that you might want to ship with a project.

```Bash
$ govend --hold github.com/hashicorp/terraform
```

> Note: When using the `--hold` flag a manifest file like `vendor.yml` is
generated.  Essentially, opting into `--hold` is also opting into `--lock`.
Also because of the nature of hold, repos using hold cannot be pruned.

# Vendor Report Summary
If you would like to get a report summary of the number of unique packages
scanned, skipped and how many repositories were downloaded, run `govend -v -r`.

```bash
→ govend -v -r
github.com/kr/fs
github.com/BurntSushi/toml
github.com/spf13/cobra
github.com/inconshreveable/mousetrap
github.com/spf13/pflag
gopkg.in/yaml.v2
gopkg.in/check.v1

packages scanned: 7
packages skipped: 0
repos downloaded: 7
```

# Vendor Scan
You may want to scan your code to determine how many third party dependencies
are
in your project.
To do so run `govend -s <path/to/dir>`. You can also specify different output
formats.

**TXT**
```bash
$ govend -s packages
github.com/kr/fs
gopkg.in/yaml.v2
```

**JSON**
```bash
$ govend -s -f json packages
[
  "github.com/kr/fs",
  "gopkg.in/yaml.v2"
]
```

**YAML**
```bash
$ govend -s -f yaml packages
- github.com/kr/fs
- gopkg.in/yaml.v2
```
**XML**
```bash
$ govend -s -f xml packages
<string>gopkg.in/yaml.v2</string>
<string>github.com/kr/fs</string>
```

# More Flags

You can run `govend -h` to find more flags and options.

# Vendor Supported Go Versions

  * Go 1.4 or less - Go does not support vendoring
  * Go 1.5 - vendor via `GO15VENDOREXPERIMENT=1`
  * Go 1.6 - vendor unless `GO15VENDOREXPERIMENT=0`
  * Go 1.7+ - vendor always despite the value of `GO15VENDOREXPERIMENT`

For further explanation please read https://golang.org/doc/go1.6#go_command.

# Windows Support
`govend` works on Windows, but please report any bugs.

# Contributing
Simply fork the code and send a pull request.

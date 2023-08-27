# GRT (GO Repo Template)

[![CI][badge-build]][build]
[![GoDoc][go-docs-badge]][go-docs]
[![GoReportCard][go-report-card-badge]][go-report-card]
[![License][badge-license]][license]

## Motivation

This repo aims to save some bootstrapping time when creating a new GO repo.
By using it as a template when creating new projects, you will gain:

- A set of pre-commit hooks to test that:
    * Your GO code is correctly formatted (using `gofmt`)
    * Your package imports adhere to GO standards
    * There are no linting issues
    * There are no inefficient assignments
- GH actions CI configuration for:
    * Testing code pushes and pull requests (using the same tools as above)
    * Releasing assets with [GoReleaser](https://goreleaser.com)
- A template for CLI GO applications (usage, option processing, etc)
- A README template with links and badges for: 
    * CI
    * License
    * GO docs
    * GO report card

## Configuration files
- [.pre-commit-config.yaml](.pre-commit-config.yaml): a set of pre-commit hooks relevant to GO code
- [.goreleaser.yaml](.goreleaser.yaml): [GoReleaser](https://goreleaser.com/quick-start) configuration file
- [.github/workflows/go.yml](.github/workflows/go.yml): GH action for post push and pull request testing
- [.github/workflows/release.yml](.github/workflows/release.yml): GH action to trigger upon tag creation

## Bootstrapping steps
- Fork this repo
- Under `https://github.com/YOUR_ORG/grt/settings`, make sure `Template repository` is checked
- When creating a new GO repo, select `grt` as the `Repository template`
- Clone the new repo
- Set the desired LICENSE (default is MIT)

### pre-commit hooks

Install the needed GO packages:
```sh
go install github.com/go-critic/go-critic/cmd/gocritic@latest
go install golang.org/x/tools/cmd/goimports@latest
go install golang.org/x/lint/golint@latest
go install github.com/gordonklaus/ineffassign@latest
```

Install the `pre-commit` util:
```sh
pip install pre-commit
```

Generate `.git/hooks/pre-commit`:
```sh
pre-commit install
```

Following that, these tests will run every time you invoke `git commit`:
```sh
go fmt...................................................................Passed
go imports...............................................................Passed
go vet...................................................................Passed
go lint..................................................................Passed
go-critic................................................................Passed
shellcheck...............................................................Passed
```

To manually run all tests on all repo files, invoke:

```sh
pre-commit run --all-files
```

## GoReleaser

With GoReleaser, you can:

- Cross-compile your Go project
- Release to GitHub, GitLab and Gitea
- Create nightly builds
- Create Docker images and manifests
- Create Linux packages and Homebrew taps
- Sign artifacts, checksums and container images
- Announce new releases on Twitter, Slack, Discord and others
- Generate SBOMs (Software Bill of Materials) for binaries and container images

This repo includes a [basic GoReleaser config](.goreleaser.yaml) that will produce binaries for Linux, Darwin (what
people refer to as MacOS and shouldn't), FreeBSD and NetBSD. You can tweak it as you please but one necessary change is
in [line 21](https://github.com/jessp01/grt/blob/master/.goreleaser.yaml#L21) where `main` should point to your entry
file. Also included, is [.github/workflows/release.yml](.github/workflows/release.yml), a GH action to trigger upon tag creation. This file will work out of the box but of course, you may edit it to your liking.


## Demo files

The below files are mostly included for demonstration purposes. They should be edited to reflect your project.

### grt.go

Includes only one function: `ReplaceTokens()` which replaces tokens in input file and outputs to new file. This is invoked
from `cmd/grt.go` and `grt_test.go`. Once you've gone through the basic bootstrapping, this file should be removed from
your template repo. 

### cmd/grt.go

This file serves a dual purpose:
- It demonstrate the use of the `github.com/urfave/cli` package for processing CLI args (both long and short options are
  supported)
- It assists in replacing tokens in the `README.tmpl.md` template file to produce a README skeleton for your project.

After you're done viewing this README, take the below actions:
 
- Run `go run cmd/grt.go -i README.tmpl.md -o README.md -r github.com/YOUR_ORG/YOUR_REPO_NAME`
- Edit README.md and fill in the different sections

Since this file includes metadata that is likely to be common to all your projects (author name, email and organisation
name, etc), you'll probably want to edit it on the repo you've designated as your template (the clone of the `grt`
repo).

### grt_test.go

Included to illustrate the writing of unit tests which can be called by invoking `go test -v`. This contents of this
file should be completely replaced to reflect tests for your newly created repo.

## Contributing

Contributions are, of course, welcome. Please submit a pull request.

## License

Code is under MIT License.

[license]: ./LICENSE
[badge-license]: https://img.shields.io/github/license/jessp01/grt.svg?v=0.21.0
[go-docs-badge]: https://godoc.org/github.com/jessp01/grt?status.svg
[go-docs]: https://godoc.org/github.com/jessp01/grt
[go-report-card-badge]: https://goreportcard.com/badge/github.com/jessp01/grt?v=0.21.0
[go-report-card]: https://goreportcard.com/report/github.com/jessp01/grt
[badge-build]: https://github.com/jessp01/grt/actions/workflows/go.yml/badge.svg
[build]: https://github.com/jessp01/grt/actions/workflows/go.yml

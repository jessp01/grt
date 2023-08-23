# GRT (GO Repo Template)

[![CI][badge-build]][build]
[![GoDoc][go-docs-badge]][go-docs]
[![GoReportCard][go-report-card-badge]][go-report-card]
[![License][badge-license]][license]

## Motivation

This repo aims to save some boostrapping time when creating a new GO repo.

## Contents
- [.pre-commit-config.yaml](.pre-commit-config.yaml): a set of pre-commit hooks relevant to GO code
- [.goreleaser.yaml](.goreleaser.yaml): [goreleaser](https://goreleaser.com/quick-start) configuration file
- [.github/workflows/go.yml](.github/workflows/go.yml): GH action for post push and pull request testing
- [.github/workflows/release.yml](.github/workflows/release.yml): GH action to trigger upon tag creation

## Bootstrapping steps
- Fork this repo
- Under `https://github.com/YOUR_ORG/grt/settings`, make sure `Template repository` is checked
- When creating a new GO repo, select `grt` as the `Repository template`
- Clone the new repo
- Run `go run cmd/grt.go -i README.tmpl.md -o README.md -r github.com/YOUR_ORG/YOUR_REPO_NAME`
- Edit README.md and fill in the different sections
- Set the desired LICENSE (default is AGPLv3)

### pre-commit hooks

### grt.go

### grt_test.go

### cmd/grt.go

### Goreleaser

[license]: ./LICENSE
[badge-license]: https://img.shields.io/github/license/jessp01/grt.svg
[go-docs-badge]: https://godoc.org/github.com/jessp01/grt?status.svg
[go-docs]: https://godoc.org/github.com/jessp01/grt
[go-report-card-badge]: https://goreportcard.com/badge/github.com/jessp01/grt
[go-report-card]: https://goreportcard.com/report/github.com/jessp01/grt
[badge-build]: https://github.com/jessp01/grt/actions/workflows/go.yml/badge.svg
[build]: https://github.com/jessp01/grt/actions/workflows/go.yml

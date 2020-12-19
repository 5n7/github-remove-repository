# GitHub Remove Repository

<p align="center">
  <a href="https://github.com/skmatz/github-remove-repository/actions?query=workflow%3Arelease">
    <img
      src="https://github.com/skmatz/github-remove-repository/workflows/release/badge.svg"
      alt="release"
    />
  </a>
  <a href="https://goreportcard.com/report/github.com/skmatz/github-remove-repository">
    <img
      src="https://goreportcard.com/badge/github.com/skmatz/github-remove-repository"
      alt="go report card"
    />
  </a>
  <a href="./LICENSE">
    <img
      src="https://img.shields.io/github/license/skmatz/github-remove-repository"
      alt="license"
    />
  </a>
  <a href="./go.mod">
    <img
      src="https://img.shields.io/github/go-mod/go-version/skmatz/github-remove-repository"
      alt="go version"
    />
  </a>
  <a href="https://github.com/skmatz/github-remove-repository/releases/latest">
    <img
      src="https://img.shields.io/github/v/release/skmatz/github-remove-repository"
      alt="release"
    />
  </a>
</p>

## Install

### Binary

Download from [GitHub Releases](https://github.com/skmatz/github-remove-repository/releases).

### Source

```sh
go get github.com/skmatz/github-remove-repository/...
```

## Usage

```bash
> ghrm skmatz github-remove-repository
```

```console
> ghrm --help

ghrm is the GitHub repository remover.

Usage:
  ghrm [flags]
  ghrm [command]

Available Commands:
  completion  Output shell completion (bash/fish/powershell/zsh)
  help        Help about any command
  token       Prompt for the GitHub access token
  version     Show version

Flags:
  -f, --force     skip confirmation
  -h, --help      help for ghrm
  -V, --version   show version

Use "ghrm [command] --help" for more information about a command.
```

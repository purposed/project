# Project
Project is an opinionated repository management tool. It believes your source code should live
in a tree stemming from a single root (`~/` by default) (similar to a `GOPATH`).

## Installation
Installation is simple, simply:
```shell
$ go get github.com/dalloriam/project
```

## Usage

### Fetching repositories

```shell
$ project clone git@github.com/some/repo.git
```

Project will fetch the source code & clone it in the `${ROOT}/src/github.com/some/repo/` directory.
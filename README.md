# Project
Project is an opinionated repository management tool. It believes your source code should live
in a tree stemming from a single root (`~/` by default) (similar to a `GOPATH`).

## Installation
Installation is simple, simply:
```shell
$ go get github.com/purposed/project
```

## Configuration
Project is configured via environment variables.

* `PURPOSED_OWNER` : Default owner of  repositories. (default: `${USER}`)
* `PURPOSED_SOURCEPROVIDER` : Default source hosting provider to use (default: `github.com`)
* `PURPOSED_SOURCEROOT` : Default directory for the source root. (default: `~/`)

## Usage

### Listing managed projects

```shell
$ project list
```

###  Syncing managed projects

```shell
$ project sync
```

### Fetching repositories

```shell
$ project clone github.com/user/project
```

Project will fetch the source code & clone it in the `${ROOT}/src/github.com/user/project/` directory.

# 4210 Final Project

[![Go Report Card](https://goreportcard.com/badge/github.com/kangjoa/gostart)](https://goreportcard.com/report/github.com/kangjoa/gostart)

Final project description and requirements found [here](https://github.com/Tech-at-DU/ACS-4210-Strongly-Typed-Languages?tab=readme-ov-file#%66%69%6E%61%6C%2D%70%72%6F%6A%65%63%74).

## Description

gostart is a CLI tool that instantly generates standardized, lightweight Go project boilerplate.

```zsh
new-project/
├─ .gitignore
├─ go.mod
├─ LICENSE
├─ main.go
└─ README.md
```

## How to use

### Install

Navigate to the directory where you want to start your Go project.

```zsh
go install github.com/kangjoa/gostart@latest
```

### Create a new project

```zsh
gostart init new-project
```

### Create a new project with author and email added to README.md

```zsh
gostart init new-project --author "Your name" --email "yourEmail@email.com"
```

### Save custom default author and email values to your home directory in a yaml file (persists data in a file). Your new projects will automatically contain these values.

```zsh
gostart config set author "Your name"
gostart config set email "yourEmail@email.com"

# Check current configuration values:
gostart config
```

### Get help

```zsh
gostart --help
gostart init --help
gostart config --help
```

## Testing

2 table-driven tests and 1 benchmark test are in `cmd/init_test.go`. To run tests, navigate to that directory and then run:

```zsh
go test
```

Get verbose test output

```zsh
go test -v
```

Get test coverage

```zsh
go test -cover
```

Run the benchmark test

```zsh
go test -bench=
```

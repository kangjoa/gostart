# 4210 Final Project

Final project description and requirements found [here](https://github.com/Tech-at-DU/ACS-4210-Strongly-Typed-Languages?tab=readme-ov-file#%66%69%6E%61%6C%2D%70%72%6F%6A%65%63%74).

## Description

gostart is A CLI tool that instantly generates standardized, lightweight Go project boilerplate.

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

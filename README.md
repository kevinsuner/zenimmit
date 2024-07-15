# Zenimmit

A configurable cli tool to help you write better git commit messages.

[![Go version](https://img.shields.io/github/go-mod/go-version/kevinsuner/zenimmit)](https://github.com/kevinsuner/zenimmit/blob/master/go.mod)
[![License](https://img.shields.io/github/license/kevinsuner/zenimmit)](https://github.com/kevinsuner/zenimmit/blob/master/LICENSE)

## Installation

- Using `go install`
    ```bash
    go install github.com/kevinsuner/zenimmit@latest
    ```
    **NOTE**: Remember to add the `/go/bin` folder to your path


- Using a pre-built binary
    - Download the latest binary release [here](https://github.com/kevinsuner/zenimmit/releases)
    - Move the binary to your folder of choice
    - Make the binary executable
    - Add the downloaded binary or the folder containing it to your path


- Using the source code
  - Clone this repository on your machine
  - From the root of the cloned directory run `go build -o zenimmit`
  - Move the binary to your folder of choice
  - Add the newly built binary or the folder containing it to your path

## Usage

Zenimmit lets you define your own `types` and `scopes` for the commit through
a file named `zenimmit.yaml`, which needs to be placed at the root of your project,
and looks like the following:

```yaml
types: [
  "build",
  "ci",
  "docs",
  "feat",
  # ...
]

scopes: [
  "components",
  "stepper",
  "core",
  "config",
  # ...
]
```

Although it is customizable I'd recommend sticking with the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) 
guide for the `types` part. 

Following that it's just a matter of staging some files, running the `zenimmit`
binary and going through each step, at the end Zenimmit will use the output
produced by the `git commit` command as an indicator of successful execution.

## Credits

Zenimmit makes use of a variety of open-source projects including:

- [github.com/golang/go](https://github.com/golang/go)
- [github.com/charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)
- [github.com/charmbracelet/bubbles](https://github.com/charmbracelet/bubbles)
- [github.com/spf13/viper](https://github.com/spf13/viper)
- [github.com/conventional-commits/conventionalcommits.org](https://github.com/conventional-commits/conventionalcommits.org)
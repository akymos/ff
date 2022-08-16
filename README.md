# ff

![Release](https://github.com/akymos/ff/actions/workflows/release.yml/badge.svg?branch=)
[![Go Report Card](https://goreportcard.com/badge/github.com/akymos/ff)](https://goreportcard.com/report/github.com/akymos/ff)


ff is a command-line tool to manage favorite folders, creating an alias, to be used via shell directly with the cd command.

[![asciicast](https://asciinema.org/a/UCxUx5TkKEUEitKyg8FEUZFCi.svg)](https://asciinema.org/a/UCxUx5TkKEUEitKyg8FEUZFCi)

* [Installation](#installation)
  * [Prebuilt binary](#prebuilt-binary)
  * [Update](#update)
* [Usage](#usage)
* [Todo](#todo)

# Features
- Quickly create the alias of the folder you are in.
- Manage many aliases (tested with 200+ aliases)
- In case of conflict between an alias and a directory present, gives priority to the second.
- Once you have created the alias you do not have to close the terminal to be able to use it immediately.

# Installation
## Prebuilt binary
Download the prebuilt binary from [here (stable)](https://github.com/akymos/ff/releases/latest) or [here (all with pre-release)](https://github.com/akymos/ff/releases) and run the following command:
```bash
tar -xf <prebuilt_archive> ff && sudo mv ff /usr/local/bin
```
Next add the following line to `~/.zshrc` or `~/.bashrc`: 
```bash
source "$(ff alias)"
```
## Update
To update the installed version of ff to the latest stable version, run the following command:
```bash
ff self-update
```

# Usage
```bash
ff is a command-line tool to manage favorite folders, creating an alias,
to be used via shell directly with the cd command.

Usage:
  ff [command]

Examples:
$ ff add alias_name
$ ff add alias2 /tmp
$ ff update alias_name /var

Available Commands:
  add         Create a directory alias
  alias       Prints out the path to the alias file
  delete      Interactively allows you to delete an alias
  get         Return the raw path of an alias
  help        Help about any command
  list        List saved aliases
  panic       WARINING!! This delete all saved data
  self-update Update ff to the latest version
  update      Update a directory alias

Flags:
  -h, --help      help for ff
  -v, --version   version for ff

Use "ff [command] --help" for more information about a command.
```

# TODO
- [ ] clean up the code
- [x] make a better readme
- [ ] windows support
- [ ] ......

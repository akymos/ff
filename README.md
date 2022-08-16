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

# Installation
## Prebuilt binary
Download the prebuilt binary from [here](https://github.com/akymos/ff/releases/latest) and run the following command:
```bash
tar -xf <prebuilt_archive> ff && sudo mv ff /usr/local/bin
```
Nest add the following line to `~/.zshrc` or `~/.bashrc`: 
```bash
source "$(ff alias)"
```
## Update
For update the installed version of ff, run the following command:
```bash
ff self-update
```

# Usage
```bash
$ ff
ff is a command-line tool to manage favorite folders, creating an alias,
to be used via shell directly with the cd command.

Usage:
  ff [command]

Available Commands:
  add         Create a directory alias
  alias       Prints out the path to the alias file.
  delete      Delete a directory alias.
  get         Return the raw path of an alias.
  help        Help about any command
  list        List saved aliases.
  panic       WARINING!! This delete all saved data
  self-update Update ff to the latest version.
  update      Update a directory alias

Flags:
  -h, --help      help for ff
  -v, --version   version for ff

Use "ff [command] --help" for more information about a command.
```

# TODO
- [ ] clean up the code
- [ ] make a better readme
- [ ] windows support
- [ ] ......

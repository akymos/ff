# ff
ff is a command-line tool to manage favorite folders, creating an alias, to be used via shell directly with the cd command.

[![asciicast](https://asciinema.org/a/UCxUx5TkKEUEitKyg8FEUZFCi.svg)](https://asciinema.org/a/UCxUx5TkKEUEitKyg8FEUZFCi)

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
  update      Update a directory alias

Flags:
  -h, --help      help for ff
  -v, --version   version for ff

Use "ff [command] --help" for more information about a command.
```

# TODO
- [ ] clean up the code
- [X] use github workflow to build the binary
- [ ] use github.com/manifoldco/promptui for edit command
- [ ] make a better readme
- [ ] ......

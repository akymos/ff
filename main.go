package main

import (
	"fmt"
	"os"
)

func main() {
	err := initConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	arg := ""
	arg1 := ""
	arg2 := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	if len(os.Args) > 2 {
		arg1 = os.Args[2]
	}
	if len(os.Args) > 3 {
		arg2 = os.Args[3]
	}

	switch arg {
	case "add", "a", "+":
		if arg2 == "" || arg2 == "." || arg2 == "./" {
			p, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			arg2 = p
		}
		err = localDb.add(arg1, arg2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = populateAlias()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		writeDb()
		fmt.Printf("Alias %s added for folder %s\n", arg1, arg2)
		fmt.Printf("Now run \n%s\nor restart the shell\n", "source \"$(ff alias)\"")
	case "delete", "del", "d", "-":
		err = localDb.del(arg1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = populateAlias()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		writeDb()
		fmt.Printf("Alias %s deleted\n", arg1)
		fmt.Printf("Now run \n%s\nor restart the shell\n", "source \"$(ff alias)\"")
	case "update", "upd", "u":
		if arg2 == "" || arg2 == "." || arg2 == "./" {
			p, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
			}
			arg2 = p
		}
		localDb.update(arg1, arg2)
		err = populateAlias()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		writeDb()
		fmt.Printf("Alias %s updated to %s\n", arg1, arg2)
		fmt.Printf("Now run \n%s\nor restart the shell\n", "source \"$(ff alias)\"")
	case "list", "ls", "l":
		list := localDb.findAll()
		if len(list) == 0 {
			fmt.Println("Empty")
			os.Exit(0)
		}
		for k, v := range list {
			fmt.Printf("Alias: %s -> Folder: %s\n", k, v)
		}
		os.Exit(0)
	case "alias":
		err = populateAlias()
		if err != nil {
			return
		}
		fmt.Println(getAlias())
	default:
		fmt.Printf(`ff is a command-line tool to manage favorite folders, creating an alias,
to be used via shell directly with the cd command.

Usage:
  ff [command] [alias] [?path]

Available Commands:
  add, a, +           Add an alias.
  delete, del, d, -   Remove an alias.
  update, upd, u      Update an alias.
  list, ls, l         List all aliases.
  alias               Prints out the path to the alias file.

Eamples:
  ff add current_folder_alias
  ff + alias_name /path/to/folder
  ff update alias_name /path/to/new/folder
  ff ls
  ff del alias_name
`)
	}
	os.Exit(0)
}

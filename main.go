package main

import "github.com/akymos/ff/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}

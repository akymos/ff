package main

import "github.com/akymos/ff/cmd"

var date string
var commit string
var version string

func main() {
	cmd.Execute(version, date, commit)
}

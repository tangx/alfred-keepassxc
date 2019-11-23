package main

import (
	"os"

	"github.com/tangx/alfred-keepassxc/cmd"
)

func main() {
	if len(os.Args) > 1 {
		cmd.Main(os.Args[1:])
	}
}

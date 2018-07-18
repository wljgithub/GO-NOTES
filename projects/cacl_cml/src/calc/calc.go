package main

import (
	"fmt"
	"os"
	"strconv"
)

func usage() {

}
func main() {

	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		usage()
		return

	}
	switch args[0] {
	case "add":
		if len(args) != 3 {
			fmt.Println("")
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	fmt.Println("to do")
}

type command struct {
	Name string
	Desc string
}

func usage() {
	fmt.Fprintf(os.Stderr, "gone-cli is a cli tool for gone.\n\n")
	fmt.Fprintf(os.Stderr, "Usage:\n\n")
	fmt.Fprintf(os.Stderr, "\tgone-cli <command> [arguments]\n\n")
	fmt.Fprintf(os.Stderr, "The commands are:\n\n")

	cmds := []command{
		{Name: "init", Desc: "init gone code"},
	}

	for _, cmd := range cmds {
		fmt.Fprintf(os.Stderr, "\t%-10s %s\n", cmd.Name, cmd.Desc)
	}

	fmt.Fprintf(os.Stderr, "\nFor more information run\n\n")
	fmt.Fprintf(os.Stderr, "\tgone-cli help <command>\n\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
	// os.Exit(2)
}

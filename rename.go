package main

import "fmt"
import "os"

func Rename(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Missing argument: <name>")
		os.Exit(1)
	}

	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, "Missing argument: <new_name>")
		os.Exit(1)
	}

	name := args[0]
	newName := args[1]
	client, err := newClient()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if !directoryExists(client, name) {
		fmt.Fprintln(os.Stderr, "Repository", name, "doesn't exist.")
		os.Exit(1)
	}

	if directoryExists(client, newName) {
		fmt.Fprintln(os.Stderr, "Repository", newName, "already exists.")
		os.Exit(1)
	}

	command := "/bin/mv " + name + ".git " + newName + ".git"
	_, stderr, err := client.Run(command)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't rename repository.")
		fmt.Fprintln(os.Stderr, stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("Repository renamed from", name, "to", newName+".")
}

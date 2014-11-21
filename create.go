package main

import "os"
import "fmt"

func Create(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Missing argument: <name>")
		os.Exit(1)
	}

	name := args[0]
	client, err := newClient()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if repositoryExists(client, name) {
		fmt.Fprintln(os.Stderr, "Repository", name, "already exists.")
		os.Exit(1)
	}

	command := "mkdir -p " + name + ".git && cd " + name + ".git && git --bare init"
	_, stderr, err := client.Run(command)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create repository.")
		fmt.Fprintln(os.Stderr, err.Error(), stderr)
		os.Exit(1)
	}

	fmt.Println("Repository", name, "created.")
}

package commands

import "fmt"
import "os"

func Destroy(args []string) {
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

	if !directoryExists(client, name) {
		fmt.Fprintln(os.Stderr, "Repository", name, "doesn't exist.")
		os.Exit(1)
	}

	command := "rm -rf " + name + ".git"
	_, stderr, err := client.Run(command)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to destroy repository.")
		fmt.Fprintln(os.Stderr, err.Error(), stderr)
		os.Exit(1)
	}

	fmt.Println("Repository", name, "destroyed.")
}

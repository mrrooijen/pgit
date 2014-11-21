package main

import "fmt"
import "os"

func main() {
	cmd, args, ok := parseInput(os.Args)

	if !ok {
		Help()
		os.Exit(2)
	}

	switch cmd {
	case "list":
		List()
	case "clip":
		Clip(args)
	case "create":
		Create(args)
	case "rename":
		Rename(args)
	case "destroy":
		Destroy(args)
	case "version":
		Version()
	case "help":
		Help()
	default:
		fmt.Fprintln(os.Stderr, "Command", "`"+cmd+"`", "not found in pgit.")
		fmt.Fprintln(os.Stderr, "Run `pgit help` to see a list of available commands.")
		os.Exit(2)
	}
}

func parseInput(osArgs []string) (cmd string, args []string, ok bool) {

	if len(osArgs) < 2 {
		return
	}

	cmd = osArgs[1]
	args = osArgs[2:]
	ok = true

	return
}

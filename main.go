package main

import "fmt"
import "os"
import "github.com/meskyanichi/pgit/commands"

func main() {
	cmd, args, ok := parseInput(os.Args)

	if !ok {
		commands.Help()
		os.Exit(2)
	}

	switch cmd {
	case "list":
		commands.List()
	case "clip":
		commands.Clip(args)
	case "create":
		commands.Create(args)
	case "rename":
		commands.Rename(args)
	case "destroy":
		commands.Destroy(args)
	case "version":
		commands.Version()
	case "help":
		commands.Help()
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

package commands

import "fmt"

func Help() {
	fmt.Println(`
Lightweight command-line utility for managing private repositories on a server.

Usage:

  pgit command [arguments]

Commands:

  list                         list all repositories
  clip                         clip the url of a repository to your clipboard
  create <name>                create a new repository
  rename <name> <new_name>     rename an existing repository
  destroy <name>               destroy an existing repository
  version                      display the current pgit version
  help                         display this help screen

https://github.com/meskyanichi/pgit
	`)
}

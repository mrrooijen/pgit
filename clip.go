package main

import "fmt"
import "os"
import "os/exec"
import "strings"
import "errors"
import "bytes"

type lookPath func(string) (string, error)

func Clip(args []string) {
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

	defer client.Close()

	if !repositoryExists(client, name) {
		fmt.Fprintln(os.Stderr, "Repository", name, "doesn't exist.")
		os.Exit(1)
	}

	fullUrl, err := getFullUrl(name)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	cmd, stderr, err := newClipCmd(fullUrl)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't copy to clipboad.")
		fmt.Fprintln(os.Stderr, stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println(fullUrl, "copied to clipboard!")
}

func newClipCmd(input string) (cmd *exec.Cmd, stderr bytes.Buffer, err error) {
	if input == "" {
		err = errors.New("Can't clip without a value.")
		return
	}

	clipUtilWithArgs, err := getClipUtilWithArgs(exec.LookPath)

	if err != nil {
		return
	}

	cmd = exec.Command(clipUtilWithArgs[0], clipUtilWithArgs[1:]...)
	cmd.Stdin = strings.NewReader(input)
	cmd.Stderr = &stderr
	return
}

func getClipUtilWithArgs(fn lookPath) (cmdWithArgs []string, err error) {
	if path, err := fn("pbcopy"); err == nil {
		cmdWithArgs = []string{path}
	}

	if path, err := fn("xsel"); err == nil {
		cmdWithArgs = []string{path, "-i", "-b"}
	}

	if path, err := fn("xclip"); err == nil {
		cmdWithArgs = []string{path, "-selection", "clipboard"}
	}

	if len(cmdWithArgs) == 0 {
		err = errors.New(
			"Cannot copy to clipboard.\n" +
				"Couldn't find pbcopy, xsel or xclip on this machine.")
	}

	return
}

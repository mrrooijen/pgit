package commands

import "fmt"
import "os"
import "strings"

func List() {
	client, err := newClient()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	repositories, err := getRepositories(client)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	maxLength := maxLen(repositories)

	for _, name := range repositories {
		url, err := getFullUrl(name)

		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		fmt.Println(name, addWhitespace(name, maxLength), url)
	}
}

func maxLen(slice []string) (result int) {
	for _, item := range slice {
		length := len(item)
		if length > result {
			result = length
		}
	}

	return
}

func addWhitespace(name string, maxLength int) string {
	requiredWhitespace := maxLength - len(name)
	return strings.Repeat(" ", requiredWhitespace)
}

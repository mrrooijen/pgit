package commands

import "os"
import "regexp"
import "strconv"
import "strings"
import "errors"
import "github.com/meskyanichi/simplessh"

type PgitUrl struct {
	Url  string
	User string
	Host string
	Port int
}

type CommandRunner interface {
	Run(string) (stdout string, stderr string, err error)
}

func parsePgitUrl() (pgitUrl *PgitUrl, err error) {
	url := os.Getenv("PGIT_URL")

	if url == "" {
		err = errors.New("PGIT_URL not set.")
		return
	}

	pattern := "^([^@]+)@([^:]+)(?::(\\d+))?$"
	compiledPattern, _ := regexp.Compile(pattern)
	match := compiledPattern.FindStringSubmatch(url)

	if len(match) == 0 {
		err = errors.New("Couldn't parse PGIT_URL, invalid format.\n" +
			"Acceptable format: <user>@<host>[:port]")
		return
	}

	port, _ := strconv.Atoi(match[3])
	if port == 0 {
		port = 22
	}

	pgitUrl = &PgitUrl{
		Url:  url,
		User: match[1],
		Host: match[2],
		Port: port,
	}

	return
}

func newClient() (client *simplessh.Client, err error) {
	pgitUrl, err := parsePgitUrl()

	if err != nil {
		return
	}

	config := &simplessh.Config{
		User: pgitUrl.User,
		Host: pgitUrl.Host,
		Port: pgitUrl.Port,
	}
	client, err = simplessh.NewClient(config)

	if err != nil {
		err = errors.New(
			"Couldn't establish a connection with " + pgitUrl.Url + ".\n" +
				err.Error())
		return
	}

	return
}

func repositoryExists(client CommandRunner, name string) bool {
	_, _, err := client.Run("ls " + name + ".git/HEAD")
	return err == nil
}

func directoryExists(client CommandRunner, name string) bool {
	_, _, err := client.Run("ls " + name + ".git")
	return err == nil
}

func getRepositories(client CommandRunner) (repositories []string, err error) {
	stdout, stderr, err := client.Run("ls")

	if err != nil {
		err = errors.New(stderr + " " + err.Error())
		return
	}

	list := strings.Fields(stdout)
	pattern, _ := regexp.Compile(".git$")

	for _, item := range list {
		if pattern.MatchString(item) {
			repositories = append(repositories, item[:len(item)-4])
		}
	}

	return
}

func getFullUrl(name string) (url string, err error) {
	user, err := parseUrlUser()

	if err != nil {
		return
	}

	urlParts := []string{
		"ssh://" + os.Getenv("PGIT_URL"),
		"home",
		user,
		name + ".git",
	}

	url = strings.Join(urlParts, "/")

	return
}

func parseUrlUser() (user string, err error) {
	pattern, _ := regexp.Compile("([^@]+)@")
	match := pattern.FindStringSubmatch(os.Getenv("PGIT_URL"))

	if len(match) == 0 {
		err = errors.New(
			"Couldn't parse PGIT_URL, invalid format." +
				"Acceptable format: <user>@<host>[:port]")
		return
	}

	user = match[1]

	return
}

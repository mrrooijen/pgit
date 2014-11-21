package main

import "os"
import "testing"
import "errors"
import "github.com/stretchr/testify/assert"

type ClientStub struct{}

func (c *ClientStub) Run(cmd string) (string, string, error) {
	if cmd == "ls" {
		return "foo.git bar baz.git", "", nil
	}

	if cmd == "ls example.git" {
		return "", "", nil
	}

	if cmd == "ls example.git/HEAD" {
		return "", "", nil
	}

	return "", "", nil
}

type ClientErrorStub struct{}

func (c *ClientErrorStub) Run(cmd string) (string, string, error) {
	if cmd == "ls" {
		return "", "ls: foo: No such file or directory", errors.New("new error")
	}

	if cmd == "ls example.git" {
		return "", "", errors.New("")
	}

	if cmd == "ls example.git/HEAD" {
		return "", "", errors.New("")
	}

	return "", "", nil
}

func TestParseUrlUser(t *testing.T) {
	assert := assert.New(t)
	var user string
	var err error

	os.Setenv("PGIT_URL", "example.com")
	user, err = parseUrlUser()
	assert.Equal(user, "")
	assert.NotNil(err)

	os.Setenv("PGIT_URL", "git@example.com")
	user, err = parseUrlUser()
	assert.Equal(user, "git")
	assert.Nil(err)
}

func TestGetFullUrl(t *testing.T) {
	assert := assert.New(t)
	var url string
	var err error

	os.Setenv("PGIT_URL", "")
	url, err = getFullUrl("example")
	assert.Equal(url, "")
	assert.NotNil(err)

	os.Setenv("PGIT_URL", "git@example.com")
	url, err = getFullUrl("example")
	assert.Equal(url, "ssh://git@example.com/home/git/example.git")
	assert.Nil(err)
}

func TestGetRepositories(t *testing.T) {
	assert := assert.New(t)
	var repositories []string
	var err error

	clientTwo := ClientErrorStub{}
	repositories, err = getRepositories(&clientTwo)
	assert.Equal(len(repositories), 0)
	assert.Equal(err.Error(), "ls: foo: No such file or directory new error")

	clientOne := ClientStub{}
	repositories, err = getRepositories(&clientOne)
	assert.Equal(repositories, []string{"foo", "baz"})
	assert.Nil(err)
}

func TestDirectoryExists(t *testing.T) {
	assert := assert.New(t)
	var ok bool

	clientTwo := ClientErrorStub{}
	ok = directoryExists(&clientTwo, "example")
	assert.Equal(ok, false)

	clientOne := ClientStub{}
	ok = directoryExists(&clientOne, "example")
	assert.Equal(ok, true)
}

func TestRepositoryExists(t *testing.T) {
	assert := assert.New(t)
	var ok bool

	clientTwo := ClientErrorStub{}
	ok = repositoryExists(&clientTwo, "example")
	assert.Equal(ok, false)

	clientOne := ClientStub{}
	ok = repositoryExists(&clientOne, "example")
	assert.Equal(ok, true)
}

func TestParsePgitUrl(t *testing.T) {
	assert := assert.New(t)
	var pgitUrl *PgitUrl
	var err error

	os.Setenv("PGIT_URL", "")
	pgitUrl, err = parsePgitUrl()
	assert.Nil(pgitUrl)
	assert.NotNil(err)

	os.Setenv("PGIT_URL", "example.com")
	pgitUrl, err = parsePgitUrl()
	assert.Nil(pgitUrl)
	assert.NotNil(err)

	os.Setenv("PGIT_URL", "git@example.com")
	pgitUrl, err = parsePgitUrl()
	assert.Equal(pgitUrl, &PgitUrl{
		Url:  "git@example.com",
		User: "git",
		Host: "example.com",
		Port: 22,
	})
	assert.Nil(err)

	os.Setenv("PGIT_URL", "git@example.com:2222")
	pgitUrl, err = parsePgitUrl()
	assert.Equal(pgitUrl, &PgitUrl{
		Url:  "git@example.com:2222",
		User: "git",
		Host: "example.com",
		Port: 2222,
	})
	assert.Nil(err)
}

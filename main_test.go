package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestParseInputList(t *testing.T) {
	assert := assert.New(t)
	osArgs := []string{"", "list"}
	cmd, args, ok := parseInput(osArgs)

	assert.Equal(cmd, "list")
	assert.Equal(len(args), 0)
	assert.Equal(ok, true)
}

func TestParseInputCreate(t *testing.T) {
	assert := assert.New(t)
	osArgs := []string{"", "create", "repo"}
	cmd, args, ok := parseInput(osArgs)

	assert.Equal(cmd, "create")
	assert.Equal(args, []string{"repo"})
	assert.Equal(ok, true)
}

func TestParseInputRename(t *testing.T) {
	assert := assert.New(t)
	osArgs := []string{"", "rename", "repo", "newrepo"}
	cmd, args, ok := parseInput(osArgs)

	assert.Equal(cmd, "rename")
	assert.Equal(args, []string{"repo", "newrepo"})
	assert.Equal(ok, true)
}

func TestParseInputError(t *testing.T) {
	assert := assert.New(t)
	osArgs := []string{""}
	cmd, args, ok := parseInput(osArgs)

	assert.Equal(cmd, "")
	assert.Equal(len(args), 0)
	assert.Equal(ok, false)
}

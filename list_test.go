package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestMaxLen(t *testing.T) {
	assert := assert.New(t)

	slice := []string{"bash", "golang", "zsh"}
	length := maxLen(slice)
	assert.Equal(length, 6)
}

func TestAddWhitespace(t *testing.T) {
	assert := assert.New(t)

	whitespace := addWhitespace("golang", 10)
	assert.Equal(whitespace, "    ")
}

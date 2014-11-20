package commands

import "testing"
import "os/exec"
import "errors"
import "github.com/stretchr/testify/assert"

func TestGetClipUtilWithArgs(t *testing.T) {
	assert := assert.New(t)
	var cmdWithArgs []string
	var err error

	pbcopyLookPath := func(path string) (string, error) {
		if path == "pbcopy" {
			return "/bin/pbcopy", nil
		}

		return "", errors.New("")
	}
	cmdWithArgs, err = getClipUtilWithArgs(pbcopyLookPath)
	assert.Equal(cmdWithArgs, []string{"/bin/pbcopy"})
	assert.Nil(err)

	xselLookPath := func(path string) (string, error) {
		if path == "xsel" {
			return "/bin/xsel", nil
		}

		return "", errors.New("")
	}
	cmdWithArgs, err = getClipUtilWithArgs(xselLookPath)
	assert.Equal(cmdWithArgs, []string{"/bin/xsel", "-i", "-b"})
	assert.Nil(err)

	xclipLookPath := func(path string) (string, error) {
		if path == "xclip" {
			return "/bin/xclip", nil
		}

		return "", errors.New("")
	}
	cmdWithArgs, err = getClipUtilWithArgs(xclipLookPath)
	assert.Equal(cmdWithArgs, []string{"/bin/xclip", "-selection", "clipboard"})
	assert.Nil(err)

	noopLookPath := func(_ string) (string, error) {
		return "", errors.New("")
	}
	cmdWithArgs, err = getClipUtilWithArgs(noopLookPath)
	assert.Equal(len(cmdWithArgs), 0)
	assert.NotNil(err)
}

func TestNewClipCmd(t *testing.T) {
	assert := assert.New(t)
	var cmd *exec.Cmd
	var err error

	cmd, _, err = newClipCmd("")
	assert.Nil(cmd)
	assert.NotNil(err)
}

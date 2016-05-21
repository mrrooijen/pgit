package simplessh

import (
	"bytes"
	"errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

type Config struct {
	User     string
	Host     string
	Port     int
	Password string
	KeyPaths []string
}

type Client struct {
	SshClient *ssh.Client
}

func NewClient(config *Config) (client *Client, err error) {
	clientConfig := &ssh.ClientConfig{
		User: config.User,
	}

	if config.Password != "" {
		password := ssh.Password(config.Password)
		clientConfig.Auth = append(clientConfig.Auth, password)
	}

	if len(config.KeyPaths) == 0 {
		keyPath := os.Getenv("HOME") + "/.ssh/id_rsa"
		key, err := makePrivateKey(keyPath)

		if err == nil {
			clientConfig.Auth = append(clientConfig.Auth, ssh.PublicKeys(key))
		}
	} else {
		keys, err := makePrivateKeys(config.KeyPaths)

		if err == nil {
			clientConfig.Auth = append(clientConfig.Auth, ssh.PublicKeys(keys...))
		}
	}

	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err == nil {
		agent := agent.NewClient(sock)
		signers, err := agent.Signers()

		if err == nil {
			clientConfig.Auth = append(clientConfig.Auth, ssh.PublicKeys(signers...))
		}
	}

	if config.Port == 0 {
		config.Port = 22
	}

	hostAndPort := config.Host + ":" + strconv.Itoa(config.Port)
	sshClient, err := ssh.Dial("tcp", hostAndPort, clientConfig)

	if err != nil {
		return client, errors.New("Failed to dial: " + err.Error())
	}

	return &Client{sshClient}, nil
}

func (client *Client) Run(command string) (stdoutStr string, stderrStr string, err error) {
	session, err := client.SshClient.NewSession()
	if err != nil {
		return
	}
	defer session.Close()

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	session.Stdout = &stdout
	session.Stderr = &stderr

	err = session.Run(command)
	stdoutStr = stdout.String()
	stderrStr = stderr.String()

	return
}

func (client *Client) Close() {
	client.SshClient.Close()
}

func makePrivateKeys(keyPaths []string) (keys []ssh.Signer, err error) {
	for _, keyPath := range keyPaths {

		if key, err := makePrivateKey(keyPath); err == nil {
			keys = append(keys, key)
		} else {
			return keys, err
		}
	}

	return
}

func makePrivateKey(keyPath string) (key ssh.Signer, err error) {
	buffer, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	key, err = ssh.ParsePrivateKey(buffer)
	if err != nil {
		return
	}

	return
}

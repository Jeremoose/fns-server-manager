package fnsservermanager

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Server struct {
	name      string
	command   string
	arguments []string
	stdOut    *bytes.Buffer
	stdErr    *bytes.Buffer
}

func NewServer(name string, command string, arguments ...string) (*Server, error) {
	if name == "" || command == "" {
		return nil, errors.New("server name and/or command are not filled")
	}
	return &Server{name: name, command: command, arguments: arguments, stdOut: new(bytes.Buffer), stdErr: new(bytes.Buffer)}, nil
}

func (s *Server) Run() error {
	if s.command == "" {
		log.Printf("SERVER %s: Command string is empty\n", s.name)
		return errors.New("command string is empty")
	}

	cmd := exec.Command(s.command, s.arguments...)
	cmd.Stdout = io.MultiWriter(os.Stdout, s.stdOut)
	cmd.Stderr = io.MultiWriter(os.Stderr, s.stdErr)

	err := cmd.Run()
	if err != nil {
		log.Printf("SERVER %s: Run() failed with %s\n", s.name, err)
		return err
	}

	return nil
}

func (s *Server) GetStdOut() (*bytes.Buffer, error) {
	if s.stdOut == nil {
		return nil, fmt.Errorf("stdOut for server %s has not been set yet", s.name)
	}
	return s.stdOut, nil
}

func (s *Server) GetStdErr() (*bytes.Buffer, error) {
	if s.stdOut == nil {
		return nil, fmt.Errorf("stdErr for server %s has not been set yet", s.name)
	}
	return s.stdErr, nil
}

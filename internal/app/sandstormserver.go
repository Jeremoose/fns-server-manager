package app

import (
	"errors"
	"log"
)

type SandstormServer struct {
	server      *Server
	logFilePath *string
}

func NewSandstormServer(name string, command string, arguments ...string) (*SandstormServer, error) {
	server, err := NewServer(name, command, arguments...)
	if err != nil {
		return nil, err
	}
	return &SandstormServer{
		server: server,
	}, nil
}

func (s *SandstormServer) Start() error {
	if s.server == nil {
		return errors.New("no server available")
	}
	if err := s.server.Run(); err != nil {
		log.Printf("couldn't start server %s\n", s.server.name)
		return err
	}
	return nil
}

func (s *SandstormServer) SetLogFilePath(logFilePath string) {
	s.logFilePath = &logFilePath
}

package server

import (
	"github.com/gpestana/sealpir/client"
)

type Server struct {
	preProcessed bool
	encryptionParams []string
	pirParams []string
}

type Reply struct {}

func New() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Preprocess() {
	if !s.preProcessed {
		//preprocess database
		s.preProcessed = true
	}
}

func (s *Server) GenerateReply(q client.Query) *Reply {
	return &Reply{}
}

func (s *Server) DecomposeToPlaintexts() *Reply {
	return &Reply{}
}

func expandQuery() {}

// others



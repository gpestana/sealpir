package sealpir

type Server struct {
	preProcessed     bool
	encryptionParams []string
	pirParams        []string
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (s *Server) Preprocess() {
	if !s.preProcessed {
		//preprocess database
		s.preProcessed = true
	}
}

func (s *Server) GenerateReply(q PirQuery) *PirReply {
	return &PirReply{}
}

func (s *Server) DecomposeToPlaintexts() *PirReply {
	return &PirReply{}
}

func expandQuery() {}

// others

package server

func (s *server) Routes() {
	s.mux.HandleFunc("/health", handleHealth())
}
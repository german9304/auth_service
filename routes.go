package server

func (s *server) Routes() {
	s.mux.HandleFunc("/health", handleHealth())
	s.mux.HandleFunc("/authorize", handleAuthenticate())
	s.mux.HandleFunc("/.well-known/openid-configuration", handleOpenId())
}

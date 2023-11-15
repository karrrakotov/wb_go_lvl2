package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,          // 1 MB
		ReadTimeout:    30 * time.Second, // 30 SEC
		WriteTimeout:   30 * time.Second, // 30 SEC
	}

	return s.httpServer.ListenAndServe()
}

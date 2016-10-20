package siad

import (
	"net"
	"net/http"
	"strings"
)

type (
	// Server creates and serves a HTTP server that offers communication with a
	// Sia API.
	Server struct {
		httpServer *http.Server
		mux        *http.ServeMux
		listener   net.Listener
	}
)

// NewServer creates a new net.http server listening on bindAddr.  Only the
// /daemon/ routes are registered by this func, additional routes can be
// registered later by calling serv.mux.Handle.
func NewServer(bindAddr string) (*Server, error) {
	// Create the listener for the server
	l, err := net.Listen("tcp", bindAddr)
	if err != nil {
		return nil, err
	}

	// Create the Server
	mux := http.NewServeMux()
	srv := &Server{
		mux:      mux,
		listener: l,
		httpServer: &http.Server{
			Handler: mux,
		},
	}

	return srv, nil
}

// Handle registers the handler for the given pattern. If a handler already exists for pattern, Handle panics.
func (srv *Server) Handle(pattern string, handler http.Handler) {
	srv.mux.Handle(pattern, handler)
}

//Serve accepts incoming connections
func (srv *Server) Serve() error {
	// The server will run until an error is encountered or the listener is
	// closed, via either the Close method or the signal handling above.
	// Closing the listener will result in the benign error handled below.
	err := srv.httpServer.Serve(srv.listener)
	if err != nil && !strings.HasSuffix(err.Error(), "use of closed network connection") {
		return err
	}
	return nil
}

// Close closes the Server's listener, causing the HTTP server to shut down.
func (srv *Server) Close() error {
	// Close the listener, which will cause Server.Serve() to return.
	if err := srv.listener.Close(); err != nil {
		return err
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	defaultPort = 9999
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	s, err := newServer()
	if err != nil {
		return err
	}

	port := fmt.Sprintf(":%d", s.port)
	err = http.ListenAndServe(port, s)
	if err != nil {
		return err
	}
	return nil
}

type server struct {
	port int
	l    *log.Logger
	r    *router
}

func newServer() (*server, error) {
	s := &server{
		port: defaultPort,
		l:    log.New(os.Stdout, "blackpage: ", log.Lshortfile),
	}
	s.routes()
	return s, nil
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.r.ServeHTTP(w, r)
}

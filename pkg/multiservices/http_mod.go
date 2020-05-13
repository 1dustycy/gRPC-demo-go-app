package multiservices

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

// HTTPMod ...
type HTTPMod struct {
	Port   string
	Server *http.Server
}

// Start ...
func (s *HTTPMod) Start() error {
	if len(s.Port) == 0 {
		return errors.New("undefined port")
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		return err
	}

	return s.Server.Serve(l)
}

// Close ...
func (s *HTTPMod) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return s.Server.Shutdown(ctx)
}

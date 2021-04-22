// Package echo contains protoc-generated output and implements a test
// and demo echo services.
// It is intended as for transflect testing only.
package echo

import (
	"context"
	"fmt"
)

// Server implements the server-side of gRPC demo Phone service.
type Server struct {
	UnimplementedEchoServer
}

// Hello is a demo echo service.
func (*Server) Hello(_ context.Context, req *HelloRequest) (*HelloResponse, error) {
	resp := &HelloResponse{Response: fmt.Sprintf("And to you: %s", req.Message)}
	return resp, nil
}

// HelloStream streaming RPC handler.
func (s *Server) HelloStream(req *HelloRequest, stream Echo_HelloStreamServer) error {
	for i := 0; i < 10; i++ {
		err := stream.Send(&HelloResponse{Response: req.Message})
		if err != nil {
			return err
		}
	}
	return nil
}

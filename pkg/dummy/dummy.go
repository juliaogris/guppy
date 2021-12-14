// Package dummy contains protoc-generated output and implements a test
// and demo dummy services.
// It is intended as for transflect testing only.
package dummy

import (
	"context"
	"fmt"
)

// Server implements the server-side of gRPC demo Phone service.
type Server struct {
	UnimplementedDummyServer
}

// Hello is a demo dummy service.
func (*Server) Say(ctx context.Context, req *SayRequest) (*SayResponse, error) {
	return &SayResponse{DoubleWord: fmt.Sprintf("%s, %s", req.Word, req.Word)}, nil
}

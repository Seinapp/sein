package utils

import (
	context "context"
)

// Ping represents a simple ping endpoint, that returns an empty value
func (s *Server) Ping(ctx context.Context, in *PingRequest) (*PingResponse, error) {
	return &PingResponse{}, nil
}

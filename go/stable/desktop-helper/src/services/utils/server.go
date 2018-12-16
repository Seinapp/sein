package utils

import (
	"go.uber.org/zap"
)

// we make sure UtilsServer implements MyStruct
var _ UtilsServer = (*Server)(nil)

// NewServer returns a new util
func NewServer(logger *zap.Logger) *Server {
	return &Server{}
}

// Server represents an Auth gRPC server
type Server struct {
	logger *zap.Logger
}

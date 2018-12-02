package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// UnaryLoggerInterceptor is a Unary interceptor that logs every incoming
// requests
func (s *Server) UnaryLoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	grpc_zap.ReplaceGrpcLogger(logger)
	return grpc_zap.UnaryServerInterceptor(logger)
}

// UnaryPanicInterceptor is a Unary interceptor that catches panics to prevent
// the server from crashing
func (s *Server) UnaryPanicInterceptor() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor()
}

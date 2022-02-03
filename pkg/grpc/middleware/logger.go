package middleware

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"grpc-streams/pkg/logger"
)

func codeToLevel(code codes.Code) zapcore.Level {
	if code == codes.OK {
		// It is DEBUG
		return zap.DebugLevel
	}
	return grpc_zap.DefaultCodeToLevel(code)
}

var (
	opts          []grpc_zap.Option
	CtxTagsUnary  grpc.UnaryServerInterceptor
	CtxTagsStream grpc.StreamServerInterceptor
	LogUnary      grpc.UnaryServerInterceptor
	LogStream     grpc.StreamServerInterceptor
)

func init() {
	fmt.Println("middleware/logger init")
	grpc_zap.ReplaceGrpcLoggerV2(logger.Log)
	opts = []grpc_zap.Option{
		grpc_zap.WithLevels(codeToLevel),
	}

	CtxTagsUnary = grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor))
	CtxTagsStream = grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor))

	LogUnary = grpc_zap.UnaryServerInterceptor(logger.Log, opts...)
	LogStream = grpc_zap.StreamServerInterceptor(logger.Log, opts...)
}

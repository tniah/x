package interceptors

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/tniah/x/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	opts = []logging.Option{
		logging.WithLogOnEvents(logging.FinishCall),
	}
)

func UnaryAuditServiceRequest(zl *logger.ZapLogger) grpc.UnaryServerInterceptor {
	return logging.UnaryServerInterceptor(requestLogger(zl), opts...)
}

func StreamAuditServiceRequest(zl *logger.ZapLogger) grpc.StreamServerInterceptor {
	return logging.StreamServerInterceptor(requestLogger(zl), opts...)
}

func requestLogger(zl *logger.ZapLogger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)
		for i := 0; i < len(fields); i += 2 {
			key := fields[i]
			val := fields[i+1]

			switch v := val.(type) {
			case string:
				f = append(f, zap.String(key.(string), v))
			case int:
				f = append(f, zap.Int(key.(string), v))
			case bool:
				f = append(f, zap.Bool(key.(string), v))
			default:
				f = append(f, zap.Any(key.(string), v))
			}
		}

		zl.Info(msg, f...)
	})
}

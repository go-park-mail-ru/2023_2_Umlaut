package interceptors

import (
	"context"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
)

func PanicRecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			log.Println(
				"Panic, ",
				"Method: ", info.FullMethod,
				//"Error: ", errPanic.(string),
				"Message: ", string(debug.Stack()),
			)
		}
	}()

	return handler(ctx, req)
}

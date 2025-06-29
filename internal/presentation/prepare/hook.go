package prepare

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func QueryServiceLifecycle(lifecycle fx.Lifecycle, server *QueryServer) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				port := 8083
				listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
				if err != nil {
					return err
				}
				reflection.Register(server.Server)
				go func() {
					log.Printf("Query service is listening on port %d", port)
					server.Server.Serve(listener)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				server.Server.GracefulStop()
				log.Println("Query service has stopped gracefully")
				return nil
			},
		},
	)
}

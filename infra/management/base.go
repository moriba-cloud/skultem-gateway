package management

import (
	"context"
	"fmt"
	"github.com/moriba-cloud/skultem-gateway/domain/auth"
	"github.com/moriba-cloud/skultem-gateway/domain/values"
	"github.com/moriba-cloud/skultem-gateway/domain/year"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"strings"
)

type (
	Services struct {
		Year  year.Service
		Value values.Service
	}
	Args struct {
		Logger *zap.Logger
	}
)

func CleanError(err string) error {
	clean := strings.ReplaceAll(err, "\n", " ")
	clean = strings.ReplaceAll(clean, "rpc error: code = Unknown desc = ", " ")
	return fmt.Errorf(clean)
}

func GrpcMetadata(ctx context.Context) context.Context {
	token := auth.ActiveAccessToken(ctx)
	md := metadata.Pairs("authorization", token)
	return metadata.NewOutgoingContext(ctx, md)
}

func NeeService(args Args) Services {
	return Services{
		Year:  NewYear(args.Logger),
		Value: NewValue(args.Logger),
	}
}

package useCase

import (
	"context"
	"github.com/xdimedrolx/moly/internal/app/service"
)

// usecase handlers
type Handlers interface {
	WebhookLogic
}

type WebhookLogic interface {
	Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error)
}

func New(services service.Services) Handlers {
	return interactor{services}
}

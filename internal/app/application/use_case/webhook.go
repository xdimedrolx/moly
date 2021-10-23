package useCase

import "context"

type EchoRequest struct {
	Message string `valid:"required" form:"message" query:"message"`
}

type EchoResponse struct {
	Message string
}

func (i interactor) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	if err := i.Validate(req); err != nil {
		return nil, err
	}

	i.services.Logger().InfoContext(ctx, req.Message)
	return &EchoResponse{Message: req.Message}, nil
}

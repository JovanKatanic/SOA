package handl

import (
	"context"
	"fmt"

	"stakeholders_service/proto/auth"
)

type AuthHandler struct {
	auth.UnimplementedStakeholderServiceServer
}

func (h AuthHandler) Greet(ctx context.Context, request *auth.Request) (*auth.Response, error) {
	return &auth.Response{
		Greeting: fmt.Sprintf("Hi %s!", request.Name),
	}, nil
}

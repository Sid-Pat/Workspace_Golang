package userservice

import (
	"context"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

// GetUser implements apiconnect.UserServiceHandler.
func (*UserService) GetUser(context.Context, *connect.Request[api.GetUserRequest]) (*connect.Response[api.User], error) {
	panic("unimplemented")
}

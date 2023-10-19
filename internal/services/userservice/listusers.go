package userservice

import (
	"context"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

// ListUsers implements apiconnect.UserServiceHandler.
func (*UserService) ListUsers(context.Context, *connect.Request[api.ListUserRequest]) (*connect.Response[api.ListUserResponse], error) {
	panic("unimplemented")
}

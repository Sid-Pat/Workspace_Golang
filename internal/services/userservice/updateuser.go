package userservice

import (
	"context"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

// UpdateUser implements apiconnect.UserServiceHandler.
func (*UserService) UpdateUser(context.Context, *connect.Request[api.UpdateUserRequest]) (*connect.Response[api.User], error) {
	panic("unimplemented")
}

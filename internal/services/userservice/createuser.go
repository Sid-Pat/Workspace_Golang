package userservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

var id int = 1

// CreateUser implements apiconnect.UserServiceHandler.
func (s *UserService) CreateUser(
	ctx context.Context,
	req *connect.Request[api.CreateUserRequest],
) (*connect.Response[api.User], error) {

	user := req.Msg.User
	if user.FirstName == "" || user.LastName == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("first name is missing"))
	}
	id += 1
	user.Id = strconv.Itoa(id)
	return connect.NewResponse(user), nil
}

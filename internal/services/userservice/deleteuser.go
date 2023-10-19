package userservice

import (
	"context"
	"errors"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
)

type UserStore struct {
	users     map[int]*api.User
	idCounter int
}

var userStore = UserStore{
	users:     make(map[int]*api.User),
	idCounter: 1,
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	req *connect.Request[api.DeleteUserRequest],
) (*connect.Response[api.User], error) {

	userIDStr := req.Msg.Id

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid user ID"))
	}

	user, exists := userStore.users[userID]
	if !exists {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("user not found"))
	}

	delete(userStore.users, userID)

	response := connect.NewResponse(user)

	return response, nil
}

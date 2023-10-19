package twitterapp_test

import (
	"context"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
	"github.com/codesmith-dev/twitter/internal/twitterapp"
)

func TestTwitterApp(t *testing.T) {
	twitterapp.Run()
	client := apiconnect.NewUserServiceClient(http.DefaultClient, "http://127.0.0.1:8080")
	create, err := client.CreateUser(context.Background(), connect.NewRequest(&api.CreateUserRequest{
		User: &api.User{
			FirstName: "",
			LastName:  "",
		},
	}))
	ok(t, err)
	notEmpty(t, create.Msg.Id)
	get, err := client.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{
		Id: create.Msg.Id,
	}))
	ok(t, err)
	equalStr(t, get.Msg.FirstName, create.Msg.FirstName)
	equalStr(t, get.Msg.LastName, create.Msg.LastName)
}

func ok(t *testing.T, err error) {
	if err != nil {
		t.FailNow()
	}
}

func notEmpty(t *testing.T, v string) {
	if v == "" {
		t.FailNow()
	}
}

func equalStr(t *testing.T, f, s string) {
	if f != s {
		t.FailNow()
	}
}

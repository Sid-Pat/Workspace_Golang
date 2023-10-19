package services

import (
	"database/sql"

	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
	"github.com/codesmith-dev/twitter/internal/services/userservice"
)

func NewUserServiceHandler(db *sql.DB) apiconnect.UserServiceHandler {
	return userservice.New(db)
}

package twitterapp

import (
	"log"
	"net/http"

	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
	"github.com/codesmith-dev/twitter/internal/services"
)

// Run runs the twitter app.
func Run() {
	http.Handle(apiconnect.NewUserServiceHandler(
		services.NewUserServiceHandler(nil),
	))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

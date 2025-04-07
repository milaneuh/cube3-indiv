package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" :example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1"))

	// TODO Start operation

	// Register GET /greeting/{name} handler.
	huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	},
	) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	http.ListenAndServe("127.0.0.1:8888", router)
}

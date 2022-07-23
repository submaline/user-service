package main

import (
	"fmt"
	userService "github.com/submaline/user-service"
	"github.com/submaline/user-service/gen/user/v1/userv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	userServiceHandler := &userService.UserService{}

	mux := http.NewServeMux()
	mux.Handle(userv1connect.NewUserServiceHandler(
		userServiceHandler))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Service listening on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

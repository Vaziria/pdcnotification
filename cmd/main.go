package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/vaziria/pdcnotification"
	"github.com/vaziria/pdcnotification/repo"
)

func main() {
	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", pdcnotification.Notification); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

func test_find() {
	user, err := repo.FindByEmail("user@gmail.com")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(user.Email)
}

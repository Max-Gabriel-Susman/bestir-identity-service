package main

import (
	"context"
	"fmt"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/handler"
)

func main() {
	ctx := context.Background()
	// concurrency shit goes hither
	run(ctx)
}

func run(ctx context.Context) {
	// cfg and setup shit right hurr

	// Start API Service
	err := handler.Handler()
	if err != nil {
		// log it or some shit
		fmt.Println("ya borked it jackass")
	}
}

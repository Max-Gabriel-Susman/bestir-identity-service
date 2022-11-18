package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/handler"
	"github.com/pkg/errors"
)

var (
	ssmParamas []byte
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

// I need to come back an grok all of this shit later
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context, _ []string) error {
	// cfg and setup shit right hurr

	// Start API Service
	h := handler.API()

	api := http.Server{
		Handler: h,
		// Addr:              "127.0.0.1:80",
		Addr:              "0.0.0.0:80",
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start listening for requests
	go func() {
		// log info about this
		serverErrors <- api.ListenAndServe()
	}()
	// Shutdown

	// logic for handling shutdown gracefully
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case <-ctx.Done():
		// log something

		// request a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}

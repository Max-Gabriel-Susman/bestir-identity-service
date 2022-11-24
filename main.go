package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/db"
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/handler"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ssmParams []byte
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func spain() {
	db.PGXSeed()
}

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
	// aws shit, currently unsupported, but soon
	//wsCfg, err := aws.NewConfig(ctx)
	//f err != nil {
	//	return errors.Wrap(err, "could not create aws sdk config")
	//

	//f _, ok := os.LookupEnv("SSM_DISABLE"); !ok {
	//	if err := awsParseSSMParams(ctx, awsCfg, ssmParams); err != nil {
	//		return err 
	//	}
	//

	// open api shit for documentation 

	// cfg and setup shit right hurr
	var cfg struct {
		ServiceName string `env:"SERVICE_NAME" envDefault:"bp-billing-service"`
		Env string `env:"ENV" envDefault:"local"`
		Database struct {
			User string `env:"BILLING_DB_USER,required"`
			Pass string `env:"BILLING_DB_PASSWORD,required"`
			Host string `env:"BILLING_DB_HOST"`
			Port string `env:"BILLING_DB_PORT envDefault:"3306"`
			DBName string `env:"BILLING_DB_NAME" envDefault:"billing"`
			Params string `env:"BILLING_DB_PARAM_OVERRIDES envDefault:"parseTime=true"`
		}
		Datadog struct {
			Disable bool `env:"DD_DISABLE"`
		}
	}

	db, err := database.Open(database.Config{
		User: cfg.Database.User,
		Password: cfg.Database.Pass,
		Host: cfg.Database.Host, 
		Name: cfg.Database.DBName,
		Params: cfg.Database.Params,
	}, cfg.ServiceName)
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer func() {
		// log.info(ctx, "stopping database")
		db.Close()
	}


	// Read in connection string
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// we gott reconfigure the service to use pgx now
	h := handler.API(handler.Deps{Conn: conn})

	// Start API Service
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

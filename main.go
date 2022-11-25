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
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/database"
	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/handler"
	env "github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
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

	// cfg and setup shit right hurr, we gotta alter it for my database setup
	var cfg struct {
		ServiceName string `env:"SERVICE_NAME" envDefault:"bp-billing-service"`
		Env         string `env:"ENV" envDefault:"local"`
		Database    struct {
			User   string `env:"Bestir_Platform_User,required"`
			Pass   string `env:"Bestir_Platform_Account,required"`
			Host   string `env:"Bestir_Identity_DB_Host"`
			Port   string `env:"Bestir_DB_Port" envDefault:"3306"`
			DBName string `env:"Bestir_DB_Name" envDefault:"identity"`
			Params string `env:"Bestir_DB_Param_Overrides" envDefault:"parseTime=true"`
		}
		Datadog struct {
			Disable bool `env:"DD_DISABLE"`
		}
	}
	if err := env.Parse(&cfg); err != nil {
		return errors.Wrap(err, "parsing configuration")
	}
	// cfg.Datadog.Disable = true

	// z = z.With(
	// 	zap.
	// )

	db, err := database.Open(database.Config{
		User:     cfg.Database.User,
		Password: cfg.Database.Pass,
		Host:     cfg.Database.Host,
		Name:     cfg.Database.DBName,
		Params:   cfg.Database.Params,
	}, cfg.ServiceName)
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer func() {
		// zl.Info(ctx, "stopping database")
		db.Close()
	}()

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

	// If DD is enabled, configure db to send stats info
	if !cfg.Datadog.Disable {
		// statsdAddress := ddAgentAddress(8125)
		// statsd, err := statsd.New(statsdAddress, statsd.WithMaxBytesPerPayload(4096))
		// if err != nil {
		// 	return errors.Wrap(err, "could not start statsd client")
		// }
		// // Start Stats Reporting for db
		// go func() {
		// 	zl.Info(ctx, "Starting reporting DB metrics", zap.String("statsd.address", statsdAddress))
		// 	defer zl.Info(ctx, "Stopped reporting DB metrics")
		// 	sr := database.NewStatsReporter(db, statsd, zl)

		// 	sr.ReportDBStats(ctx, []string{
		// 		fmt.Sprintf("service:%s", cfg.ServiceName),
		// 		fmt.Sprintf("version:%s", GitSHA),
		// 		fmt.Sprintf("env:%s", cfg.Env),
		// 		"collabs.squad:red",
		// 	}, 1)
		// }()
	}

	// CLIENTS N SHIT

	// event bridge shit

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

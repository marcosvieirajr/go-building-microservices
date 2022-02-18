package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
	"github.com/marcosvieirajr/go-multi-tier-microservices/currency/proto"
	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/data"
	"github.com/marcosvieirajr/go-multi-tier-microservices/product-api/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	loadDevEnvs(log.Default())

	var (
		haEnv       = os.Getenv("HTTP_ADDRESS")
		acEnv       = os.Getenv("ALLOWED_CORS")
		logLevelEnv = os.Getenv("LOG_LEVEL")

		httpAddr    = flag.String("http", haEnv, "HTTP service address")
		allowedCORS = flag.String("cors", acEnv, "Cross-Origin Resource Sharing")
		logLevel    = flag.String("LOG_LEVEL", logLevelEnv, "Log output level for the server [debug, info, trace]")
	)
	flag.Parse()

	dl := !strings.EqualFold("info", *logLevel)

	// l := log.New(os.Stdout, "product-api ", log.LstdFlags|log.Lshortfile)
	l := hclog.New(&hclog.LoggerOptions{
		Name:            "product-api",
		Level:           hclog.LevelFromString(*logLevel),
		Output:          os.Stdout,
		JSONFormat:      false,
		IncludeLocation: dl,
		TimeFormat:      "2006-01-02 15:04:05.000",
	})
	sl := l.StandardLogger(&hclog.StandardLoggerOptions{
		InferLevels: true,
	})

	v := data.NewValidation()

	// create client
	var opts []grpc.DialOption = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial("currency:9092", opts...)
	if err != nil {
		l.Error("error connecting to grpc server", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	cc := proto.NewCurrencyClient(conn)

	// create the handlers
	ph := handlers.NewProducts(l, v, cc)

	// create a new serve mux and register the handlers
	r := mux.NewRouter()

	listAllR := r.Methods(http.MethodGet).Subrouter()
	listAllR.HandleFunc("/products", ph.HandleListAll())

	listSingleR := r.Methods(http.MethodGet).Subrouter()
	listSingleR.HandleFunc("/products/{id:[0-9]+}", ph.HandleListSingle())

	createR := r.Methods(http.MethodPost).Subrouter()
	createR.HandleFunc("/products", ph.HandleCreate())
	createR.Use(ph.MiddlewareValidateProduct)

	updateR := r.Methods(http.MethodPut).Subrouter()
	updateR.HandleFunc("/products/{id:[0-9]+}", ph.HandleUpdate())
	updateR.Use(ph.MiddlewareValidateProduct)

	deleteR := r.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	pingR := r.Methods(http.MethodGet).Subrouter()
	pingR.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(struct{ Status string }{Status: "OK"})
	})

	ch := gohandlers.CORS(gohandlers.AllowedOriginValidator(originValidator(*allowedCORS)))

	// create a new server
	srv := http.Server{
		Addr:         *httpAddr,         // configure the bind address
		Handler:      ch(r),             // set the default handler
		ErrorLog:     sl,                // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Info("starting server", "bind_address", *httpAddr)

		err := srv.ListenAndServe()
		if err != nil {
			l.Error("error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	l.Info("shutting down server by signal", "signal", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
	return nil
}

func loadDevEnvs(l *log.Logger) {
	if err := godotenv.Load(".env"); err != nil {
		l.Println("no .env file found to loading")
	}
}

func originValidator(allowed string) func(string) bool {
	return func(origin string) bool {
		a := strings.Split(allowed, ",")
		for _, v := range a {
			if strings.HasSuffix(origin, v) {
				return true
			}
		}
		return false
	}
}

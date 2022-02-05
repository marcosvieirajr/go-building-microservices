package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/marcosvieirajr/go-multi-tier-microservices/data"
	"github.com/marcosvieirajr/go-multi-tier-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags|log.Lshortfile)

	loadDevEnvs(l)

	var (
		haEnv = os.Getenv("HTTP_ADDRESS")
		acEnv = os.Getenv("ALLOWED_CORS")

		httpAddr    = flag.String("http", haEnv, "HTTP service address")
		allowedCORS = flag.String("cors", acEnv, "Cross-Origin Resource Sharing")
	)
	flag.Parse()

	v := data.NewValidation()

	// create the handlers
	ph := handlers.NewProducts(l, v)

	// create a new serve mux and register the handlers
	r := mux.NewRouter()

	listAllR := r.Methods(http.MethodGet).Subrouter()
	listAllR.HandleFunc("/products", ph.ListAll)

	listSingleR := r.Methods(http.MethodGet).Subrouter()
	listSingleR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	createR := r.Methods(http.MethodPost).Subrouter()
	createR.HandleFunc("/products", ph.Create)
	createR.Use(ph.MiddlewareValidateProduct)

	updateR := r.Methods(http.MethodPut).Subrouter()
	updateR.HandleFunc("/products/{id:[0-9]+}", ph.Update)
	updateR.Use(ph.MiddlewareValidateProduct)

	deleteR := r.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	ch := gohandlers.CORS(gohandlers.AllowedOriginValidator(originValidator(*allowedCORS)))

	// create a new server
	srv := http.Server{
		Addr:         *httpAddr,         // configure the bind address
		Handler:      ch(r),             // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Printf("starting server on port %v", *httpAddr)

		err := srv.ListenAndServe()
		if err != nil {
			l.Printf("error starting server: %s\n", err)
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
	l.Printf("got signal: %v. trying graceful shutdown", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
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

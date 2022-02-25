package internal

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	logger "go.lafronz.com/tools/logger/stackdriver"
)

// serverDetails struct
type serverDetails struct {
	port string
}

var (
	s       serverDetails
	version string = os.Getenv("version")
)

func init() {
	// use PORT environment variable, or default to 8080
	s.port = "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		s.port = fromEnv
	} else {
		logger.Warning("No Port Provided, using default: %s", s.port)
	}
}

// StartServer func
func StartServer() {
	var wait time.Duration

	r := mux.NewRouter()

	setUpRoutes(r)

	// Server Settings
	svr := &http.Server{
		Handler:      r,
		Addr:         ":" + s.port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		// start the web server on port and accept requests
		logger.Info("Server listening on port %s", s.port)
		if err := svr.ListenAndServe(); err != nil {
			logger.Error("%s", err)
		}
	}()

	gracefulStop := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT, SIGTERM, or Interrupt
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	// Block until we receive our signal.
	<-gracefulStop

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	svr.Shutdown(ctx)

	logger.Info("Web server shut down")

	close(gracefulStop)
}

func setUpRoutes(r *mux.Router) {
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/version", versionHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/federation", federationHandler)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(version))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://developers.stellar.org/docs/glossary/federation/", http.StatusTemporaryRedirect)
}

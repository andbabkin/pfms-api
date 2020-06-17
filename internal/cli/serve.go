package cli

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	httpi "github.com/andbabkin/pfms-api/internal/http"
)

// Serve is a command which runs a server
type Serve struct{}

// Execute command
func (s *Serve) Execute(args []string) error {
	port := DefaultPort
	if len(args) > 0 {
		_, err := strconv.ParseUint(args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Wrong argument. The port should be an integer value.\nExample: pfms serve %s", DefaultPort)
		}
		port = args[0]
	}

	fmt.Printf("The server listens to :%s...\n", port)

	startHTTPSServer(`:` + port)

	return nil
}

func startHTTPSServer(addr string) {
	mux := http.NewServeMux()
	httpi.Handle(mux)
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Run server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServeTLS("public.crt", "private.key"); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	fmt.Println("")
	log.Println("Shutting down...")
	srv.Shutdown(ctx)
	// Optionally, we could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if the application should wait for other services
	// to finalize based on context cancellation.
	os.Exit(0)
}

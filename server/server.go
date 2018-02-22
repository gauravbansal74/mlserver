package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gauravbansal74/mlserver/pkg/logger"
)

// Config - server configuration
type Config struct {
	Hostname  string `json:"Hostname"`  // Server name
	UseHTTP   bool   `json:"UseHTTP"`   // Listen on HTTP
	UseHTTPS  bool   `json:"UseHTTPS"`  // Listen on HTTPS
	HTTPPort  int    `json:"HTTPPort"`  // HTTP port
	HTTPSPort int    `json:"HTTPSPort"` // HTTPS port
	CertFile  string `json:"CertFile"`  // HTTPS certificate
	KeyFile   string `json:"KeyFile"`   // HTTPS private key
}

// Run starts the HTTP and/or HTTPS listener
func Run(httpHandlers http.Handler, httpsHandlers http.Handler, s Config) {
	if s.UseHTTP && s.UseHTTPS {
		go func() {
			startHTTPS(httpsHandlers, s)
		}()
		startHTTP(httpHandlers, s)
	} else if s.UseHTTP {
		startHTTP(httpHandlers, s)
	} else if s.UseHTTPS {
		startHTTPS(httpsHandlers, s)
	} else {
		log.Fatal("Config file doesn't specify a lietener to start")
	}
}

// startHTTP starts the HTTP listener
func startHTTP(handlers http.Handler, s Config) {
	logger.Info("Server started at HTTP", logger.Fields{"HOST": httpAddress(s)})
	// Start the HTTP listener
	log.Fatal(http.ListenAndServe(httpAddress(s), handlers).Error(), "error while starting HTTP server")
}

// startHTTPs starts the HTTPS listener
func startHTTPS(handlers http.Handler, s Config) {
	logger.Info("Server started at HTTPS", logger.Fields{"HOST": httpAddress(s)})
	// Start the HTTPS listener
	log.Fatal(http.ListenAndServeTLS(httpsAddress(s), s.CertFile, s.KeyFile, handlers).Error(), "error while starting HTTPS server")
}

// httpAddress returns the HTTP address
func httpAddress(s Config) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPPort)
}

// httpsAddress returns the HTTPS address
func httpsAddress(s Config) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPSPort)
}

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var serverCount = 0

const (
	SERVER1 = "https://www.google.com"
	SERVER2 = "https://www.facebook.com"
	SERVER3 = "https://www.yahoo.com"
	PORT    = "1337"
)

// Serves a reverse proxy
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	proxy.ServeHTTP(res, req)
}

func logRequestPayload(proxyURL string) {
	log.Printf("Proxy url: %s\n", proxyURL)
}

// Returns one of the servers
func getProxyURL() string {
	var servers = []string{SERVER1, SERVER2, SERVER3}

	server := servers[serverCount]
	serverCount++

	if serverCount >= len(servers) {
		serverCount = 0
	}

	return server
}

// Handles request and redirects
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	url := getProxyURL()

	logRequestPayload(url)

	serveReverseProxy(url, res, req)
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)

	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

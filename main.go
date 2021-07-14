package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// Serve a reverse proxy for a given url
func handle(res http.ResponseWriter, req *http.Request) {
	log.Printf("Handling request to %s %s %s from %s with user agent %s", req.Method, req.Host, req.RequestURI, req.RemoteAddr, req.UserAgent())

	// parse the url
	parsedUrl, urlParseErr := url.Parse(os.Getenv("PROXY_TARGET"))
	if urlParseErr != nil {
		panic(urlParseErr)
	}

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(parsedUrl)

	// Update the headers to allow for SSL redirection
	req.URL.Host = parsedUrl.Host
	req.URL.Scheme = parsedUrl.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = parsedUrl.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

func main() {
	target := os.Getenv("PROXY_TARGET")
	host, getHostErr := os.Hostname()
	port := os.Getenv("HOST_PORT")

	if getHostErr != nil {
		panic(getHostErr)
	}

	if len(target) < 1 {
		panic("Target url not configured. Please set \"PROXY_TARGET\" environmental variable.")
	}

	// Log setup values
	log.Printf("Proxying requests to http://%s:%s/ -> %s\n", host, port, target)

	http.HandleFunc("/", handle)
	if serveErr := http.ListenAndServe(":" + port, nil); serveErr != nil {
		panic(serveErr)
	}
}

package server

import (
    "io"
    "log"
    "net/http"
    "fmt"
)

// handleProxy logic intercepts the client request and forwards it
func handleProxy(w http.ResponseWriter, r *http.Request) {
    // 1. Prepare the outgoing request to the destination server
    // We use the URL provided in the client's request
    proxyReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
    if err != nil {
        http.Error(w, "Failed to create proxy request", http.StatusInternalServerError)
        return
    }

    // 2. Copy headers from the original request to the proxy request
    for name, values := range r.Header {
        for _, value := range values {
            proxyReq.Header.Add(name, value)
        }
    }

    // 3. Execute the request using a default HTTP client
    transport := http.DefaultTransport
    resp, err := transport.RoundTrip(proxyReq)
    if err != nil {
        http.Error(w, "Failed to reach destination", http.StatusBadGateway)
        return
    }
    defer resp.Body.Close()

    // 4. Copy response headers back to the client
    for name, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }

    // 5. Set the status code and stream the body
    w.WriteHeader(resp.StatusCode)
    io.Copy(w, resp.Body)
}

func Proxy() {
    mux := http.NewServeMux()
    
    // The "/" pattern catches all incoming requests to act as a proxy
    mux.HandleFunc("/", handleProxy)

    fmt.Println("Proxy Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
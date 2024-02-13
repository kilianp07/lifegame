 # Package server

This package contains the implementation of a simple HTTP and WebSocket server written in Go. It exposes two endpoints: `/` for serving static files and `/ws` for handling WebSocket connections.

## Contents

1. [Introduction](#introduction)
2. [Code Overview](#code-overview)
3. [Serve Home Function](#serve-home-function)
4. [Run Function](#run-function)
5. [Hub](#hub)
6. [Testing the Server](#testing-the-server)

## Introduction <a name="introduction"></a>

This Go code is designed to create and run a simple HTTP server with WebSocket support. It's an excellent starting point for anyone looking to understand how to implement basic HTTP routing, handling WebSocket connections, and creating concurrent programs in Go.

## Code Overview <a name="code-overview"></a>

The `server` package consists of three functions: `serveHome`, `Run`, and an initialized `hub`.

1. The `serveHome` function handles HTTP requests for the root path (`/`) using the HTTP `HandleFunc()` function. It checks whether the request is for a valid route and method before serving static HTML content as a response.
2. The `Run` function initializes an instance of the `hub`, starts its goroutine, sets up the HTTP endpoints, and starts the HTTP server.
3. The `hub` object handles WebSocket connections. It is initialized in the `Run` function and started using its `run()` method.

## Serve Home Function <a name="serve-home-function"></a>

The `serveHome` function checks if a request matches the root path (`/`) and only accepts GET requests. If the conditions are met, it serves static HTML content using the `fmt.Fprintf()` method to write the response body. Otherwise, an error is returned depending on the mismatched condition.

```go
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "%s", htmlContent)
}
```

## Run Function <a name="run-function"></a>

The `Run` function initializes the `hub`, starts its goroutine, sets up the HTTP endpoints using the `http.HandleFunc()` method, and starts the HTTP server using the `ListenAndServe()` method.

```go
func Run() {
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Chat server error : ", err)
	}
}
```

## Hub <a name="hub"></a>

The `hub` is an instance of the WebSocket server that handles incoming connections and manages user sessions. It is initialized in the `Run()` function and started using its `run()` method.

## Testing the Server <a name="testing-the-server"></a>

To test this server, create a new Go file with the `package main` declaration at the top, import the `server` package, and call its `Run()` function to start the server. To interact with the WebSocket endpoint, you can use a WebSocket testing tool like [websocat](https://github.com/vi/websocat) or write your own Go client for testing purposes.

```go
package main

import "server"

func main() {
	server.Run()
}
```


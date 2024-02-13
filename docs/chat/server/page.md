 # server Package Documentation

This package contains the code for a simple chat application using Go as the backend and JavaScript for the frontend. The application uses WebSockets for real-time communication between clients and the server.

## Constants

The following constant is defined in this package:

### htmlContent

This constant holds the HTML content of the chat application. It includes both the JavaScript code that handles WebSocket communication and the CSS styles to format the chat window.

```go
const htmlContent = `...`
```

## Creating a WebSocket Server

To create a WebSocket server, use the standard library's net/http package along with the upgrader middleware:

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

type ChatMessage struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading", err)
        return
    }
    defer ws.Close()

    for {
        msgType, msg, err := ws.ReadMessage()
        if err != nil {
            log.Println("Error reading:", err)
            break
        }

        switch msgType {
        case websocket.TextMessage:
            fmt.Println("Received text message:", string(msg))
            // Broadcast the message to all connected clients
            broadcast <- ChatMessage{Message: string(msg)}
        }
    }
}
```

This example sets up a simple HTTP server on port 8080. When a client connects, it upgrades the connection to a WebSocket and starts listening for messages from the client. Once a message is received, it's broadcasted to all connected clients using a `chan<- ChatMessage` channel.

## Creating an HTML Frontend

The provided `htmlContent` constant holds the HTML code of our simple chat application. When loaded in a web browser that supports WebSockets, it will establish a connection with the server and start sending and receiving messages in real-time.

To serve this HTML content from your Go application, you can use the net/http package:

```go
func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        w.Header().Set("Content-Type", "text/html")
        w.Write([]byte(htmlContent))
    } else {
        // Handle WebSocket connections here...
    }
}
```

This simple example sets up a handler that serves the HTML content when the route '/' is requested via GET method. When a client sends a message through the web application, the JavaScript code will handle sending it to the server over the established WebSocket connection.

## Running the Application

To run this application, create a new Go file (e.g., `main.go`), add the above code snippets and any missing dependencies, then build and run the application using the 'go' command:

```sh
go build && go run main.go
```

This will start the chat server on port 8080, and you can access it in your web browser at `http://localhost:8080`.


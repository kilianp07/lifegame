 # server Package Documentation

This package is a simple WebSocket server implementation in Go using the Gorilla WebSocket library. It provides a basic Hub and handling of individual WebSocket connections.

## Constants

The package defines some constants that are used throughout the codebase for configuring various aspects of the server:

- `writeWait`: Time allowed to write a message to the peer.
- `pongWait`: Time allowed to read the next pong message from the peer.
- `pingPeriod`: Send pings to peer with this period. Must be less than pongWait.
- `maxMessageSize`: Maximum message size allowed from peer.

## Global Variables

The package defines some global variables used throughout the codebase:

- `newline`: A byte slice containing a single newline character.
- `space`: A byte slice containing a single space character.
- `upgrader`: An instance of the Gorilla WebSocket upgrader with custom read and write buffer sizes.

## Client Type

The `Client` type is used to represent a single WebSocket connection. It holds a reference to the hub, the underlying WebSocket connection, and a buffered channel for sending messages.

### Methods

#### `readPump()`

The `readPump()` method pumps messages from the WebSocket connection to the hub. It is started in a per-connection goroutine and ensures that there is at most one reader on a connection by executing all reads from this goroutine.

#### `writePump()`

The `writePump()` method pumps messages from the hub to the WebSocket connection. A goroutine running writePump is started for each connection, and the application ensures that there is at most one writer to a connection by executing all writes from this goroutine.

## serveWs Function

The `serveWs()` function handles WebSocket requests from the peer. It upgrades the HTTP response to a WebSocket connection using the Gorilla WebSocket upgrader, creates a new client instance, and starts the writePump() and readPump() goroutines for that client.


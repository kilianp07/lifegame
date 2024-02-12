 # Server Package Documentation

This package implements a simple hub for managing a set of active clients and broadcasting messages to them. The hub is implemented as a `Hub` type, which maintains the following components:

- `clients`: A map storing pointers to all currently registered clients.
- `broadcast`: A channel used to send messages to be broadcasted to all connected clients.
- `register`: A channel through which new clients register with the hub.
- `unregister`: A channel through which clients request to unregister from the hub.

## Hub Type

```go
type Hub struct {
	clients          map[*Client]bool
	broadcast        chan []byte
	register         chan *Client
	unregister       chan *Client
}
```

## Functions

### `newHub() *Hub`

Creates and initializes a new `Hub` instance. It returns a fresh Hub with initialized channels and an empty client map.

```go
func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}
```

### `func (h *Hub) run()`

The `run` method is responsible for managing the hub's main loop. It listens to three different channels for client registration, unregistration and broadcasting messages. The logic in this function ensures that only registered clients are allowed to receive messages and that any disconnected clients are removed from the list of active clients.

```go
func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		// ...
		}
	}
}
```


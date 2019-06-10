## go-ws-broadcast

A simple implementation of a concurrent broadcast server that dispatches a message to the connected clients via the websocket protocol.

![Diagram](https://i.imgur.com/l78X7jH.png)

Components:

## [Broadcast Server](https://github.com/zianwar/go-ws-broadcast/tree/master/server)
  consists of:
- A counter that gets incremented frequently and its value will be broadcasted to all clients.
- A hub component that registers and de-registers clients and broadcasts the counter value to the clients.

## [Web client](https://github.com/zianwar/go-ws-broadcast/tree/master/webclient)
This is a basic React app that acts as a client to the broadcast server, which:
 - Establishes a websocket connection to the server and displays a green background upon a successfull connection.
 - Starts reading the counter value and displaying it.
 - Changes the background to red when the websocket connection closes or errors out.

## Demo

[![Demo](https://img.youtube.com/vi/e7rBuhho3ks/maxresdefault.jpg)](https://youtu.be/e7rBuhho3ks)

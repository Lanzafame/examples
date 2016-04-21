## Folder information

This folder contains a simple  how to use iris websocket package, but the original example is taken from  [Gorila websocket/example](https://github.com/gorilla/websocket/tree/master/examples/echo)

## Client and server example

This example shows a simple client and server.

The server echoes messages sent to it. The client sends a message every second
and prints all messages received.

To run the example, start the server:

    $ go run server/server.go

Next, start the client:

    $ go run client/client.go

The server includes a simple web client. To use the client, open
http://127.0.0.1:8080 in the browser and follow the instructions on the page.

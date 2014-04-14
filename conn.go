package main

import (
    "net"
)

/* Structure to store socket connections */
type connection struct {
    socket *net.TCPConn /* The websocket connection */
    node string         /* Node with whom connection is established */
}

/*
 * Writes a message to an outgoing socket connection.
 */
func (c *connection) writer(msg *message) (bool) {
    buffer, err := serializeMessage(msg)
    if err != nil {
        return false
    }

    _, err = c.socket.Write(buffer)
    if err != nil {
        return false
    }

    return true
}

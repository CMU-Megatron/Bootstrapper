package main

import (
    "net"
)

/* Structure to store socket connections */
type connection struct {
    socket *net.TCPConn  /* The websocket connection */
    wsocket *net.TCPConn /* Write socket connection */
    node string          /* Node with whom connection is established */
}

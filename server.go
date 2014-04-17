package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    service := ":1800"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        conn, err := listener.AcceptTCP()
        if err != nil {
            continue
        }

        go handleRequest(conn)
    }
}

func handleRequest(conn *net.TCPConn) {
    var buf [BUFF_SIZE]byte
    var msg message
    defer conn.Close()

    n, err := conn.Read(buf[0:])
    if err != nil {
        return
    }

    err = deserializeMessage(buf, n, &msg)
    if err != nil {
        return
    }

    c := &connection {
                       socket : conn,
                       node   : msg.data,
                     }

    mqMutex.Lock()
    defer mqMutex.Unlock()

    if masterQueue.Len() == 0 || msg.messageType == INIT {
        /* If the node does not already know that it is the master, notify */
        if (msg.messageType == INIT) {
            msg.messageType = INIT

            /* Declare master */
            if c.writer(&msg) == false {
                return
            }
        }

        insertMasterQueue(c)
        mqMutex.Unlock()

        /*
         * Wait for the bootstrapping process to complete. Once bootstrapping
         * is complete, master node will close the socket from it's end.
         */
        conn.Read(buf[0:])
        mqMutex.Lock()
        removeMasterQueue(c)
    } else {
        masterConn := getCurrentMaster()
        msg.messageType = BOOTSTRAP
        msg.data = masterConn.node

        /* Send master data to the node */
        c.writer(&msg)
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(-1)
    }
}

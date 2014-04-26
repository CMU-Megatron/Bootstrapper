package main

import (
    "fmt"
    "net"
    "os"
    "encoding/gob"
)

func main() {
    service := ":1800"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    gob.Register(Message{})

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
    fmt.Println("New connection..")
    defer conn.Close()

    var data interface{}
    encoder := gob.NewEncoder(conn)
    decoder := gob.NewDecoder(conn)

    err := decoder.Decode(&data)
    if err != nil {
        fmt.Println("Deserialization err: ", err)
        return
    }

    msg := data.(Message)
    c := &connection {
                       socket : conn,
                       node   : msg.Data,
                     }

    mqMutex.Lock()
    defer mqMutex.Unlock()
    fmt.Println("Queue length : ", masterQueue.Len())

    if masterQueue.Len() == 0 || msg.MessageType == INIT {
        /* If the node does not already know that it is the master, notify */
        if (msg.MessageType != INIT) {
            msg.MessageType = INIT

            /* Declare master */
            fmt.Println("Declaring " + msg.Data + "as master")
            var data2 interface{}
            data2 = &msg
            err := encoder.Encode(&data2)
            if err != nil {
                fmt.Println(err)
                return
            }
        }

        insertMasterQueue(c)
        mqMutex.Unlock()

        /*
         * Wait for the bootstrapping process to complete. Once bootstrapping
         * is complete, master node will stop sending heart-beats and close the
         * socket from it's end.
         */
        for {
            err := decoder.Decode(&data)
            if err != nil {
                msg = data.(Message)
                fmt.Println("Bootstrapping completed for " + msg.Data)
                fmt.Println("Deserialization err: ", err)
                return
            }
        }

        fmt.Println("Bootstrapping complete")
        mqMutex.Lock()
        removeMasterQueue(c)
    } else {
        masterConn := getCurrentMaster()
        msg.MessageType = BOOTSTRAP
        fmt.Println("Telling " + msg.Data + " that " + masterConn.node + " is the master")
        msg.Data = masterConn.node

        /* Send master Data to the node */
        data := &msg
        encoder.Encode(&data)
    }
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(-1)
    }
}

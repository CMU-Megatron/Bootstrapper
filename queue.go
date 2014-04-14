package main

import (
     "sync"
     "container/list"
)

var masterQueue = list.New()
var mqMutex = &sync.RWMutex{}

/*
 * Inserts the connection of a master node in a queue.
 *
 * Note: It is assumed that hubMutex lock is already held before calling this
 *       function.
 */
func insertMasterQueue(conn *connection) {
    masterQueue.PushBack(conn)
}

/*
 * Fetches the current master node that needs to be serviced.
 *
 * Note: It is assumed that hubMutex lock is already held before calling this
 *       function.
 */
func getCurrentMaster() (*connection) {
    e := masterQueue.Front()
    if e == nil {
        return nil
    }

    return e.Value.(*connection)
}

/*
 * Removes connection of a master node who has completed its bootstrapping
 * process or has crashed.
 *
 * Note: It is assumed that hubMutex lock is already held before calling this
 *       function.
 */
func removeMasterQueue(conn *connection) {
    for e := masterQueue.Front(); e != nil; e = e.Next() {
        curr := e.Value.(*connection)
        if curr == conn {
            masterQueue.Remove(e)
            return
        }
    }
}

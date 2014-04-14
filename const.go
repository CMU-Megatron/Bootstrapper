package main

type GameStatus int
type MessageType int

const (
    NODE_COUNT int	= 3
    BUFF_SIZE int	= 512
    DELIMITER string	= " "
)

/* Game status */
const (
    CONFIG GameStatus  = 0
    PLAY   GameStatus  = 1
)

/* Message Types */
const (
    INIT		MessageType = 0 /* Declare master */
    BOOTSTRAP           MessageType = 1 /* Configuration message */
    MULTICAST		MessageType = 2 /* Multicast message */
    REPLY_MULTICAST	MessageType = 3 /* Reply to a multicast message */
    CONLOST		MessageType = 4 /* Connection lost */
    CONCHECK		MessageType = 5 /* Connection check */
    CONVOTE		MessageType = 6 /* Connection vote */
    CONCLOSE		MessageType = 7 /* Connection close */
    ECHO		MessageType = 8 /* Echo message */
    REPLY_ECHO		MessageType = 9 /* Reply to an echo message */
)

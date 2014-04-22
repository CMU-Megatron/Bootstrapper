package main

type GameStatus int
type MessageType int

const (
    NODE_COUNT int	= 4
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
    HEART_BEAT		MessageType = 2 /* Heart-beat message */
    CONNECT		MessageType = 3 /* Connect message */
    MULTICAST		MessageType = 4 /* Multicast message */
    REPLY_MULTICAST	MessageType = 5 /* Reply to a multicast message */
    CONLOST		MessageType = 6 /* Connection lost */
    CONCHECK		MessageType = 7 /* Connection check */
    CONVOTE		MessageType = 8 /* Connection vote */
    CONCLOSE		MessageType = 9 /* Connection close */
)

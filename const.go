package main

type GameStatus int
type MessageType int

const (
    NODE_COUNT int	= 2
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
    BOOTSTRAP           MessageType = 1 /* Configuration Message */
    HEART_BEAT		MessageType = 2 /* Heart-beat Message */
    CONNECT		MessageType = 3 /* Connect Message */
    MULTICAST		MessageType = 4 /* Multicast Message */
    REPLY_MULTICAST	MessageType = 5 /* Reply to a multicast Message */
    CONLOST		MessageType = 6 /* Connection lost */
    CONCHECK		MessageType = 7 /* Connection check */
    CONVOTE		MessageType = 8 /* Connection vote */
    CONCLOSE		MessageType = 9 /* Connection close */
)

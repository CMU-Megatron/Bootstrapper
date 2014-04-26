package main

/* Structure of a Message */
type Message struct {
    Sender string           /* Message source */
    OriSender string        /* Original Sender */
    seqNum uint32           /* Sequence number */
    VecStamp Timestamp      /* Vector timestamp */
    MessageType MessageType /* Type of Message to be conveyed */
    Data string             /* Message payload */
}

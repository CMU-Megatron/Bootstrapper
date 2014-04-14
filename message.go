package main

import (
    "bytes"
    "encoding/gob"
)

/* Structure of a message */
type message struct {
    sender string           /* Message source */
    oriSender string        /* Original sender */
    seqNum uint32           /* Sequence number */
    vecStamp timestamp      /* Vector timestamp */
    messageType MessageType /* Type of message to be conveyed */
    data string             /* Message payload */
}

/*
 * Encoding function to convert message into raw bytes.
 */
func (m *message) GobEncode() ([]byte, error) {
    buf := new(bytes.Buffer)
    encoder := gob.NewEncoder(buf)
    err := encoder.Encode(m.sender)
    if err != nil {
        return nil, err
    }

    err = encoder.Encode(m.oriSender)
    if err != nil {
        return nil, err
    }

    err = encoder.Encode(m.seqNum)
    if err != nil {
        return nil, err
    }

    for i := 0; i < NODE_COUNT; i++ {
        err = encoder.Encode(m.vecStamp.vector[i])
        if err != nil {
            return nil, err
        }
    }

    err = encoder.Encode(m.messageType)
    if err!=nil {
        return nil, err
    }

    err = encoder.Encode(m.data)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}

/*
 * Decoding function to construct message from raw bytes.
 */
func (m *message) GobDecode(buf []byte) error {
    r := bytes.NewBuffer(buf)
    decoder := gob.NewDecoder(r)

    err := decoder.Decode(&m.sender)
    if err != nil {
        return err
    }

    err = decoder.Decode(&m.oriSender)
    if err != nil {
        return err
    }

    err = decoder.Decode(&m.seqNum)
    if err != nil {
        return err
    }

    for i := 0; i < NODE_COUNT; i++ {
        err = decoder.Decode(&m.vecStamp.vector[i])
        if err != nil {
            return err
        }
    }

    err = decoder.Decode(&m.messageType)
    if err!=nil {
        return err
    }

    return decoder.Decode(&m.data)
}

/*
 * Serializes outbound data.
 */
func serializeMessage(msg *message) ([]byte, error) {
    buffer := new(bytes.Buffer)
    enc := gob.NewEncoder(buffer)
    err := enc.Encode(msg)

    return buffer.Bytes(), err
}

/*
 * Deserializes inbound data.
 */
func deserializeMessage(buf [BUFF_SIZE]byte, n int, msg *message) (error) {
    buffer := bytes.NewBuffer(buf[0:n])
    dec := gob.NewDecoder(buffer)
    err := dec.Decode(msg)

    return err
}

package protocol

import "github.com/bharadwaja-rao-d/syncing/diff"

type messageType int;

const (
    HandShake messageType = iota
    Update
    Close
)

// Mostly b/w client and server
type CSMessage[T diff.EditScript | string] struct {
//type CSMessage struct{
    Mtype messageType
    From string
    Content T
}

// Mostly b/w Integrator and client
type IMessage struct {
}

package replica

import (
	"forTest/message"
	"forTest/node"

	"forTest/identity"
)

type Replica struct {
	node.Node
	global_sequence chan message.Sequencer_Message
	eventChan       chan interface{}
}

func NewReplica(id identity.NodeID) *Replica {
	r := new(Replica)
	r.Node = node.NewNode(id)
	r.eventChan = make(chan interface{})
	r.global_sequence = make(chan message.Sequencer_Message)

	return r
}


func (r *Replica) Start(){
	
}
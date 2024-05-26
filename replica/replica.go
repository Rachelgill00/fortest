package replica

import (
	"fmt"
	"forTest/message"
	"forTest/node"

	"forTest/identity"

	"forTest/mempool"
)

type Replica struct {
	node.Node
	pd              *mempool.Producer
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

func (r *Replica) Start() {
	txns_batch := r.pd.GeneratePayload()
	for _, tx := range txns_batch {
		fmt.Println(tx.ID)
		fmt.Println(tx.Timestamp)
	}
}

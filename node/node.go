package node

import (
	"forTest/config"
	"forTest/identity"
	"sync"
)

type Node interface {
	ID() identity.NodeID
	//Run()
	//Register(m interface{}, f interface{}) //接收两个interface{}类型的参数，通常代表某种消息类型（m）和对应的消息处理函数（f）。
}

type node struct {
	id identity.NodeID

	MessageChan chan interface{}
	TxChan      chan interface{}
	sync.RWMutex
}

func (n *node) ID() identity.NodeID {
	return n.id
}

func NewNode(id identity.NodeID) Node {
	return &node{
		id:          id,
		MessageChan: make(chan interface{}, config.Configuration.BufferSize),
		TxChan:      make(chan interface{}, config.Configuration.BufferSize),
	}
}

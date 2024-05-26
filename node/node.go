package node

import (
	"forTest/config"
	"forTest/identity"
	"reflect"
	"sync"
)

type Node interface {
	ID() identity.NodeID
	//Run()
	//Register(m interface{}, f interface{}) //接收两个interface{}类型的参数，通常代表某种消息类型（m）和对应的消息处理函数（f）。
}

type node struct {
	id          identity.NodeID
	handles     map[string]reflect.Value
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
		handles:     make(map[string]reflect.Value),
	}
}

// Register a handle function for each message type
func (n *node) Register(m interface{}, f interface{}) {
	t := reflect.TypeOf(m)
	fn := reflect.ValueOf(f)

	if fn.Kind() != reflect.Func {
		panic("handle function is not func")
	}

	if fn.Type().In(0) != t {
		panic("func type is not t")
	}

	if fn.Kind() != reflect.Func || fn.Type().NumIn() != 1 || fn.Type().In(0) != t {
		panic("register handle function error")
	}
	n.handles[t.String()] = fn
}

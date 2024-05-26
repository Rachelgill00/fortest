package main

import (
	"forTest/identity"
	"forTest/node"
	"forTest/replica"
	"sync"
)

func initReplica(id identity.NodeID) {
	r := replica.NewReplica(id)
	r.Start()

}

func main() {

	node1 := node.NewNode("node1")
	node2 := node.NewNode("node2")
	node3 := node.NewNode("node3")
	node4 := node.NewNode("node4")

	nodes := []node.Node{node1, node2, node3, node4}

	var wg sync.WaitGroup
	for _, n := range nodes {
		wg.Add(1)
		go func(n node.Node) {
			defer wg.Done()
			initReplica(n.ID())
		}(n)
	}
	wg.Wait()
	/*
		batchSize := 200
		txs := make([]*message.Transaction, 0)

		for i := 0; i < batchSize; i++ {
			tx := message.NewTransaction(i)
			txs = append(txs, tx)
		}

		for _, tx := range txs {
			fmt.Println(tx.ID)
			fmt.Println(tx.Timestamp)
		}
	*/

}

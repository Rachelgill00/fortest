package main

import (
	"fmt"
	"forTest/message"
)

func main() {
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
}

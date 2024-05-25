package message

import (
	"strconv"
	"time"

	"github.com/gitferry/bamboo/identity"
)

type Transaction struct {
	Timestamp time.Time
	NodeID    identity.NodeID // forward by node
	ID        string
	//C         chan TransactionReply // reply channel created by request receiver
	ReadSet  []string
	WriteSet []string
}

func NewTransaction(id int) *Transaction {
	tx := new(Transaction)
	tx.ID = strconv.Itoa(id)
	tx.Timestamp = time.Now()

	return tx
}

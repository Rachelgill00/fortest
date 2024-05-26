package message

import (
	"strconv"
	"time"

	"github.com/gitferry/bamboo/identity"
	"github.com/gitferry/bamboo/types"
)

type Transaction struct {
	Timestamp time.Time
	NodeID    identity.NodeID // forward by node
	ID        string
	//C         chan TransactionReply // reply channel created by request receiver
	ReadSet  []string
	WriteSet []string
}

func NewTransaction(id int, nodeid identity.NodeID) *Transaction {
	tx := new(Transaction)
	tx.NodeID = nodeid
	tx.ID = strconv.Itoa(id)
	tx.Timestamp = time.Now()

	return tx
}

/**************************
 *     Calvin Related     *
 **************************/

//type Sequencer struct {
//}

type Sequencer_Message struct {
	NodeID  identity.NodeID
	CurView types.View
	TXN     *Transaction
	//Txn     []*Transaction
}

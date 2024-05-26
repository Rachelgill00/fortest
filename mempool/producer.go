package mempool

import (
	"forTest/message"

	"github.com/gitferry/bamboo/config"
)

type Producer struct {
	mempool *MemPool
}

func NewProducer() *Producer {
	return &Producer{
		mempool: NewMemPool(),
	}
}

// mempool에서 transaction payload 생선
func (pd *Producer) GeneratePayload() []*message.Transaction {
	return pd.mempool.some(config.Configuration.BSize)
}

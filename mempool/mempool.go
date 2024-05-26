package mempool

import "github.com/gitferry/bamboo/config"

type MemPool struct {
	*Backend
}

// NewTransactions creates a new memory pool for transactions.
// Mempool的构造函数，用于创建一个Mempool实例。
func NewMemPool() *MemPool {
	mp := &MemPool{
		Backend: NewBackend(config.GetConfig().MemSize),
	}

	return mp
}

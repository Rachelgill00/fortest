package identity

import "strconv"

// NodeID represents a generic identifier in format of Zone.Node
type NodeID string

// NewNodeID returns a new NodeID type given two int number of zone and node
// 函数主要功能是根据传入的node参数创建一个新的NodeID
func NewNodeID(node int) NodeID {
	if node < 0 {
		node = -node
	}
	// return NodeID(fmt.Sprintf("%d.%d", zone, node))
	return NodeID(strconv.Itoa(node))
}

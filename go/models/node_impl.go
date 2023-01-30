package models

type NodeImpl struct {
	node   *Node
	nodeCb *NodeCB

	// if a gong object has a corresponding elt in the diagram
	// DiagramEltInterface points to the diagram elt
	DiagramEltInterface
}

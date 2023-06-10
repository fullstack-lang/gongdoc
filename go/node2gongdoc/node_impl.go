package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

type NodeImpl struct {
	node   *gongdoc_models.Node
	nodeCb *NodeCB
}

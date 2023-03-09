package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func applyGongNodesConfiguration(node *gongdoc_models.Node, isCheckboxDisabled, isChecked, hasToBeDisabledValue bool) {

	node.IsCheckboxDisabled = isCheckboxDisabled
	node.IsChecked = isChecked

	nodeImpl := node.Impl

	if nodeImpl != nil {
		nodeImpl.SetHasToBeCheckedValue(false)
		nodeImpl.SetHasToBeDisabledValue(hasToBeDisabledValue)
	}

	for _, _node := range node.Children {
		applyGongNodesConfiguration(_node, isCheckboxDisabled, isChecked, hasToBeDisabledValue)
	}
}

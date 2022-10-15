package models

import (
	"log"
	"strconv"
)

type CallbacksSingloton struct {
	ClassdiagramsRootNode *Node
}

func (callbacksSingloton CallbacksSingloton) OnAfterUpdate(
	stage *StageStruct,
	staged, new *Node) {

	if staged.IsChecked != new.IsChecked {
		// log.Println("Node " + staged.Name + " is updated with value IsChecked to " + strconv.FormatBool(new.IsChecked))

		// setting the value of the staged node	to the new value
		staged.IsChecked = new.IsChecked

		// parse all nodes and check wether, if checked/not checked, the is the updated node
		for _, node := range callbacksSingloton.ClassdiagramsRootNode.Children {
			if node == staged {
				continue
			}
			if new.IsChecked && node.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should unchecked")
				node.IsChecked = false
				node.Commit()
			}
			if !new.IsChecked && !node.IsChecked {
				// log.Println("Node " + node.Name + " is checked and should unchecked")
				node.IsChecked = true
				node.Commit()
			}
		}
	}
	if staged.IsExpanded != new.IsExpanded {
		log.Println("Node " + staged.Name + " is updated with value IsExpanded to " + strconv.FormatBool(new.IsExpanded))

		// setting the value of the staged node	to the new value
		staged.IsChecked = new.IsChecked
	}
}

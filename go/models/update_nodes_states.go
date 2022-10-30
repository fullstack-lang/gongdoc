package models

// updateNodesStates updates the tree of symbols
// according to the selected diagram
func updateNodesStates(stage *StageStruct, callbacksSingloton *CallbacksSingloton) {

	// get the editable state of the package
	var pkglet *Pkgelt
	for _pkgelt := range *GetGongstructInstancesSet[Pkgelt]() {
		pkglet = _pkgelt
	}

	// map of gognstruct nodes according to their name
	gongstructNodes := make(map[string]*Node)

	// map of gong fileds according to their name
	gongfieldNodes := make(map[string]*Node)

	// unckeck nodes and construct the map
	for _, gognstructNode := range callbacksSingloton.GongstructsRootNode.Children {

		gognstructNode.IsCheckboxDisabled = !pkglet.Editable

		gognstructNode.IsChecked = false
		gongstructNodes[gognstructNode.Name] = gognstructNode

		for _, gongfieldNode := range gognstructNode.Children {
			gongfieldNode.IsChecked = false
			gongfieldNode.IsCheckboxDisabled = !pkglet.Editable
			gongfieldNodes[gognstructNode.Name+"."+gongfieldNode.Name] = gongfieldNode
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range callbacksSingloton.ClassdiagramsRootNode.Children {

		classdiagramNode.HasEditButton = pkglet.Editable
		classdiagramNode.HasDeleteButton = pkglet.Editable

		if classdiagramNode.IsChecked {
			// get the diagram
			classDiagram := classdiagramNode.Classdiagram

			// get the referenced gongstructs
			for _, classshape := range classDiagram.Classshapes {
				gongstruct := classshape.GongStruct
				gongstructNodes[gongstruct.Name].IsChecked = true

				for _, field := range classshape.Fields {
					gongfieldNodes[gongstruct.Name+"."+field.Name].IsChecked = true
				}
			}
		}
	}

	for _, stateDiagramNode := range callbacksSingloton.StateDiagramsRootNode.Children {
		stateDiagramNode.HasEditButton = pkglet.Editable
		stateDiagramNode.HasDeleteButton = pkglet.Editable
	}
	stage.Commit()
}

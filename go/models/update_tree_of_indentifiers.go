package models

// updateTreeOfIndentifiers updates the tree of symbols
// according to the selected diagram
func updateTreeOfIndentifiers(callbacksSingloton *CallbacksSingloton) {

	// map of gognstruct nodes according to their name
	gongstructNodes := make(map[string]*Node)

	// map of gong fileds according to their name
	gongfieldNodes := make(map[string]*Node)

	// unckeck nodes and construct the map
	for _, gognstructNode := range callbacksSingloton.GongstructsRootNode.Children {
		gognstructNode.IsChecked = false
		gongstructNodes[gognstructNode.Name] = gognstructNode

		for _, gongfieldNode := range gognstructNode.Children {
			gongfieldNode.IsChecked = false
			gongfieldNodes[gognstructNode.Name+"."+gongfieldNode.Name] = gongfieldNode
		}
	}

	// get the selected diagram and collect what are its referenced
	// gongstructs
	for _, classdiagramNode := range callbacksSingloton.ClassdiagramsRootNode.Children {
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

}

package diagrammer

import (
	"log"

	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/maticons/maticons"
)

type Diagrammer struct {
	model     Model
	portfolio Portfolio
	treeStage *gongtree_models.StageStruct

	// selectedDiagram is the currently selected diagram
	selectedDiagram Diagram

	map_portfolioNode_treeNode map[PortfolioNode]*gongtree_models.Node
	map_elementNode_treeNode   map[ElementNode]*gongtree_models.Node
}

func NewDiagrammer(
	model Model,
	portfolio Portfolio,
	treeStage *gongtree_models.StageStruct,

) (diagrammer *Diagrammer) {
	diagrammer = new(Diagrammer)

	diagrammer.model = model
	diagrammer.portfolio = portfolio
	diagrammer.treeStage = treeStage

	diagrammer.map_portfolioNode_treeNode = make(map[PortfolioNode]*gongtree_models.Node)
	diagrammer.map_elementNode_treeNode = make(map[ElementNode]*gongtree_models.Node)

	return
}

func (diagrammer *Diagrammer) AddPortfiolioNodeTreeNodeEntry(portfolioNode PortfolioNode, treeNode *gongtree_models.Node) {
	diagrammer.map_portfolioNode_treeNode[portfolioNode] = treeNode
}

func (diagrammer *Diagrammer) RemovePortfiolioNodeTreeNodeEntry(portfolioNode PortfolioNode) {
	delete(diagrammer.map_portfolioNode_treeNode, portfolioNode)
}

func (diagrammer *Diagrammer) GetPortfiolioNodeFromTreeNode(portfolioNode PortfolioNode) (treeNode *gongtree_models.Node) {

	var ok bool
	treeNode, ok = diagrammer.map_portfolioNode_treeNode[portfolioNode]
	if !ok {
		log.Println("unknown node", portfolioNode.GetName())
		for tn, pn := range diagrammer.map_portfolioNode_treeNode {
			log.Println(tn.GetName(), pn.GetName())
			log.Printf("%p %p\n", pn, tn)
		}
		log.Printf("entry node %s, %p\n", portfolioNode.GetName(), portfolioNode)

		return
	} else {
		return
	}
}

// FillUpModelTree ranges over Model Root Nodes
// and recursively fill up the modelTree
func (diagrammer *Diagrammer) FillUpModelTree(modelTree *gongtree_models.Tree) {

	for _, node := range diagrammer.model.GetChildren() {
		treeNode := diagrammer.modelNode2NodeTree(node, diagrammer.treeStage)
		modelTree.RootNodes = append(modelTree.RootNodes, treeNode)
	}
}

func (diagrammer *Diagrammer) modelNode2NodeTree(modelNode ModelNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: modelNode.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = modelNode.IsExpanded()

	if elementNode, ok := modelNode.(ElementNode); ok {
		treeNode.HasCheckboxButton = true
		treeNode.IsCheckboxDisabled = true
		diagrammer.map_elementNode_treeNode[elementNode] = treeNode
	}

	treeNode.Impl = &ModelNodeImpl{
		diagrammer: diagrammer,
		modelNode:  modelNode}

	for _, childrenModelNode := range modelNode.GetChildren() {
		childrenTreeNode := diagrammer.modelNode2NodeTree(childrenModelNode, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}

// FillUpPortfolioTree ranges over Portfolio Root Nodes
// and recursively fill up the Tree UI from the Portfolio tree
func (diagrammer *Diagrammer) FillUpPortfolioTree(portfolioTree *gongtree_models.Tree) {
	diagrammer.portfolio.GenerateTree()

	for _, portfolioNode := range diagrammer.portfolio.GetChildren() {
		// log.Printf("FillUpPortfolioTree %s %p\n", portfolioNode.GetName(), portfolioNode)
		treeNode := diagrammer.portfolioNode2NodeTree(portfolioNode, diagrammer.treeStage)
		portfolioTree.RootNodes = append(portfolioTree.RootNodes, treeNode)
	}
	diagrammer.generatePortfolioNodesButtons()
}

func (diagrammer *Diagrammer) portfolioNode2NodeTree(portfolioNode PortfolioNode, treeStage *gongtree_models.StageStruct) (treeNode *gongtree_models.Node) {
	treeNode = (&gongtree_models.Node{Name: portfolioNode.GetName()}).Stage(treeStage)
	treeNode.IsExpanded = portfolioNode.IsExpanded()

	// log.Printf("portfolioNode2NodeTree %s %p\n", portfolioNode.GetName(), portfolioNode)
	diagrammer.AddPortfiolioNodeTreeNodeEntry(portfolioNode, treeNode)

	if diagramNode, ok := portfolioNode.(PortfolioDiagramNode); ok {
		_ = diagramNode
		treeNode.HasCheckboxButton = true
		treeNode.Impl = &PortfolioDiagramNodeImpl{
			diagrammer:           diagrammer,
			portfolioDiagramNode: diagramNode}
	}

	treeNode.IsCheckboxDisabled = !diagrammer.portfolio.IsInSelectionMode()

	for _, childrenPortfolioNode := range portfolioNode.GetChildren() {
		childrenTreeNode := diagrammer.portfolioNode2NodeTree(childrenPortfolioNode, treeStage)
		treeNode.Children = append(treeNode.Children, childrenTreeNode)
	}

	return
}

func (diagrammer *Diagrammer) generatePortfolioNodesButtons() {

	for _, portfolioNode := range diagrammer.portfolio.GetChildren() {
		// log.Printf("generatePortfolioNodesButtons %s %p\n", portfolioNode.GetName(), portfolioNode)

		// here the value of "class diagrams" node has changed, 0xc0014665e8
		diagrammer.generatePortfolioNodesButtonsRecursive(portfolioNode)
	}
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) generatePortfolioNodesButtonsRecursive(portfolioNode PortfolioNode) {

	// log.Printf("generatePortfolioNodesButtonsRecursive %s %p\n", portfolioNode.GetName(), portfolioNode)

	treeNode := diagrammer.GetPortfiolioNodeFromTreeNode(portfolioNode)

	// remove all buttons
	for _, _button := range treeNode.Buttons {
		_button.Unstage(diagrammer.treeStage)
	}
	treeNode.Buttons = make([]*gongtree_models.Button, 0)

	if portfolioCategoryNode, ok := portfolioNode.(PortfolioCategoryNode); ok {

		if portfolioCategoryNode.HasAddDiagramButton() {
			addDiagramButton := (&gongtree_models.Button{
				Name: portfolioNode.GetName() + " " + string(maticons.BUTTON_add),
				Icon: string(maticons.BUTTON_add)}).Stage(diagrammer.treeStage)
			treeNode.Buttons = append(treeNode.Buttons, addDiagramButton)
			addDiagramButton.Impl = NewDiagramButtonAddImpl(
				portfolioCategoryNode,
				diagrammer,
				treeNode,
				diagrammer.treeStage,
			)
		}
	}

	if portfolioDiagramNode, ok := portfolioNode.(PortfolioDiagramNode); ok {
		if portfolioDiagramNode.GetDiagram() == diagrammer.selectedDiagram {
			treeNode.IsChecked = true

			type nodeStateEnum int

			const (
				IDLE nodeStateEnum = iota
				IN_RENAME_MODE
				IN_EDIT_MODE
			)
			nodeState := IDLE

			if portfolioDiagramNode.IsInEditMode() {
				nodeState = IN_EDIT_MODE
			} else {
				if portfolioDiagramNode.IsInRenameMode() {
					nodeState = IN_RENAME_MODE
				}
			}

			switch nodeState {
			case IDLE:
				if portfolioDiagramNode.HasEditButton() {
					editDiagramButton := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_add),
						Icon: string(maticons.BUTTON_draw)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, editDiagramButton)
					editDiagramButton.Impl = NewPortfolioDiagramNodeButtonEditImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
					)
				}
				if portfolioDiagramNode.HasDiagramRenameButton() {
					renameDiagramButton := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_add),
						Icon: string(maticons.BUTTON_edit)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, renameDiagramButton)
					renameDiagramButton.Impl = NewPortfolioDiagramNodeButtonRenameImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
					)
				}
				if portfolioDiagramNode.HasDuplicateButton() {
					button := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_file_copy),
						Icon: string(maticons.BUTTON_file_copy)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, button)
					button.Impl = NewPortfolioDiagramNodeButtonDuplicateImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
					)
				}
				if portfolioDiagramNode.HasDeleteButton() {
					button := (&gongtree_models.Button{
						Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_file_copy),
						Icon: string(maticons.BUTTON_delete)}).Stage(diagrammer.treeStage)
					treeNode.Buttons = append(treeNode.Buttons, button)
					button.Impl = NewPortfolioDiagramNodeButtonRemoveImpl(
						portfolioDiagramNode,
						diagrammer,
						treeNode,
						diagrammer.treeStage,
					)
				}
			case IN_EDIT_MODE:
				saveDiagramButton := (&gongtree_models.Button{
					Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_add),
					Icon: string(maticons.BUTTON_draw)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, saveDiagramButton)
				saveDiagramButton.Impl = NewPortfolioDiagramNodeButtonSaveImpl(
					portfolioDiagramNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
				)
				editCancelDiagramButton := (&gongtree_models.Button{
					Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_add),
					Icon: string(maticons.BUTTON_edit_off)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, editCancelDiagramButton)
				editCancelDiagramButton.Impl = NewPortfolioDiagramNodeButtonEditCancelImpl(
					portfolioDiagramNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
				)
			case IN_RENAME_MODE:
				renameCancelDiagramButton := (&gongtree_models.Button{
					Name: portfolioDiagramNode.GetName() + " " + string(maticons.BUTTON_add),
					Icon: string(maticons.BUTTON_edit_off)}).Stage(diagrammer.treeStage)
				treeNode.Buttons = append(treeNode.Buttons, renameCancelDiagramButton)
				renameCancelDiagramButton.Impl = NewPortfolioDiagramNodeButtonRenameCancelImpl(
					portfolioDiagramNode,
					diagrammer,
					treeNode,
					diagrammer.treeStage,
				)
			}
		} else {
			treeNode.IsChecked = false
		}
	}

	for _, childrenPortfolioNode := range portfolioNode.GetChildren() {
		diagrammer.generatePortfolioNodesButtonsRecursive(childrenPortfolioNode)
	}
}

func (diagrammer *Diagrammer) CommitTreeStage() {
	diagrammer.treeStage.Commit()
}

func (diagrammer *Diagrammer) GetMap_elementNode_treeNode() map[ElementNode]*gongtree_models.Node {
	return diagrammer.map_elementNode_treeNode
}

// computeModelNodeStatus parses all nodes in the Model Tree
// and unchecks all nodes unless nodes matches an element in the diagram
func (diagrammer *Diagrammer) computeModelNodeStatus(map_ModelNode_Shape map[ModelNode]Shape) {
	for modelNode, treeNode := range diagrammer.map_elementNode_treeNode {
		treeNode.IsChecked = false
		if _, ok := map_ModelNode_Shape[modelNode]; ok {
			treeNode.IsChecked = true
		}
	}
}

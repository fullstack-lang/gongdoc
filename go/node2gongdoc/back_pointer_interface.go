package node2gongdoc

// BackPointerInterface is what a gong object need to perform on the gong node
// It only need to disable / uncheck them
type BackPointerInterface interface {
	DisableNodeCheckbox()
	CheckNode()
}

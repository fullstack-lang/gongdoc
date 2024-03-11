package adapter

type ElementNodeBase struct {
	portfolioAdapter *PortfolioAdapter
}

func (base *ElementNodeBase) IsExpanded() bool {
	return false
}

func (base *ElementNodeBase) SetIsExpanded(bool) {}

func (base *ElementNodeBase) IsNameEditable() bool {
	return false
}

func (base *ElementNodeBase) CanBeAddedToDiagram() bool {
	return true
}

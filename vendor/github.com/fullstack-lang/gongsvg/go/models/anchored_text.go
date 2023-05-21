package models

type AnchoredText struct {
	Name    string
	Content string

	X_Offset float64
	Y_Offset float64

	// "bold", "normal", ...
	FontWeight string

	Presentation
	Animates []*Animate

	Impl AnchoredTextImplInterface
}

func (anchoredText *AnchoredText) OnAfterUpdate(stage *StageStruct, _, frontAnchoredText *AnchoredText) {

	if anchoredText.Impl != nil {
		anchoredText.Impl.AnchoredTextUpdated(frontAnchoredText)
	}
}

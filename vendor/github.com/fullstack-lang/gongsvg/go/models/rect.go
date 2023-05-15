package models

import "log"

type Rect struct {
	Name string

	X, Y, Width, Height, RX float64
	Presentation

	Animations []*Animate

	IsSelectable bool // alllow selected
	IsSelected   bool

	CanHaveLeftHandle bool
	HasLeftHandle     bool

	CanHaveRightHandle bool
	HasRightHandle     bool

	CanHaveTopHandle bool
	HasTopHandle     bool

	CanHaveBottomHandle bool
	HasBottomHandle     bool

	CanMoveHorizontaly bool
	CanMoveVerticaly   bool

	RectAnchoredTexts []*RectAnchoredText
	RectAnchoredRects []*RectAnchoredRect

	Impl RectImplInterface
}

// OnAfterUpdate, notice that rect == stagedRect
func (rect *Rect) OnAfterUpdate(stage *StageStruct, _, frontRect *Rect) {

	log.Println("Rect, OnAfterUpdate", rect.Name)

	// behavior logic
	if frontRect.IsSelected != rect.IsSelected {
		rect.IsSelected = frontRect.IsSelected
		if frontRect.IsSelected && frontRect.CanHaveLeftHandle {
			rect.HasLeftHandle = true
		} else {
			rect.HasLeftHandle = false
		}
		if frontRect.IsSelected && frontRect.CanHaveRightHandle {
			rect.HasRightHandle = true
		} else {
			rect.HasRightHandle = false
		}
		if frontRect.IsSelected && frontRect.CanHaveTopHandle {
			rect.HasTopHandle = true
		} else {
			rect.HasTopHandle = false
		}
		if frontRect.IsSelected && frontRect.CanHaveBottomHandle {
			rect.HasBottomHandle = true
		} else {
			rect.HasBottomHandle = false
		}
		rect.Commit(stage)
	}

	if rect.X != frontRect.X ||
		rect.Y != frontRect.Y ||
		rect.Width != frontRect.Width ||
		rect.Height != frontRect.Height {

		rect.X = frontRect.X
		rect.Y = frontRect.Y
		rect.Width = frontRect.Width
		rect.Height = frontRect.Height

		rect.Commit(stage)

		if rect.Impl != nil {
			rect.Impl.RectUpdated(rect)
		}
	}
}

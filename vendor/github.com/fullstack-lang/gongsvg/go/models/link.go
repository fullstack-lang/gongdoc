package models

import "log"

type Link struct {
	Name string

	Type LinkType

	Start           *Rect
	StartAnchorType AnchorType

	End           *Rect
	EndAnchorType AnchorType

	// if link type is floating orthogonal ratio, from 0 to 1,
	// where the anchor starts on the edge (horizontal / vertical)
	StartOrientation OrientationType
	StartRatio       float64
	EndOrientation   OrientationType
	EndRatio         float64

	// in case StartOrientation is the same as EndOrientation,
	// there is a perpendicular line that reach the corner at
	// CornerOffsetRatio
	CornerOffsetRatio float64

	// corner radius
	CornerRadius float64

	// Arrows
	HasEndArrow  bool
	EndArrowSize float64

	// to be displayed at the end
	TextAtArrowEnd []*AnchoredText

	// to be displayed at the start
	TextAtArrowStart []*AnchoredText

	// for non floating orthogonal anchors
	ControlPoints []*Point

	Presentation
}

func (link *Link) OnAfterUpdate(stage *StageStruct, _, frontLink *Link) {

	log.Println("Link, OnAfterUpdate", link.Name)
}

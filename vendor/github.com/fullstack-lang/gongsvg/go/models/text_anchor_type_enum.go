package models

type TextAnchorType string

// values for https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/text-anchor
const (
	TEXT_ANCHOR_LEFT   TextAnchorType = "left"
	TEXT_ANCHOR_RIGHT  TextAnchorType = "right"
	TEXT_ANCHOR_CENTER TextAnchorType = "middle"
)

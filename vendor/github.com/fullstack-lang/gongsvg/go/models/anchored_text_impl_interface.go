package models

type AnchoredTextImplInterface interface {

	// AnchoredTextUpdated function is called each time a AnchoredText is modified
	AnchoredTextUpdated(updatedAnchoredText *AnchoredText)
}

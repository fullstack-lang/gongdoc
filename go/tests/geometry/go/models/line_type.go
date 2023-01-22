package models

type LineTypeString string

// values for EnumType
const (
	CONTINUOUS LineTypeString = "CONTINUOUS"
	DOTTED     LineTypeString = "DOTTED"
)

type LineTypeInt int

// values for EnumType
const (
	CONTINUOUS_ZERO LineTypeInt = iota
	DOTTED_ONE
)

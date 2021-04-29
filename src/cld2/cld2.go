package cld2

type ILanguages interface {
	GetEstimates() []IEstimate
	GetTextBytes() int
	GetReliable() bool
}

type IEstimate interface {
	GetLanguage() uint16
	GetPercent() int
	GetNormScore() float64
}

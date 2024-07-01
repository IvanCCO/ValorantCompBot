package valorant

type ValorantAgent string

const (
	Jeet    ValorantAgent = "Jeet"
	Raze    ValorantAgent = "Raze"
	Iso     ValorantAgent = "Iso"
	Sage    ValorantAgent = "Sage"
	Cypher  ValorantAgent = "Cypher"
	Phoenix ValorantAgent = "Phoenix"
	Viper   ValorantAgent = "Viper"
)

var AllCharacters = []ValorantAgent{Jeet, Raze, Iso, Sage, Cypher}

package valorant

import (
	"math/rand"
)

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

var AllCharacters = []ValorantAgent{Jeet, Raze, Iso, Sage, Cypher, Phoenix, Viper}

func GetRandomCharacters(n int) []ValorantAgent {
	rand.Shuffle(len(AllCharacters), func(i, j int) { AllCharacters[i], AllCharacters[j] = AllCharacters[j], AllCharacters[i] })

	if n > len(AllCharacters) {
		n = len(AllCharacters)
	}

	return AllCharacters[:n]
}

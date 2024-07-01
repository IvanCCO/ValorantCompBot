package commands

import (
	"fmt"

	"github.com/IvanCCO/ValorantCompBot/config"
	"github.com/IvanCCO/ValorantCompBot/valorant"
)

func RandomCommand(ctx config.Context) {
	ctx.Reply(mockedValorantComp())
}

func mockedValorantComp() string {
	agents := valorant.GetRandomCharacters(5)
	result := ""

	for i, agent := range agents {
		result += fmt.Sprintf("%d. %s\n", i+1, agent)
	}

	return result
}

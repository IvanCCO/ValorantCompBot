package commands

import (
	"fmt"
	"math/rand"

	"github.com/IvanCCO/ValorantCompBot/internal/api"
	"github.com/IvanCCO/ValorantCompBot/internal/config"
)

func RandomCommand(ctx config.Context) {
	agents, err := api.GetAgents()
	if err != nil {
		return
	}

	if len(agents) < 5 {
		return
	}

	rand.Shuffle(len(agents), func(i, j int) { agents[i], agents[j] = agents[j], agents[i] })

	selectedAgents := agents[:5]
	result := "Boa Sorte 🤣:\n"
	for i, agent := range selectedAgents {
		result += fmt.Sprintf("%d. %s\n", i+1, agent.DisplayName)
	}

	ctx.Reply(result)
}

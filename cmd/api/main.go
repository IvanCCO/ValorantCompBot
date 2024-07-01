package main

import (
	"fmt"
	"log"

	"github.com/IvanCCO/ValorantCompBot/internal/api"
)

func main() {
	agents, err := api.GetAgents()
	if err != nil {
		log.Fatal(err)
	}

	for _, agent := range agents {
		fmt.Printf("UUID: %s, DisplayName: %s\n", agent.UUID, agent.DisplayName)
	}
}

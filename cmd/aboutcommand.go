package cmd

import (
	"github.com/IvanCCO/ValorantCompBot/config"
)

const ABOUT_MESSAGE = `
  Bot criado para ajudar os times a escolher os personagens quando estão com dúvidas
  da composição. 

  __criado por @Oliversss__
  `

func AboutCommand(ctx config.Context) {
	ctx.Reply(ABOUT_MESSAGE)
}

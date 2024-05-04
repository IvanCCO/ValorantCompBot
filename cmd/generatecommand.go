package cmd

import (
	"bytes"

	"github.com/IvanCCO/ValorantCompBot/config"
)

func GenerateCommand(ctx config.Context) {
	buffer := bytes.NewBufferString("Choose Valorant Agent composition: \n")
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}

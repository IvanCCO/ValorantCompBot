package commands

import (
	"bytes"
	"fmt"

	"github.com/IvanCCO/ValorantCompBot/internal/config"
)

func HelpCommand(ctx config.Context) {
	cmds := ctx.CmdHandler.GetCmds()
	buffer := bytes.NewBufferString("Commands: \n")
	for cmdName, cmdStruct := range cmds {
		if len(cmdName) == 1 {
			continue
		}
		msg := fmt.Sprintf("\t %s - %s\n", cmdName, cmdStruct.GetHelp())
		buffer.WriteString(msg)
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}

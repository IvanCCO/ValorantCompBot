package cmd

import (
	"bytes"
	"fmt"

	"github.com/IvanCCO/ValorantCompBot/config"
)

func HelpCommand(ctx config.Context) {
	cmds := ctx.CmdHandler.GetCmds()
	buffer := bytes.NewBufferString("Commands: \n")
	for cmdName, cmdStruct := range cmds {
		if len(cmdName) == 1 {
			continue
		}
		msg := fmt.Sprintf("\t %s%s - %s\n", ctx.Conf.Prefix, cmdName, cmdStruct.GetHelp())
		buffer.WriteString(msg)
	}
	str := buffer.String()
	ctx.Reply(str[:len(str)-2])
}

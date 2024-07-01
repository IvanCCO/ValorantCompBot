package config

type (
	Command func(Context)

	CommandStruct struct {
		command Command
		help    string
	}

	CmdMap map[string]CommandStruct

	CommandHandler struct {
		cmds CmdMap
	}
)

func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

func (handler CommandHandler) GetCmds() CmdMap {
	return handler.cmds
}

func (handler CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := handler.cmds[name]
	return &cmd.command, found
}

func (handler CommandHandler) Register(name string, command Command, helpmsg string) {
	cmdstruct := CommandStruct{command: command, help: helpmsg}
	handler.cmds[name] = cmdstruct
	if len(name) > 1 {
		handler.cmds[name[:1]] = cmdstruct
	}
}

func (command CommandStruct) GetHelp() string {
	return command.help
}

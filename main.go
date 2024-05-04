package main

import (
	"fmt"
	"strings"

	"github.com/IvanCCO/ValorantCompBot/cmd"
	"github.com/IvanCCO/ValorantCompBot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	conf       *config.Config
	CmdHandler *config.CommandHandler
	botId      string
	PREFIX     string
)

func init() {
	conf = config.LoadConfig("config.json")
	PREFIX = conf.Prefix
}

func main() {
	CmdHandler = config.NewCommandHandler()
	registerCommands()
	discord, err := discordgo.New(conf.BotToken)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
		return
	}
	discord.AddHandler(commandHandler)
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botId || user.Bot {
		return
	}
	content := message.Content
	if len(content) <= len(PREFIX) {
		return
	}
	if content[:len(PREFIX)] != PREFIX {
		return
	}
	content = content[len(PREFIX):]
	if len(content) < 1 {
		return
	}
	args := strings.Fields(content)
	name := strings.ToLower(args[0])
	command, found := CmdHandler.Get(name)
	if !found {
		return
	}
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		return
	}
	guild, err := discord.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println("Error getting guild,", err)
		return
	}
	ctx := config.NewContext(discord, guild, channel, user, message, CmdHandler)
	ctx.Args = args[1:]
	c := *command
	c(*ctx)
}

func registerCommands() {
	CmdHandler.Register("help", cmd.HelpCommand, "Gives you this help message!")
}

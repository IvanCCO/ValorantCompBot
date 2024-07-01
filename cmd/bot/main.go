package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IvanCCO/ValorantCompBot/config"
	"github.com/IvanCCO/ValorantCompBot/internal/commands"
	"github.com/bwmarrin/discordgo"
)

var (
	CmdHandler *config.CommandHandler
	Token      string
	PREFIX     string
)

func init() {
	Token = ""
	PREFIX = "!V"
}

func main() {
	CmdHandler = config.NewCommandHandler()
	registerCommands()
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating discord session,", err)
		return
	}
	discord.AddHandler(commandHandler)

	discord.Identify.Intents = discordgo.IntentsAll

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == discord.State.User.ID || user.Bot {
		return
	}
	content := message.Content
	if len(content) <= len(PREFIX) {
		return
	}
	if content[:len(PREFIX)] != PREFIX {
		return
	}

	args := strings.Fields(content)
	name := strings.ToLower(args[1])
	command, found := CmdHandler.Get(name)
	if !found {
		return
	}

	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("ChannelId ,", message.ChannelID)
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
	CmdHandler.Register("sobre", commands.AboutCommand, "Sobre o ValorantCompBot")
	CmdHandler.Register("ajuda", commands.HelpCommand, "Mostra os comandos disponíveis e sua descrição")
	CmdHandler.Register("aleatorio", commands.RandomCommand, "Composição aleatória de Valorant - Atualmente mocado")
}

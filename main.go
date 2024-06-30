package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IvanCCO/ValorantCompBot/cmd"
	"github.com/IvanCCO/ValorantCompBot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	CmdHandler *config.CommandHandler
	Token      string
	PREFIX     string
)

func init() {
	Token = ""
	PREFIX = "/"
}

func main() {
	CmdHandler = config.NewCommandHandler()
	registerCommands()
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating discord session,", err)
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
	// Skipa caso o produtor da mensagem seja o próprio BOT
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
	name := strings.ToLower(args[0])
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
	CmdHandler.Register("/help", cmd.HelpCommand, "Gives you this help message!")
}

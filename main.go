package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IvanCCO/ValorantCompBot/cmd"
	"github.com/IvanCCO/ValorantCompBot/config"
	"github.com/bwmarrin/discordgo"
)

var CmdHandler *config.CommandHandler

func main() {
	CmdHandler = config.NewCommandHandler()
	registerCommands()
	discord, err := discordgo.New("Bot " + "MTI1Njk3OTMxNTA1MzE3MDc4MA.G0FboV._wtHoAuzST8o_uGW0XpO4gkG2WvBMMvpk3XW7U")
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

	// Cleanly close down the Discord session.
	discord.Close()
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Skipa caso o produtor da mensagem seja o próprio BOT
	user := message.Author
	if user.ID == discord.State.User.ID || user.Bot {
		return
	}
	content := message.Content

	fmt.Println(content)

	args := strings.Fields(content)
	command, found := CmdHandler.Get("/help")
	if !found {
		return
	}
	channel, err := discord.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel,", err)
		msg, _ := json.Marshal(message)
		fmt.Println("Message ,", string(msg))
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
